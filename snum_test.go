package snum

import (
	"fmt"
	"strings"
	"testing"
)

func Test_trim(_t *testing.T) {
	fn := func(_s_num__input, _s_num__expected string, _n_len_int, _n_len_decimal int) {
		pt_snum := &Snum{}
		pt_snum.Init(_n_len_int, _n_len_decimal)
		pt_snum.SetStr(_s_num__input)
		pt_snum.TrimDigit()

		s_num_result := pt_snum.String()

		if s_num_result != _s_num__expected {
			_t.Errorf("invalid trim | result : [%s] | expected : [%s]", s_num_result, _s_num__expected)
		}
		fmt.Println("before :", _s_num__input)
		fmt.Println("after :", s_num_result)
		fmt.Println()
	}

	fn("1", "1", 1, 1)
	fn("10", "0", 1, 1)
	fn("0.1", "0.1", 1, 1)
	fn("0.01", "0", 1, 1)
	fn("0.11", "0.1", 1, 1)
	fn("123456789.987654321", "9.9", 1, 1)
	fn("123456789.987654321", "9.98", 1, 2)
	fn("123456789.987654321", "89.987", 2, 3)
	fn("123456789.987654321", "89.987", 2, 3)

	// default
	fn("-1"+strings.Repeat("0", DEF_headerLenInteger-1), "-1"+strings.Repeat("0", DEF_headerLenInteger-1), 0, 0)
	fn("-"+strings.Repeat("9", DEF_headerLenInteger), "-"+strings.Repeat("9", DEF_headerLenInteger), 0, 0)
	fn("-0."+strings.Repeat("0", DEF_headerLenDecimal-1)+"1", "-0."+strings.Repeat("0", DEF_headerLenDecimal-1)+"1", 0, 0)
	fn("-"+strings.Repeat("9", DEF_headerLenInteger)+"."+strings.Repeat("9", DEF_headerLenDecimal), "-"+strings.Repeat("9", DEF_headerLenInteger)+"."+strings.Repeat("9", DEF_headerLenDecimal), 0, 0)
	fn(strings.Repeat("9", DEF_headerLenInteger)+"."+strings.Repeat("8", DEF_headerLenDecimal), strings.Repeat("9", DEF_headerLenInteger)+"."+strings.Repeat("8", DEF_headerLenDecimal), 0, 0)

	// 초과
	fn(strings.Repeat("9", DEF_headerLenInteger-10)+"."+strings.Repeat("8", DEF_headerLenDecimal+10), strings.Repeat("9", DEF_headerLenInteger-10)+"."+strings.Repeat("8", DEF_headerLenDecimal), 0, 0)
	fn(strings.Repeat("9", DEF_headerLenInteger+10)+"."+strings.Repeat("8", DEF_headerLenDecimal-10), strings.Repeat("9", DEF_headerLenInteger)+"."+strings.Repeat("8", DEF_headerLenDecimal-10), 0, 0)
}

func Test_set_get(_t *testing.T) {
	fn := func(_s_num__input string, _s_num__expected string) {
		// 기대값 설정
		{
			// 기대값 입력이 비어있으면 입력값과 동일한 것이 정답인 것으로 인정한다.
			if _s_num__expected == "" {
				_s_num__expected = _s_num__input
			}
		}

		// String_set, String_get
		// string -> bigint -> string
		pt_snum := &Snum{}
		pt_snum.Init(0, 0)
		{
			err := pt_snum.SetStr(_s_num__input)
			if err != nil {
				_t.Errorf("String__set\n%s:%v\n%s:%v\n", "_s_num__input :", _s_num__input, "err", err)
			}

			s_num__recovery := pt_snum.String()
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
			fn("1"+strings.Repeat("0", DEF_headerLenInteger-1), "")
			fn(strings.Repeat("9", DEF_headerLenInteger), "")
			fn("0."+strings.Repeat("0", DEF_headerLenDecimal-1)+"1", "")
			fn(strings.Repeat("9", DEF_headerLenInteger)+"."+strings.Repeat("9", DEF_headerLenDecimal), "")
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
			fn("10000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000", "")

			fn("-1"+strings.Repeat("0", DEF_headerLenInteger-1), "")
			fn("-"+strings.Repeat("9", DEF_headerLenInteger), "")
			fn("-0."+strings.Repeat("0", DEF_headerLenDecimal-1)+"1", "")
			fn("-"+strings.Repeat("9", DEF_headerLenInteger)+"."+strings.Repeat("9", DEF_headerLenDecimal), "")
			fn(strings.Repeat("9", DEF_headerLenInteger+10)+"."+strings.Repeat("8", DEF_headerLenDecimal), "") // 양수 최대값
		}
	}
}

