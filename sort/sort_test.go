package sort

import (
	"bytes"
	"encoding/json"
	"fmt"
	"sort"
	"strings"
	"testing"

	"github.com/gokch/snum_sort/snum"
	"github.com/stretchr/testify/require"
)

// Byte_encode, Byte_decode
// string -> bigint -> binary -> bigint -> string
func TestSort_encode_decode(t *testing.T) {
	fn := func(input string) {
		numSort := New(input)

		enc := numSort.Encode()
		err := numSort.Decode(enc)
		require.NoError(t, err)

		recovery := numSort.GetStr()

		require.Equal(t, input, recovery)
	}

	// 0
	fn("0")

	// 1 이상
	fn("1")
	fn("12")
	fn("123")
	fn("123.1")
	fn("123.12")
	fn("123.123")
	fn("12.3123")
	fn("1.23123")
	fn("101")
	fn("2020")
	fn("30303030")
	fn("404040404")
	fn("102030405060708090")

	// 1 미만
	fn("0.1")
	fn("0.123123")
	fn("0.0123123")
	fn("0.00123123")
	fn("0.000123123")
	fn("0.000000000000000000000000000001")
	fn("0.0000000000000000000000000000001")
	fn("0.00000000000000000000000000000001")

	// 소수
	fn("1.01")
	fn("1.001")
	fn("1.000000000000000000000000000001")
	fn("1.0000000000000000000000000000001")
	fn("1.00000000000000000000000000000001")

	fn("1.0123456789012345678")
	fn("1.01234567890123456789")
	fn("10.1234567890123456789")
	fn("101.234567890123456789")
	fn("1012.34567890123456789")
	fn("10123.4567890123456789")
	fn("101234.567890123456789")
	fn("1012345.67890123456789")
	fn("10123456.7890123456789")
	fn("101234567.890123456789")
	fn("1012345678.90123456789")
	fn("10123456789.0123456789")
	fn("101234567890.123456789")
	fn("1012345678901.23456789")
	fn("10123456789012.3456789")
	fn("101234567890123.456789")
	fn("1012345678901234.56789")
	fn("10123456789012345.6789")
	fn("101234567890123456.789")
	fn("1012345678901234567.89")
	fn("10123456789012345678.9")
	fn("10123456789012345678.9012345678")
	fn("1.012345678901234567891")
	fn("1.0123456789012345678901")

	// 자릿수 한도
	{
		// 정수는 96 자리 까지 ok
		fn("1" + strings.Repeat("0", DEF_digitIntegerMax-1))
		fn(strings.Repeat("9", DEF_digitIntegerMax))
		// 소수는 32 자리 까지 ok
		fn("0." + strings.Repeat("0", DEF_digitDecimalMax-1) + "1")

		// 양수 최대값
		fn(strings.Repeat("9", DEF_digitIntegerMax) + "." + strings.Repeat("9", DEF_digitDecimalMax))
	}

	// -1 이하
	fn("-1")
	fn("-12")
	fn("-123")
	fn("-123.1")
	fn("-123.12")
	fn("-123.123")
	fn("-12.3123")
	fn("-1.23123")
	fn("-1")
	fn("-1.1")
	fn("-1.01")
	fn("-1.001")
	fn("-1.0001")
	fn("-1.00001")
	fn("-1.000001")
	fn("-1.0000001")
	fn("-101")
	fn("-2020")
	fn("-30303030")
	fn("-404040404")

	// -1 초과
	fn("-0.1")
	fn("-0.123123")
	fn("-0.0123123")
	fn("-0.00123123")
	fn("-0.000123123")
	fn("-0.000000000000000000000000000001")
	fn("-0.0000000000000000000000000000001")
	fn("-0.00000000000000000000000000000001")

	// 소수
	fn("-1.01")
	fn("-1.001")
	fn("-1.000000000000000000000000000001")
	fn("-1.0000000000000000000000000000001")
	fn("-1.00000000000000000000000000000001")
	fn("-1.0123456789012345678")
	fn("-1.01234567890123456789")
	fn("-10.1234567890123456789")
	fn("-101.234567890123456789")
	fn("-1012.34567890123456789")
	fn("-10123.4567890123456789")
	fn("-101234.567890123456789")
	fn("-1012345.67890123456789")
	fn("-10123456.7890123456789")
	fn("-101234567.890123456789")
	fn("-1012345678.90123456789")
	fn("-10123456789.0123456789")
	fn("-101234567890.123456789")
	fn("-1012345678901.23456789")
	fn("-10123456789012.3456789")
	fn("-101234567890123.456789")
	fn("-1012345678901234.56789")
	fn("-10123456789012345.6789")
	fn("-101234567890123456.789")
	fn("-1012345678901234567.89")
	fn("-10123456789012345678.9")
	fn("-10123456789012345678.9012345678")
	fn("-1.012345678901234567891")
	fn("-1.0123456789012345678901")

	// 자릿수 한도
	{
		// 정수는 96 자리 까지 ok
		fn("-1" + strings.Repeat("0", DEF_digitIntegerMax-1))
		fn("-" + strings.Repeat("9", DEF_digitIntegerMax))
		// 소수는 32 자리 까지 ok
		fn("-0." + strings.Repeat("0", DEF_digitDecimalMax-1) + "1")

		// 음수 최소값
		fn("-" + strings.Repeat("9", DEF_digitIntegerMax) + "." + strings.Repeat("9", DEF_digitDecimalMax))
	}

}

