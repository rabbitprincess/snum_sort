package snum

import (
	"bytes"
	"fmt"
	"sort"
	"strings"
	"testing"
)

func Test__encode_decode(_t *testing.T) {
	fn := func(_s_num__input string, _s_num__expected string) {
		// 기대값 설정
		{
			// 기대값 입력이 비어있으면 입력값과 동일한 것이 정답인 것으로 인정한다.
			if _s_num__expected == "" {
				_s_num__expected = _s_num__input
			}
		}

		// Byte_encode, Byte_decode
		// string -> bigint -> binary -> bigint -> string
		pt_sorted := &T_Encoder{}
		pt_sorted.Init()
		{
			err := pt_sorted.Set__str(_s_num__input)
			if err != nil {
				_t.Errorf("Set__str\n%s:%v\n%s:%v\n", "_s_num__input", _s_num__input, "err", err)
				return
			}

			bt_num, err := pt_sorted.Encode()
			if err != nil {
				_t.Errorf("Decode\n%s:%v\n%s:%v\n", "_s_num__input", _s_num__input, "err", err)
				return
			}

			// for debug
			// fmt.Printf("%30s\n", fmt.Sprintf("%02x", bt_num))

			err = pt_sorted.Decode(bt_num)
			if err != nil {
				_t.Errorf("Encode\n%s:%v\n%s:%v\n", "_s_num__input", _s_num__input, "err", err)
				return
			}

			s_num__recovery, err := pt_sorted.Get__str()
			if err != nil {
				_t.Errorf("Get__str\n%s:%v\n%s:%v\n", "_s_num__input", _s_num__input, "err", err)
				return
			}

			if _s_num__expected != s_num__recovery {
				_t.Errorf("s_num__expected != s_num__recovery\n - %s : %v\n - %s : %v\n\n", "s_num__expected", _s_num__expected, "s_num__recovery", s_num__recovery)
				return
			}
		}
	}

	// 0
	{
		fn("0", "")
		fn(".0", "0")
		fn("0.0", "0")
	}
	// 양수
	{
		// 1 이상
		fn("1", "")
		fn("12", "")
		fn("123", "")
		fn("123.1", "")
		fn("123.12", "")
		fn("123.123", "")
		fn("12.3123", "")
		fn("1.23123", "")

		// 1 미만
		fn("0.1", "")
		fn("0.123123", "")
		fn("0.0123123", "")
		fn("0.00123123", "")
		fn("0.000123123", "")

		fn("1.01", "")
		fn("1.0123456789012345678", "")
		fn("1.01234567890123456789", "")
		fn("10.1234567890123456789", "")
		fn("101.234567890123456789", "")
		fn("1012.34567890123456789", "")
		fn("10123.4567890123456789", "")
		fn("101234.567890123456789", "")
		fn("1012345.67890123456789", "")
		fn("10123456.7890123456789", "")
		fn("101234567.890123456789", "")
		fn("1012345678.90123456789", "")
		fn("10123456789.0123456789", "")
		fn("101234567890.123456789", "")
		fn("1012345678901.23456789", "")
		fn("10123456789012.3456789", "")
		fn("101234567890123.456789", "")
		fn("1012345678901234.56789", "")
		fn("10123456789012345.6789", "")
		fn("101234567890123456.789", "")
		fn("1012345678901234567.89", "")
		fn("10123456789012345678.9", "")
		fn("10123456789012345678.9012345678", "")
		fn("1.012345678901234567891", "")
		fn("1.0123456789012345678901", "")

		// 자릿수 한도
		{
			// 정수는 96 자리 까지 ok
			fn("1"+strings.Repeat("0", DEF_b1_header__max_len__standard-1), "")
			fn(strings.Repeat("9", DEF_b1_header__max_len__standard), "")
			// 소수는 32 자리 까지 ok
			fn("0."+strings.Repeat("0", DEF_b1_header__max_len__decimal-1)+"1", "")

			// 양수 최대값
			fn(strings.Repeat("9", DEF_b1_header__max_len__standard)+"."+strings.Repeat("9", DEF_b1_header__max_len__decimal), "")
		}
	}

	// 음수
	{
		// -1 이하
		fn("-1", "")
		fn("-12", "")
		fn("-123", "")
		fn("-123.1", "")
		fn("-123.12", "")
		fn("-123.123", "")
		fn("-12.3123", "")
		fn("-1.23123", "")

		// -1 초과
		fn("-0.1", "")
		fn("-0.123123", "")
		fn("-0.0123123", "")
		fn("-0.00123123", "")
		fn("-0.000123123", "")

		fn("-1.01", "")
		fn("-1.0123456789012345678", "")
		fn("-1.01234567890123456789", "")
		fn("-10.1234567890123456789", "")
		fn("-101.234567890123456789", "")
		fn("-1012.34567890123456789", "")
		fn("-10123.4567890123456789", "")
		fn("-101234.567890123456789", "")
		fn("-1012345.67890123456789", "")
		fn("-10123456.7890123456789", "")
		fn("-101234567.890123456789", "")
		fn("-1012345678.90123456789", "")
		fn("-10123456789.0123456789", "")
		fn("-101234567890.123456789", "")
		fn("-1012345678901.23456789", "")
		fn("-10123456789012.3456789", "")
		fn("-101234567890123.456789", "")
		fn("-1012345678901234.56789", "")
		fn("-10123456789012345.6789", "")
		fn("-101234567890123456.789", "")
		fn("-1012345678901234567.89", "")
		fn("-10123456789012345678.9", "")
		fn("-10123456789012345678.9012345678", "")
		fn("-1.012345678901234567891", "")
		fn("-1.0123456789012345678901", "")

		// 자릿수 한도
		{
			// 정수는 96 자리 까지 ok
			fn("-1"+strings.Repeat("0", DEF_b1_header__max_len__standard-1), "")
			fn("-"+strings.Repeat("9", DEF_b1_header__max_len__standard), "")
			// 소수는 32 자리 까지 ok
			fn("-0."+strings.Repeat("0", DEF_b1_header__max_len__decimal-1)+"1", "")

			// 음수 최소값
			fn("-"+strings.Repeat("9", DEF_b1_header__max_len__standard)+"."+strings.Repeat("9", DEF_b1_header__max_len__decimal), "")
		}
	}
}