//------------------------------------------------------------------------------------------------------//
//

func Test_copy(_t *testing.T) {
	fn := func(_sn_input string) {
		pt_ori := &Snum{}
		pt_ori.Init(0, 0)
		err := pt_ori.SetStr(_sn_input)
		if err != nil {
			_t.Errorf("input string set error | err - %v", err)
			return
		}
		sn_ori, err := pt_ori.GetStr()
		if err != nil {
			_t.Errorf("string get error | err - %v", err)
			return
		}

		pt_copy := &Snum{}
		pt_copy.Init(0, 0)
		pt_copy.decimal.Copy(pt_ori.decimal)
		sn_copy, err := pt_copy.GetStr()
		if err != nil {
			_t.Errorf("string get error | err - %v", err)
			return
		}
		if sn_ori != sn_copy {
			_t.Errorf("%s copy error | expected : %s, result : %v\n", _sn_input, sn_ori, sn_copy)
		}
	}

	fn("0")
	fn("1")
	fn("12")
	fn("123")
	fn("1234")
	fn("123.4")
	fn("12.34")
	fn("1.234")
	fn("0.1234")
	fn("0.01234")
	fn("0.001234")
	fn("0.0001234")

	fn("-0")
	fn("-1")
	fn("-12")
	fn("-123")
	fn("-1234")
	fn("-123.4")
	fn("-12.34")
	fn("-1.234")
	fn("-0.1234")
	fn("-0.01234")
	fn("-0.001234")
	fn("-0.0001234")

	fn("-0." + strings.Repeat("0", DEF_headerLenDecimal-1) + "1")                                         // 음수 최대값
	fn("0." + strings.Repeat("0", DEF_headerLenDecimal-1) + "1")                                          // 양수 최소값
	fn("-" + strings.Repeat("9", DEF_headerLenInteger) + "." + strings.Repeat("9", DEF_headerLenDecimal)) // 음수 최소값
	fn(strings.Repeat("9", DEF_headerLenInteger) + "." + strings.Repeat("9", DEF_headerLenDecimal))       // 양수 최대값
}

