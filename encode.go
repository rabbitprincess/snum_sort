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
		is_minus, b1_len_header = t.header__decode(_num[0])
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
		n_len__decimal = t.header__make__len_decimal(n_len__total, b1_len_header)
	}

	// snum 세팅 ( T_Snum 사용 )
	{
		t.Snum.SetRaw(s_raw, n_len__decimal, is_minus)
	}
	return nil
}

//------------------------------------------------------------------------------------------//
// util ( header )

func (t *Encoder) header__decode(_b1_header byte) (is_minus bool, b1_len_standard byte) {
	if _b1_header&DEF_headerBitMaskSign == 0 {
		// 부호(+-) 추출
		is_minus = true
		// 음수일 경우 헤더 보수처리
		_b1_header = ^_b1_header
	}

	// 헤더에서 정수길이만 추출
	b1_len_standard = _b1_header & DEF_headerBitMaskStandardLen

	return is_minus, b1_len_standard
}

func (t *Encoder) header__make__len_decimal(_n_len int, _b1_len_starndard byte) (n_len_decimal int) {
	// 소수 길이 추출
	n_len_decimal = _n_len - int(_b1_len_starndard) + DEF_headerLenDecimal - 1 // -1 이유 = 1의 자리가 0번 idx 지만 길이는 1의 자리가 len 1 이기 때문에 1 감소로 1의 자리를 0 번으로 맞춘다.
	return n_len_decimal
}

func (t *Encoder) makePosStartDot(_n_len__total int, _n_len__decimal int) (n_pos_start_dot int) {
	// 소수점 시작 위치 추출
	n_pos_start_dot = _n_len__total - _n_len__decimal + DEF_headerLenDecimal - 1 // -1 이유 = 1의 자리가 0번 idx 지만 길이는 1의 자리가 len 1 이기 때문에 1 감소로 1의 자리를 0 번으로 맞춘다.
	return n_pos_start_dot
}

func (t *Encoder) makeHeader(_n_pos_start_dot int, is_minus bool) (b1_header byte) {
	// 헤더 제작 - 제작시 양수로 가정하고 제작 후 -> 후 처리에서 음수를 반영
	b1_header = DEF_headerValueSignPlus | byte(_n_pos_start_dot)

	// 음수의 경우 비트반전
	if is_minus == true {
		b1_header = ^b1_header
	}
	return b1_header
}
