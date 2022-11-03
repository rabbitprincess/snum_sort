package sort

import (
	"bytes"
	"fmt"
	"sort"
	"strings"
	"testing"
)

func Test__encode_decode(_t *testing.T) {
	fn := func(input string, expect string) {
		if expect == "" {
			expect = input
		}

		// Byte_encode, Byte_decode
		// string -> bigint -> binary -> bigint -> string
		numSort := NewSnumSort(input)

		enc, err := numSort.Encode()
		if err != nil {
			_t.Error(err)
			return
		}

		err = numSort.Decode(enc)
		if err != nil {
			_t.Error(err)
			return
		}

		recovery, err := numSort.GetStr()
		if err != nil {
			_t.Error(err)
			return
		}

		if expect != recovery {
			_t.Errorf("expect: %s, recovery: %s", expect, recovery)
			return
		}
	}

	// 0
	{
		fn("0", "0")
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
			fn("1"+strings.Repeat("0", DEF_digitIntegerMax-1), "")
			fn(strings.Repeat("9", DEF_digitIntegerMax), "")
			// 소수는 32 자리 까지 ok
			fn("0."+strings.Repeat("0", DEF_digitDecimalMax-1)+"1", "")

			// 양수 최대값
			fn(strings.Repeat("9", DEF_digitIntegerMax)+"."+strings.Repeat("9", DEF_digitDecimalMax), "")
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
			fn("-1"+strings.Repeat("0", DEF_digitIntegerMax-1), "")
			fn("-"+strings.Repeat("9", DEF_digitIntegerMax), "")
			// 소수는 32 자리 까지 ok
			fn("-0."+strings.Repeat("0", DEF_digitDecimalMax-1)+"1", "")

			// 음수 최소값
			fn("-"+strings.Repeat("9", DEF_digitIntegerMax)+"."+strings.Repeat("9", DEF_digitDecimalMax), "")
		}
	}
}

func Test_encode__sort(_t *testing.T) {
	type Input struct {
		sn string
		bt []byte
	}

	oris := make([]*Input, 0, 100)
	sorts := make([]*Input, 0, 100)

	fn_input := func(snum string) {
		sorted := NewSnumSort(snum)
		bt, err := sorted.Encode()
		if err != nil {
			_t.Errorf("input - %s | err - %v", snum, err)
		}

		data := &Input{sn: snum, bt: bt}
		oris = append(oris, data)
		sorts = append(sorts, data)
	}

	print := func() {
		fmt.Printf("-----------------------------------------\n")
		fmt.Printf("%10s - %10s - %s", "orignal", "sort", "[]byte(ori기준)\n")
		for i := 0; i < len(oris); i++ {
			fmt.Printf("%10s - %10s - %v\n", oris[i].sn, sorts[i].sn, oris[i].bt)
		}
		fmt.Printf("-----------------------------------------\n")
	}

	check := func() {
		isExistErr := false
		for i := 0; i < len(oris); i++ {
			// 같지 않으면 에러 출력
			if oris[i] != sorts[i] {
				if isExistErr == false {
					_t.Errorf("-----------------------------------------\n")
					_t.Errorf("%10s - %10s - %s", "orignal", "sort", "[]byte(ori기준)\n")
				}
				isExistErr = true
				_t.Errorf("%10s - %10s %08b\n", oris[i].sn, sorts[i].sn, oris[i].bt)
			}
		}
		if isExistErr == true {
			_t.Errorf("-----------------------------------------\n")
			print()
		}
	}
	fn_input("-" + strings.Repeat("9", DEF_digitIntegerMax) + "." + strings.Repeat("9", DEF_digitDecimalMax)) // 음수 최소값
	fn_input("-" + strings.Repeat("9", DEF_digitIntegerMax))
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
	fn_input("-0." + strings.Repeat("0", DEF_digitDecimalMax-1) + "1") // 음수 최대값
	fn_input("0")
	fn_input("0." + strings.Repeat("0", DEF_digitDecimalMax-1) + "1") // 양수 최소값
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
	fn_input(strings.Repeat("9", DEF_digitIntegerMax))
	fn_input(strings.Repeat("9", DEF_digitIntegerMax) + "." + strings.Repeat("9", DEF_digitDecimalMax)) // 양수 최대값

	sort.SliceStable(sorts, func(i, j int) bool {
		cmp := bytes.Compare(sorts[i].bt, sorts[j].bt)
		if cmp == 1 {
			return false
		}
		return true
	})

	check()
	// print()
}