func Test_calc(_t *testing.T) {
	type T_Input struct {
		sn_input_1       string
		sn_input_2       string
		sn_expected__add string
		sn_expected__sub string
		sn_expected__mul string
		sn_expected__div string
	}

	fn := func(_t_input T_Input) {
		arrs_calc := []string{"add", "sub", "mul", "div"}
		var s_op string
		var err error
		var sn_result string
		var sn_expect string
		for _, s_calc := range arrs_calc {
			switch s_calc {
			case "add":
				sn_result, err = AddStr(_t_input.sn_input_1, _t_input.sn_input_2)
				s_op = "+"
				sn_expect = _t_input.sn_expected__add
			case "sub":
				sn_result, err = SubStr(_t_input.sn_input_1, _t_input.sn_input_2)
				s_op = "-"
				sn_expect = _t_input.sn_expected__sub
			case "mul":
				sn_result, err = MulStr(_t_input.sn_input_1, _t_input.sn_input_2)
				s_op = "*"
				sn_expect = _t_input.sn_expected__mul
			case "div":
				sn_result, err = DivStr(_t_input.sn_input_1, _t_input.sn_input_2)
				s_op = "/"
				sn_expect = _t_input.sn_expected__div
			}
			if err != nil {
				_t.Errorf("err - %v", err)
			}
			if sn_expect != sn_result {
				_t.Errorf("\n%s %s %s\n%10s - %s\n%10s - %s\n", _t_input.sn_input_1, s_op, _t_input.sn_input_2, "result", sn_result, "expect", sn_expect)
			}
		}
	}

	fn(T_Input{"111", "111", "222", "0", "12321", "1"})
	fn(T_Input{"111", "-111", "0", "222", "-12321", "-1"})
	fn(T_Input{"+111", "-111", "0", "222", "-12321", "-1"})
	fn(T_Input{"-111", "111", "0", "-222", "-12321", "-1"})
	fn(T_Input{"222", "-111", "111", "333", "-24642", "-2"})
	fn(T_Input{"-222", "111", "-111", "-333", "-24642", "-2"})
	fn(T_Input{"-222", "-111", "-333", "-111", "24642", "2"})

	fn(T_Input{"0", "0", "0", "0", "0", "NaN34"})
	fn(T_Input{"50", "0", "50", "50", "0", "Infinity"})

	fn(T_Input{"111", "111", "222", "0", "12321", "1"})
	fn(T_Input{"1", "9", "10", "-8", "9", "0.11111111111111111111"})
	fn(T_Input{"12", "98", "110", "-86", "1176", "0.12244897959183673469"})
	fn(T_Input{"12", "987", "999", "-975", "11844", "0.0121580547112462006"})
	fn(T_Input{"12", "9876", "9888", "-9864", "118512", "0.00121506682867557715"})
	fn(T_Input{"987", "123", "1110", "864", "121401", "8.02439024390243902439"})
	fn(T_Input{"987", "12", "999", "975", "11844", "82.25"})
	fn(T_Input{"987", "1", "988", "986", "987", "987"})
	fn(T_Input{"123.123", "987.987", "1111.11", "-864.864", "121643.923401", "0.12462006079027355623"})
	fn(T_Input{"123", "987.987", "1110.987", "-864.987", "121522.401", "0.12449556522504850772"})
	fn(T_Input{"123.123", "987", "1110.123", "-863.877", "121522.401", "0.12474468085106382978"})
	fn(T_Input{"987", "123.123", "1110.123", "863.877", "121522.401", "8.01637387003240661777"})
	fn(T_Input{"987.987", "123", "1110.987", "864.987", "121522.401", "8.03241463414634146341"})
	fn(T_Input{"123.987", "123", "246.987", "0.987", "15250.401", "1.00802439024390243902"})
	fn(T_Input{"123", "0.987", "123.987", "122.013", "121.401", "124.620060790273556231"})
	fn(T_Input{"123", "0.0987", "123.0987", "122.9013", "12.1401", "1246.20060790273556231003"})
	fn(T_Input{"123", "0.00987", "123.00987", "122.99013", "1.21401", "12462.0060790273556231003"})
	fn(T_Input{"123", "123.00987", "246.00987", "-0.00987", "15130.21401", "0.99991976253612819849"})
	fn(T_Input{"0.00123", "0.00987", "0.0111", "-0.00864", "0.0000121401", "0.12462006079027355623"})
	fn(T_Input{"0.00123", ".00987", "0.0111", "-0.00864", "0.0000121401", "0.12462006079027355623"})
	fn(T_Input{".00123", "0.00987", "0.0111", "-0.00864", "0.0000121401", "0.12462006079027355623"})
	fn(T_Input{".00123", ".00987", "0.0111", "-0.00864", "0.0000121401", "0.12462006079027355623"})
	fn(T_Input{"0.1", "9", "9.1", "-8.9", "0.9", "0.01111111111111111111"})
	fn(T_Input{"0.01", "9", "9.01", "-8.99", "0.09", "0.00111111111111111111"})
	fn(T_Input{"0.001", "9", "9.001", "-8.999", "0.009", "0.00011111111111111111"})
	fn(T_Input{"0.123", "123", "123.123", "-122.877", "15.129", "0.001"})
	fn(T_Input{"123", "0.123", "123.123", "122.877", "15.129", "1000"})

}

func Test_abs_neg(_t *testing.T) {
	type T_Input struct {
		sn_input         string
		sn_expected__abs string
		sn_expected__neg string
	}

	fn := func(_t_input T_Input) {
		s_output, err := AbsStr(_t_input.sn_input)
		if err != nil {
			_t.Errorf("input string set error | err - %v", err)
			return
		}
		if s_output != _t_input.sn_expected__abs {
			_t.Errorf("%s abs error | expected : %v, result : %v\n", _t_input.sn_input, _t_input.sn_expected__abs, s_output)
		}

		s_output, err = NegStr(_t_input.sn_input)
		if err != nil {
			_t.Errorf("input string set error | err - %v", err)
			return
		}
		if s_output != _t_input.sn_expected__neg {
			_t.Errorf("%s neg error | expected : %v, result : %v\n", _t_input.sn_input, _t_input.sn_expected__neg, s_output)
		}
	}

	fn(T_Input{"1", "1", "-1"})
	fn(T_Input{"-1", "1", "1"})
	fn(T_Input{"0", "0", "0"})
	fn(T_Input{"-0", "0", "0"})
	fn(T_Input{"0.123456789", "0.123456789", "-0.123456789"})
	fn(T_Input{"-0.123456789", "0.123456789", "0.123456789"})
	fn(T_Input{"123456789.123456789", "123456789.123456789", "-123456789.123456789"})
	fn(T_Input{"-123456789.123456789", "123456789.123456789", "123456789.123456789"})
}