func Test_sort(t *testing.T) {
	type Input struct {
		snum   string
		encode []byte
	}

	oris := make([]*Input, 0, 100)
	sorts := make([]*Input, 0, 100)

	input := func(snum string) {
		sorted := New(snum)
		bt := sorted.Encode()

		data := &Input{snum: snum, encode: bt}
		oris = append(oris, data)
		sorts = append(sorts, data)
	}

	sortAndCheck := func() {
		// sort checks
		sort.SliceStable(sorts, func(i, j int) bool {
			cmp := bytes.Compare(sorts[i].encode, sorts[j].encode)
			if cmp < 0 {
				return true
			}
			return false
		})
		// cmp ori and sorts
		require.EqualValues(t, oris, sorts)
	}
	input("-" + strings.Repeat("9", DEF_digitIntegerMax) + "." + strings.Repeat("9", DEF_digitDecimalMax)) // 음수 최소값
	input("-" + strings.Repeat("9", DEF_digitIntegerMax))
	input("-10000")
	input("-9999")
	input("-1.2")
	input("-1.199999")
	input("-1.19")
	input("-1.1899")
	input("-1.189")
	input("-1.1119")
	input("-1.1112")
	input("-1.11111")
	input("-1.111109")
	input("-1.1111")
	input("-1.111")
	input("-1.11")
	input("-1.1")
	input("-1.0901")
	input("-1.09")
	input("-1.089")
	input("-1.0001")
	input("-1.0000000000000000001")
	input("-1")
	input("-0.9")
	input("-0.1")
	input("-0.01")
	input("-0.001")
	input("-0." + strings.Repeat("0", DEF_digitDecimalMax-1) + "1") // 음수 최대값
	input("0")
	input("0." + strings.Repeat("0", DEF_digitDecimalMax-1) + "1") // 양수 최소값
	input("0.001")
	input("0.01")
	input("0.1")
	input("0.9")
	input("1")
	input("1.0000000000000000000000000000001")
	input("1.000000000000000000000000000001")
	input("1.09")
	input("1.1")
	input("1.11")
	input("1.19")
	input("9")
	input("9999")
	input("10000")
	input(strings.Repeat("9", DEF_digitIntegerMax))
	input(strings.Repeat("9", DEF_digitIntegerMax) + "." + strings.Repeat("9", DEF_digitDecimalMax)) // 양수 최대값

	sortAndCheck()
}

func TestSort_header(t *testing.T) {
	fn := func(lenRaw, lenDecimal int) {
		header := encodeHeader(lenRaw, lenDecimal)
		lenDecimalNew := decodeHeader(lenRaw, header)
		require.EqualValues(t, lenDecimal, lenDecimalNew)
	}

	fn(1, 32)   // "0"
	fn(1, 32)   // "0.000000000000000000000000000000001"
	fn(1, 31)   // "0.00000000000000000000000000000001"
	fn(1, 30)   // "0.0000000000000000000000000000001"
	fn(33, 32)  // "1.00000000000000000000000000000001"
	fn(1, 0)    // "1"
	fn(2, 0)    // "10"
	fn(2, 1)    // "10.1"
	fn(96, 0)   // 96 자리 정수
	fn(128, 32) // 96 자리 정수 + 32 자리 소수
}

func TestSort_json(t *testing.T) {
	fn_test := func(sn string) {
		snumSort := New(snum.New(sn))

		enc, err := json.MarshalIndent(&snumSort, "", "\t")
		fmt.Println(sn, string(enc))
		require.NoError(t, err)

		snumSortNew := New(snum.New(0))
		err = json.Unmarshal(enc, &snumSortNew)
		require.NoError(t, err)

		require.Equal(t, snumSort.String(), snumSortNew.String())
	}

	fn_test("11")
	fn_test("10.00000000000000000000000000000001")
	fn_test("10")
	fn_test("1.2")
	fn_test("1.1")
	fn_test("1")
	fn_test("0.1")
	fn_test("0.123123")
	fn_test("0.0123123")
	fn_test("0.00123123")
	fn_test("0.000123123")
	fn_test("0.000000000000000000000000000001")
	fn_test("0.0000000000000000000000000000001")
	fn_test("0.00000000000000000000000000099999")
	fn_test("0.00000000000000000000000000009999")
	fn_test("0.00000000000000000000000000000999")
	fn_test("0.00000000000000000000000000000099")
	fn_test("0.00000000000000000000000000000009")
	fn_test("0.00000000000000000000000000000001")
	fn_test("0")
	fn_test("-0.00000000000000000000000000000001")
	fn_test("-0.00000000000000000000000000000009")
	fn_test("-0.00000000000000000000000000000099")
	fn_test("-0.00000000000000000000000000000999")
	fn_test("-0.00000000000000000000000000009999")
	fn_test("-0.00000000000000000000000000099999")
	fn_test("-0.0000000000000000000000000000001")
	fn_test("-0.000000000000000000000000000001")
	fn_test("-0.000123123")
	fn_test("-0.00123123")
	fn_test("-0.0123123")
	fn_test("-0.1")
	fn_test("-0.123123")
	fn_test("-1")
	fn_test("-1.1")
	fn_test("-1.11")
	fn_test("-1.111")
	fn_test("-1.1111")
	fn_test("-1.2")
	fn_test("-10")
	fn_test("-10.00000000000000000000000000000011")
	fn_test("-10.00000000000000000000000000000010")
	fn_test("-10.00000000000000000000000000000001")
	fn_test("-11")
}
