package snum

type T_Encoder struct {
	Snum
}

func (t *T_Encoder) Init() {
	t.Snum.Init(DEF_b1_header__max_len__standard, DEF_b1_header__max_len__decimal)
}

func (t *T_Encoder) Encode() (bt_ret []byte, err error) {
	// 문자열 추출 ( T_Snum 사용 )
	var s_raw string
	var n_len__total int
	var n_len__decimal int
	var is_minus bool
	{
		s_raw, n_len__decimal, is_minus = t.Snum.Get__raw()
		// s_num 이 0일 경우 후처리
		if s_raw == "0" {
			n_len__decimal = DEF_b1_header__max_len__decimal
		}
		n_len__total = len(s_raw)
	}

	// 헤더 제작
	n_pos_start_dot := t.header__make__pos_start_dot(n_len__total, n_len__decimal)
	b1_header := t.header__make_header(n_pos_start_dot, is_minus)

	// 데이터 제작
	var bt_num__4bit []byte
	{
		bt_num__4bit = make([]byte, 0, n_len__total) // n_len / 2 + (n_len%2!=0)?1:0
		bt_num__ori := []byte(s_raw)
		for i := 0; i < n_len__total; i++ {
			b1_one_num_bit := bt_num__ori[i] - byte('0')
			if i%2 == 0 {
				bt_num__4bit = append(bt_num__4bit, b1_one_num_bit<<4)
			} else {
				bt_num__4bit[i/2] += b1_one_num_bit
			}
		}
		// 부호가 음수(-) 일 경우 데이터 비트 반전
		if is_minus == true {
			n_len_data := len(bt_num__4bit)
			for i := 0; i < n_len_data; i++ {
				bt_num__4bit[i] = ^bt_num__4bit[i] // 비트 반전
			}

			// 음수일 경우 정렬을 위해 0xFF 추가
			bt_num__4bit = append(bt_num__4bit, 0xFF)
		}
	}

	// 헤더와 데이터를 합쳐 bt_ret 제작
	bt_ret = make([]byte, 0, DEF_n_header__size+(n_len__total/2))
	bt_ret = append(bt_ret, b1_header)
	bt_ret = append(bt_ret, bt_num__4bit...)
	return bt_ret, nil
}

func (t *T_Encoder) Decode(_bt_num []byte) (err error) {
	if len(_bt_num) < DEF_n_bt__len_min_total {
		return Err_header_not_enongh
	}

	// 헤더 정보 추출 - 부호 / 길이
	var is_minus bool
	var b1_len_header byte
	{
		is_minus, b1_len_header = t.header__decode(_bt_num[0])
		// _bt_num 이 0 일 경우 처리
		if len(_bt_num) == 2 && _bt_num[1] == 0 {
			b1_len_header = byte(DEF_b1_header__max_len__decimal)
		}
	}

	// 데이터 추출 및 big int 설정
	var s_raw string
	var n_len__decimal int
	{
		bt_data := _bt_num[1:]
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
		t.Snum.Set__raw(s_raw, n_len__decimal, is_minus)
	}
	return nil
}

//------------------------------------------------------------------------------------------//
// util ( header )

func (t *T_Encoder) header__decode(_b1_header byte) (is_minus bool, b1_len_standard byte) {
	if _b1_header&DEF_b1_header__bit_mask__sign == 0 {
		// 부호(+-) 추출
		is_minus = true
		// 음수일 경우 헤더 보수처리
		_b1_header = ^_b1_header
	}

	// 헤더에서 정수길이만 추출
	b1_len_standard = _b1_header & DEF_b1_header__bit_mask__standard_len

	return is_minus, b1_len_standard
}

func (t *T_Encoder) header__make__len_decimal(_n_len int, _b1_len_starndard byte) (n_len_decimal int) {
	// 소수 길이 추출
	n_len_decimal = _n_len - int(_b1_len_starndard) + DEF_b1_header__max_len__decimal - 1
	return n_len_decimal
}

func (t *T_Encoder) header__make__pos_start_dot(_n_len__total int, _n_len__decimal int) (n_pos_start_dot int) {
	// 소수점 시작 위치 추출
	n_pos_start_dot = _n_len__total - _n_len__decimal + DEF_b1_header__max_len__decimal - 1
	return n_pos_start_dot
}

func (t *T_Encoder) header__make_header(_n_pos_start_dot int, is_minus bool) (b1_header byte) {
	// 헤더 제작 - 제작시 양수로 가정하고 제작 후 -> 후 처리에서 음수를 반영
	b1_header = DEF_b1_header__value__sign__plus | byte(_n_pos_start_dot)

	// 음수의 경우 비트반전
	if is_minus == true {
		b1_header = ^b1_header
	}
	return b1_header
}