func Test_encode__sort(_t *testing.T) {
	type T_data struct {
		s_num  string
		bt_num []byte
	}

	arrpt_ori := make([]*T_data, 0, 100)
	arrpt_sort := make([]*T_data, 0, 100)

	fn_input := func(_s_num string) {
		var err error

		pt_bt_sorted := &T_Encoder{}
		pt_bt_sorted.Init()
		pt_bt_sorted.Set__str(_s_num)
		bt_encode, err := pt_bt_sorted.Encode()
		if err != nil {
			_t.Errorf("input - %s | err - %v", _s_num, err)
		}

		pt_data := &T_data{
			s_num:  _s_num,
			bt_num: bt_encode,
		}
		arrpt_ori = append(arrpt_ori, pt_data)
		arrpt_sort = append(arrpt_sort, pt_data)
	}

	fn_print := func() {
		fmt.Printf("-----------------------------------------\n")
		fmt.Printf("%10s - %10s - %s", "orignal", "sort", "[]byte(ori기준)\n")
		for i := 0; i < len(arrpt_ori); i++ {
			fmt.Printf("%10s - %10s - %v\n", arrpt_ori[i].s_num, arrpt_sort[i].s_num, arrpt_ori[i].bt_num)
		}
		fmt.Printf("-----------------------------------------\n")
	}

	fn_error_check := func() {
		is_exist_error := false
		for i := 0; i < len(arrpt_ori); i++ {
			// 같지 않으면 에러 출력
			if arrpt_ori[i] != arrpt_sort[i] {
				if is_exist_error == false {
					_t.Errorf("-----------------------------------------\n")
					_t.Errorf("%10s - %10s - %s", "orignal", "sort", "[]byte(ori기준)\n")
				}
				is_exist_error = true
				_t.Errorf("%10s - %10s %08b\n", arrpt_ori[i].s_num, arrpt_sort[i].s_num, arrpt_ori[i].bt_num)
			}
		}
		if is_exist_error == true {
			_t.Errorf("-----------------------------------------\n")
			fn_print()
		}
	}
	fn_input("-" + strings.Repeat("9", DEF_b1_header__max_len__standard) + "." + strings.Repeat("9", DEF_b1_header__max_len__decimal)) // 음수 최소값
	fn_input("-" + strings.Repeat("9", DEF_b1_header__max_len__standard))
	fn_input("-10000")
	fn_input("-9999")
	fn_input("-1.2")
	fn_input("-1.199999")
	fn_input("-1.19")
	fn_input("-1.1899")
	fn_input("-1.189")
	fn_input("-1.1112")
	fn_input("-1.11111")
	fn_input("-1.111109")
	fn_input("-1.1111")
	fn_input("-1.111")
	fn_input("-1.11")
	fn_input("-1.1")
	fn_input("-1.0901")
	fn_input("-1.09")
	fn_input("-1.089")
	fn_input("-1")
	fn_input("-0.9")
	fn_input("-0.1")
	fn_input("-0.01")
	fn_input("-0.001")
	fn_input("-0." + strings.Repeat("0", DEF_b1_header__max_len__decimal-1) + "1") // 음수 최대값
	fn_input("0")
	fn_input("0." + strings.Repeat("0", DEF_b1_header__max_len__decimal-1) + "1") // 양수 최소값
	fn_input("0.001")
	fn_input("0.01")
	fn_input("0.1")
	fn_input("0.9")
	fn_input("1")
	fn_input("1.09")
	fn_input("1.1")
	fn_input("1.11")
	fn_input("1.19")
	fn_input("9")
	fn_input("9999")
	fn_input("10000")
	fn_input(strings.Repeat("9", DEF_b1_header__max_len__standard))
	fn_input(strings.Repeat("9", DEF_b1_header__max_len__standard) + "." + strings.Repeat("9", DEF_b1_header__max_len__decimal)) // 양수 최대값

	sort.SliceStable(arrpt_sort, func(_i, _j int) bool {
		n_cmp := bytes.Compare(arrpt_sort[_i].bt_num, arrpt_sort[_j].bt_num)
		if n_cmp == 1 {
			return false
		}
		return true
	})

	fn_error_check()
}
