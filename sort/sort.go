package sort

import (
	"errors"
	"math/big"

	"github.com/gokch/snum_sort/snum"
)

func NewSnumSort[T snum.SnumConst](num T) *SnumSort {
	return &SnumSort{
		Snum: *snum.NewSnum(num),
	}
}

type SnumSort struct {
	snum.Snum
}

func (t *SnumSort) Encode() (enc []byte, err error) {
	// get raw
	bigRaw, lenDecimal, isMinus := t.Snum.GetRaw()
	raw := bigRaw.String()
	if raw == "0" { // 0 일 경우 전처리
		lenDecimal = DEF_digitDecimalMax
	}
	lenTotal := len(raw)

	// make header
	posStartDot := t.makePosStartDot(lenTotal, lenDecimal)
	header := t.encodeHeader(posStartDot)

	// make body
	dataCompress := make([]byte, 0, lenTotal) // len / 2 + (len%2!=0)?1:0
	numOri := []byte(raw)
	for i := 0; i < lenTotal; i++ {
		b4 := numOri[i] - byte('0')
		if i%2 == 0 {
			dataCompress = append(dataCompress, b4<<4)
		} else {
			dataCompress[i/2] += b4
		}
	}
	// if minus, reverse bit
	if isMinus == true {
		header = ^header
		lenData := len(dataCompress)
		for i := 0; i < lenData; i++ {
			dataCompress[i] = ^dataCompress[i]
		}

		dataCompress = append(dataCompress, 0xFF) // append last 0xFF
	}

	// make enc
	enc = make([]byte, 0, DEF_lenHeader+(lenTotal/2))
	enc = append(enc, header)
	enc = append(enc, dataCompress...)
	return enc, nil
}

func (t *SnumSort) Decode(enc []byte) (err error) {
	if len(enc) < DEF_lenTotalMin {
		return errors.New("too short")
	}

	// Decode header info
	isMinus, lenHeader := t.decodeHeader(enc[0])
	if len(enc) == 2 && enc[1] == 0 { // 0 일 경우 후처리
		lenHeader = byte(DEF_digitDecimalMax)
	}

	data := enc[1:]
	if isMinus == true {
		data = data[:len(data)-1]        // separate last 0xFF
		for i := 0; i < len(data); i++ { // reverse bit
			data[i] = ^data[i]
		}
	}

	// make string
	var sRaw string
	for i := 0; i < len(data); i++ {
		high4bit := data[i] >> 4
		low4bit := data[i] - (high4bit << 4)
		sRaw += string('0' + high4bit)
		sRaw += string('0' + low4bit)
	}

	// set total, decimal len
	lenTotal := len(sRaw)
	lenDecimal := t.makeLenDecimal(lenTotal, lenHeader)

	// set snum
	big, _ := big.NewInt(0).SetString(sRaw, 10)
	t.Snum.SetRaw(big, lenDecimal, isMinus)
	return nil
}

//------------------------------------------------------------------------------------------//
// util ( header )

func (t *SnumSort) decodeHeader(header byte) (isMinus bool, lenStandard byte) {
	if header&DEF_headerBitMaskSign == 0 {
		// 부호(+-) 추출
		isMinus = true
		header = ^header
	}

	// 헤더에서 정수길이만 추출
	lenStandard = header & DEF_headerBitMaskStandardLen

	return isMinus, lenStandard
}

func (t *SnumSort) makeLenDecimal(len int, lenStarndard byte) (lenDecimal int) {
	// 소수 길이 추출 ( 자릿수는 1번이 0번 idx 지만 길이는 1의 자리가 1 이기 때문에 1 감소 필요 )
	lenDecimal = len - int(lenStarndard) + DEF_digitDecimalMax - 1
	return lenDecimal
}

func (t *SnumSort) makePosStartDot(lenTotal int, lenDecimal int) (posStartDot int) {
	// 소수점 시작 위치 추출 ( 자릿수는 1번이 0번 idx 지만 길이는 1의 자리가 1 이기 때문에 1 감소 필요 )
	posStartDot = lenTotal - lenDecimal + DEF_digitDecimalMax - 1
	return posStartDot
}

func (t *SnumSort) encodeHeader(posStartDot int) (header byte) {
	header = DEF_headerBitMaskSign | byte(posStartDot)
	return header
}