func Test_cmp(_t *testing.T) {
	type T_Input struct {
		sn_input_1  string
		sn_input_2  string
		sn_expected int
	}

	fn := func(_t_input T_Input) {
		n_cmp, err := CmpStr(_t_input.sn_input_1, _t_input.sn_input_2)
		if err != nil {
			_t.Errorf("input string set error | err - %v", err)
			return
		}
		if n_cmp != _t_input.sn_expected {
			_t.Errorf("%s - %s cmp error | expected : %v, result : %v\n", _t_input.sn_input_1, _t_input.sn_input_2, _t_input.sn_expected, n_cmp)
		}
	}

	// positive
	{
		fn(T_Input{"0.123", "123", -1})
		fn(T_Input{"1.23", "23", -1})
		fn(T_Input{"12.3", "123", -1})

		fn(T_Input{"123", "0.123", 1})
		fn(T_Input{"123", "1.23", 1})
		fn(T_Input{"123", "12.3", 1})

		fn(T_Input{"123", "123", 0})
		fn(T_Input{"1230", "123", 1})
		fn(T_Input{"12300", "123", 1})
		fn(T_Input{"123", "1230", -1})
		fn(T_Input{"123", "12300", -1})
	}

	// negative
	{
		fn(T_Input{"-0.123", "-123", 1})
		fn(T_Input{"-1.23", "-23", 1})
		fn(T_Input{"-12.3", "-123", 1})

		fn(T_Input{"-123", "-0.123", -1})
		fn(T_Input{"-123", "-1.23", -1})
		fn(T_Input{"-123", "-12.3", -1})

		fn(T_Input{"-123", "-123", 0})
		fn(T_Input{"-1230", "-123", -1})
		fn(T_Input{"-12300", "-123", -1})
		fn(T_Input{"-123", "-1230", 1})
		fn(T_Input{"-123", "-12300", 1})
	}

	// special case
	{
		fn(T_Input{"0", "-0", 0})
		fn(T_Input{"0", "-0.0", 0})
		fn(T_Input{"0.1", "-0.1", 1})
		fn(T_Input{"0.01", "-0.1", 1})
		fn(T_Input{"0.001", "-0.1", 1})

		fn(T_Input{"123.321", "123.3211", -1})
		fn(T_Input{"123.321", "1233.21", -1})
		fn(T_Input{"-123.321", "-123.3211", 1})
		fn(T_Input{"-123.321", "-1233.21", 1})
	}
}

