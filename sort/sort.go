package sort

import (
	"math/big"

	"github.com/gokch/snum_sort/snum"
)

type Encoder struct {
	snum.Snum
}

func (t *Encoder) Init() {
	t.Snum.Init()
}

func (t *Encoder) Encode() (ret []byte, err error) {
	// 문자열 추출 ( Snum 사용 )
	var raw string
	var lenTotal int
	var lenDecimal int
	var isMinus bool
	{
		var bigRaw *big.Int
		bigRaw, lenDecimal, isMinus = t.Snum.GetRaw()
		raw = bigRaw.String()

		// s_num 이 0일 경우 후처리
		if raw == "0" {
			lenDecimal = DEF_headerLenDecimal
		}
		lenTotal = len(raw)
	}

	// 헤더 제작
	posStartDot := t.makePosStartDot(lenTotal, lenDecimal)
	header := t.makeHeader(posStartDot, isMinus)

	// 데이터 제작
	var numCompress []byte
	{
		numCompress = make([]byte, 0, lenTotal) // n_len / 2 + (n_len%2!=0)?1:0
		numOri := []byte(raw)
		for i := 0; i < lenTotal; i++ {
			b1_one_num_bit := numOri[i] - byte('0')
			if i%2 == 0 {
				numCompress = append(numCompress, b1_one_num_bit<<4)
			} else {
				numCompress[i/2] += b1_one_num_bit
			}
		}
		// 부호가 음수(-) 일 경우 데이터 비트 반전
		if isMinus == true {
			lenData := len(numCompress)
			for i := 0; i < lenData; i++ {
				numCompress[i] = ^numCompress[i] // 비트 반전
			}

			// 음수의 경우 무조건 끝에 역정렬 알고리즘을 위한 비교마감(cut) 수치 ( 통상올수있는 값 range 보다 더 큰수 ) 를 넣는다.
			numCompress = append(numCompress, 0xFF)
		}
	}

	// 헤더와 데이터를 합쳐 bt_ret 제작
	ret = make([]byte, 0, DEF_headerSize+(lenTotal/2))
	ret = append(ret, header)
	ret = append(ret, numCompress...)
	return ret, nil
}

func (t *Encoder) Decode(arg []byte) (err error) {
	if len(arg) < DEF_lenDataMinTotal {
		return ErrHeaderNotEnough
	}

	// 헤더 정보 추출 - 부호 / 길이
	var isMinus bool
	var lenHeader byte
	{
		isMinus, lenHeader = t.decodeHeader(arg[0])
		// _bt_num 이 0 일 경우 처리
		if len(arg) == 2 && arg[1] == 0 {
			lenHeader = byte(DEF_headerLenDecimal)
		}
	}

	// 데이터 추출 및 big int 설정
	var sRaw string
	var lenDecimal int
	{
		data := arg[1:]
		// 전처리 - 헤더에 따른 정보가 음수 일 경우
		if isMinus == true {
			// 마지막 0xFF 분리
			data = data[:len(data)-1]
			// 비트 반전(데이터)
			for i := 0; i < len(data); i++ {
				data[i] = ^data[i]
			}
		}

		// string 제작
		for i := 0; i < len(data); i++ {
			high4bit := data[i] >> 4
			low4bit := data[i] - (high4bit << 4)
			sRaw += string('0' + high4bit)
			sRaw += string('0' + low4bit)
		}

		// n_len__decimal 추출
		lenTotal := len(sRaw)
		lenDecimal = t.makeLenDecimal(lenTotal, lenHeader)
	}

	// snum 세팅 ( T_Snum 사용 )
	{
		big, _ := big.NewInt(0).SetString(sRaw, 10)
		t.Snum.SetRaw(big, lenDecimal, isMinus)
	}
	return nil
}

//------------------------------------------------------------------------------------------//
// util ( header )

func (t *Encoder) decodeHeader(header byte) (isMinus bool, lenStandard byte) {
	if header&DEF_headerBitMaskSign == 0 {
		// 부호(+-) 추출
		isMinus = true
		// 음수일 경우 헤더 보수처리
		header = ^header
	}

	// 헤더에서 정수길이만 추출
	lenStandard = header & DEF_headerBitMaskStandardLen

	return isMinus, lenStandard
}

func (t *Encoder) makeLenDecimal(len int, lenStarndard byte) (lenDecimal int) {
	// 소수 길이 추출
	lenDecimal = len - int(lenStarndard) + DEF_headerLenDecimal - 1 // -1 이유 = 1의 자리가 0번 idx 지만 길이는 1의 자리가 len 1 이기 때문에 1 감소로 1의 자리를 0 번으로 맞춘다.
	return lenDecimal
}

func (t *Encoder) makePosStartDot(lenTotal int, lenDecimal int) (posStartDot int) {
	// 소수점 시작 위치 추출
	posStartDot = lenTotal - lenDecimal + DEF_headerLenDecimal - 1 // -1 이유 = 1의 자리가 0번 idx 지만 길이는 1의 자리가 len 1 이기 때문에 1 감소로 1의 자리를 0 번으로 맞춘다.
	return posStartDot
}

func (t *Encoder) makeHeader(posStartDot int, isMinus bool) (header byte) {
	// 헤더 제작 - 제작시 양수로 가정하고 제작 후 -> 후 처리에서 음수를 반영
	header = DEF_headerValueSignPlus | byte(posStartDot)

	// 음수의 경우 비트반전
	if isMinus == true {
		header = ^header
	}
	return header
}
