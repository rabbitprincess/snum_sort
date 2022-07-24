package snum

//------------------------------------------------------------------------------------------//
// binary - sorted

type Encoder struct {
	Snum
}

func (t *Encoder) Init() {
	t.Snum.Init(DEF_headerLenInteger, DEF_headerLenDecimal)
}

func (t *Encoder) Encode() (ret []byte, err error) {
	// 문자열 추출 ( T_Snum 사용 )
	var raw string
	var lenTotal int
	var lenDecimal int
	var isMinus bool
	{
		raw, lenDecimal, isMinus = t.Snum.GetRaw()
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

func (t *Encoder) Decode(_num []byte) (err error) {
	if len(_num) < DEF_lenDataMinTotal {
		return ErrHeaderNotEnough
	}

	// 헤더 정보 추출 - 부호 / 길이
	var is_minus bool
	var b1_len_header byte
	{
		is_minus, b1_len_header = t.decodeHeader(_num[0])
		// _bt_num 이 0 일 경우 처리
		if len(_num) == 2 && _num[1] == 0 {
			b1_len_header = byte(DEF_headerLenDecimal)
		}
	}

	// 데이터 추출 및 big int 설정
	var s_raw string
	var n_len__decimal int
	{
		bt_data := _num[1:]
		// 전처리 - 헤더에 따른 정보가 음수 일 경우
		if is_minus == true {
			// 마지막 0xFF 분리
			bt_data = bt_data[:len(bt_data)-1]
			// 비트 반전(데이터)
			for i := 0; i < len(bt_data); i++ {
				bt_data[i] = ^bt_data[i]
			}
		}

		// string 제작
		for i := 0; i < len(bt_data); i++ {
			b1_num__high_4bit := bt_data[i] >> 4
			b1_num__low_4bit := bt_data[i] - (b1_num__high_4bit << 4)
			s_raw += string('0' + b1_num__high_4bit)
			s_raw += string('0' + b1_num__low_4bit)
		}

		// n_len__decimal 추출
		n_len__total := len(s_raw)
		n_len__decimal = t.makeLenDecimal(n_len__total, b1_len_header)
	}

	// snum 세팅 ( T_Snum 사용 )
	{
		t.Snum.SetRaw(s_raw, n_len__decimal, is_minus)
	}
	return nil
}

//------------------------------------------------------------------------------------------//
// util ( header )

func (t *Encoder) decodeHeader(_header byte) (isMinus bool, lenStandard byte) {
	if _header&DEF_headerBitMaskSign == 0 {
		// 부호(+-) 추출
		isMinus = true
		// 음수일 경우 헤더 보수처리
		_header = ^_header
	}

	// 헤더에서 정수길이만 추출
	lenStandard = _header & DEF_headerBitMaskStandardLen

	return isMinus, lenStandard
}

func (t *Encoder) makeLenDecimal(_len int, _lenStarndard byte) (lenDecimal int) {
	// 소수 길이 추출
	lenDecimal = _len - int(_lenStarndard) + DEF_headerLenDecimal - 1 // -1 이유 = 1의 자리가 0번 idx 지만 길이는 1의 자리가 len 1 이기 때문에 1 감소로 1의 자리를 0 번으로 맞춘다.
	return lenDecimal
}

func (t *Encoder) makePosStartDot(_lenTotal int, _lenDecimal int) (posStartDot int) {
	// 소수점 시작 위치 추출
	posStartDot = _lenTotal - _lenDecimal + DEF_headerLenDecimal - 1 // -1 이유 = 1의 자리가 0번 idx 지만 길이는 1의 자리가 len 1 이기 때문에 1 감소로 1의 자리를 0 번으로 맞춘다.
	return posStartDot
}

func (t *Encoder) makeHeader(_posStartDot int, _isMinus bool) (header byte) {
	// 헤더 제작 - 제작시 양수로 가정하고 제작 후 -> 후 처리에서 음수를 반영
	header = DEF_headerValueSignPlus | byte(_posStartDot)

	// 음수의 경우 비트반전
	if _isMinus == true {
		header = ^header
	}
	return header
}