func Test_round(_t *testing.T) {
	type T_Input struct {
		sn_input                string
		n_round_cnt             int
		sn_expected__round      string
		sn_expected__round_down string
		sn_expected__round_up   string
	}

	fn := func(_t_input T_Input) {
		pt_input := &Snum{}
		pt_input.Init(0, 0)

		var sn_result__round, sn_result__round_down, sn_result__round_up string
		{
			err := pt_input.SetStr(_t_input.sn_input)
			if err != nil {
				_t.Errorf("input string set error | err - %v", err)
				return
			}
			pt_input_round := pt_input.Copy()
			pt_input_round.Round(_t_input.n_round_cnt)
			sn_result__round = pt_input_round.String()

			pt_input_round_up := pt_input.Copy()
			pt_input_round_up.RoundUp(_t_input.n_round_cnt)
			sn_result__round_up = pt_input_round_up.String()

			pt_input_round_down := pt_input.Copy()
			pt_input_round_down.RoundDown(_t_input.n_round_cnt)
			sn_result__round_down = pt_input_round_down.String()
		}
		if sn_result__round != _t_input.sn_expected__round {
			_t.Errorf("%s - %d round error | expected : %v, result : %v\n", _t_input.sn_input, _t_input.n_round_cnt, _t_input.sn_expected__round, sn_result__round)
		}
		if sn_result__round_up != _t_input.sn_expected__round_up {
			_t.Errorf("%s - %d round_up error | expected : %v, result : %v\n", _t_input.sn_input, _t_input.n_round_cnt, _t_input.sn_expected__round_up, sn_result__round_up)
		}
		if sn_result__round_down != _t_input.sn_expected__round_down {
			_t.Errorf("%s - %d round_down error | expected : %v, result : %v\n", _t_input.sn_input, _t_input.n_round_cnt, _t_input.sn_expected__round_down, sn_result__round_down)
		}
	}

	{
		fn(T_Input{"123456789.123456789", 1, "100000000", "100000000", "200000000"})
		fn(T_Input{"123456789.123456789", 2, "120000000", "120000000", "130000000"})
		fn(T_Input{"123456789.123456789", 3, "123000000", "123000000", "124000000"})
		fn(T_Input{"123456789.123456789", 4, "123500000", "123400000", "123500000"})
		fn(T_Input{"123456789.123456789", 5, "123460000", "123450000", "123460000"})
		fn(T_Input{"123456789.123456789", 6, "123457000", "123456000", "123457000"})
		fn(T_Input{"123456789.123456789", 7, "123456800", "123456700", "123456800"})
		fn(T_Input{"123456789.123456789", 8, "123456790", "123456780", "123456790"})
		fn(T_Input{"123456789.123456789", 9, "123456789", "123456789", "123456790"})
		fn(T_Input{"123456789.123456789", 10, "123456789.1", "123456789.1", "123456789.2"})
		fn(T_Input{"123456789.123456789", 11, "123456789.12", "123456789.12", "123456789.13"})
		fn(T_Input{"123456789.123456789", 12, "123456789.123", "123456789.123", "123456789.124"})
		fn(T_Input{"123456789.123456789", 13, "123456789.1235", "123456789.1234", "123456789.1235"})
		fn(T_Input{"123456789.123456789", 14, "123456789.12346", "123456789.12345", "123456789.12346"})
		fn(T_Input{"123456789.123456789", 15, "123456789.123457", "123456789.123456", "123456789.123457"})
		fn(T_Input{"123456789.123456789", 16, "123456789.1234568", "123456789.1234567", "123456789.1234568"})
		fn(T_Input{"123456789.123456789", 17, "123456789.12345679", "123456789.12345678", "123456789.12345679"})
		fn(T_Input{"123456789.123456789", 18, "123456789.123456789", "123456789.123456789", "123456789.123456789"})

		fn(T_Input{"0.123456789", 4, "0.1235", "0.1234", "0.1235"})
		fn(T_Input{"0.023456789", 4, "0.02346", "0.02345", "0.02346"})
		fn(T_Input{"0.003456789", 4, "0.003457", "0.003456", "0.003457"})
		fn(T_Input{"0.000456789", 4, "0.0004568", "0.0004567", "0.0004568"})
		fn(T_Input{"0.000056789", 4, "0.00005679", "0.00005678", "0.00005679"})
		fn(T_Input{"0.000006789", 4, "0.000006789", "0.000006789", "0.000006789"})
		fn(T_Input{"0.000000789", 4, "0.000000789", "0.000000789", "0.000000789"})
		fn(T_Input{"0.000000089", 4, "0.000000089", "0.000000089", "0.000000089"})
		fn(T_Input{"0.000000009", 4, "0.000000009", "0.000000009", "0.000000009"})
	}
}

func Test_pow(_t *testing.T) {
	sn_ret, err := Pow("2", 10)
	if err != nil {
		_t.Fatal(err)
	}
	fmt.Println(sn_ret)
}

func Test_scale_limit(_t *testing.T) {
	fn_print := func(_n_precision int, _s_num string) {
		pt_snum := &Snum{}
		pt_snum.Init(_n_precision, 0)

		err := pt_snum.SetStr(_s_num)
		if err != nil {
			_t.Fatal(err)
		}

		fmt.Println("입력 :", _s_num)
		fmt.Println("출력 :", pt_snum.String())
		fmt.Printf("소수부 길이 : %v | 정수부 길이 : %v | 정밀도 : %v \n", pt_snum.decimal.Scale(), pt_snum.decimal.Precision()-pt_snum.decimal.Scale(), pt_snum.decimal.Precision())
		fmt.Println()
		// pt_snum.pt_decimal.Context.MinScale

	}

	fn_print(128, "0")
	fn_print(128, "-0."+strings.Repeat("0", DEF_headerLenDecimal-1)+"1")
	fn_print(128, "0."+strings.Repeat("0", DEF_headerLenDecimal-1)+"1")
	fn_print(128, "0.1")
	fn_print(128, "0.0000001")
	fn_print(128, "0.1111111")
	fn_print(128, "1000000")
	fn_print(128, "1111111")

	fn_print(128, "1000000.0000001")
	fn_print(128, "1111111.1111111")

	fn_print(128, "1111111.1111111")
	fn_print(128, strings.Repeat("9", DEF_headerLenInteger)+"."+strings.Repeat("9", DEF_headerLenDecimal))

}
