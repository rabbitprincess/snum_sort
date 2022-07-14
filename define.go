package snum

import (
	"errors"
	"regexp"
)

/*
	byte encoding 시 가능한 용량 한도
		1. bt__sorted
			1-1. 양수 : 65 byte (header 1 bt + body 64 bt)
			1-2. 음수 : 66 byte (header 1 bt + body 64 bt + 0xFF 1 bt)

		2. bt__unsorted
			2-1. 양수 : 55 byte (header 1 bt + body 54 bt)
			2-2. 음수 : 55 byte (header 1 bt + body 54 bt)
*/

/*

	buf[0]
		- 헤더
		- 255 = 정수 자릿수가 96 자리인 양수 ( 1e95 <= x < 1e96 )
		- 254 = 정수 자릿수가 95 자리인 양수 ( 1e94 <= x < 1e95 )
		- ..... ( 양의 정수 )
		- 160 = 정수 자릿수가 1 자리인 양수 ( 1 <= x < 10 )
		- 159 = 정수 자릿수가 -1 자리인 양수 ( 0.1 <= x < 1 )
		- ..... ( 양의 소수 )
		- 128 = 정수 자릿수가 -32 자리인 양수 ( 0 <= x < 1e-31 ) - !! 0 포함 !!

		- 127 = 정수 자릿수가 -32 자리인 음수 ( -1e-31 < x <= -1e-32 )
		- ..... ( 음의 소수 )
		- 96 = 정수 자릿수가 -1 자리인 음수 ( -1 < x <= -0.1 )
		- 95 = 정수 자릿수가 1 자리인 음수 ( -10 < x <= -1 )
		- ..... ( 음의 정수 )
		- 1 = 정수 자릿수가 95 자리인 음수 ( -1e95 < x <= -1e94 )
		- 0 = 정수 자릿수가 96 자리인 음수 ( -1e96 < x <= -1e95 )

	buf[1:]
		- 정수 + 소수 big.Int 에 담아 2자릿수 당 1바이트 로 압축한 byte array
		- 음수일 경우 보수로 저장

*/

const (
	DEF_s_dot                         string = "."
	DEF_s_plus                        string = "+"
	DEF_s_minus                       string = "-"
	DEF_s_num_zero                    string = "0"
	DEF_n_len_extend_decimal_for_calc int    = 20
	DEF_n_base_10                     int    = 10

	DEF_n_bt__len_min_data  int = 1 // 숫자가 zero (0) 인 경우도 1 byte 를 사용 함 - big.int 스팩이 아니라 자체 스팩
	DEF_n_bt__len_min_total int = DEF_n_header__size + DEF_n_bt__len_min_data

	DEF_n_header__size                    int  = 1
	DEF_b1_header__bit_mask__sign         byte = 128 // 1000 0000
	DEF_b1_header__bit_mask__standard_len byte = 127 // 0111 1111
	DEF_b1_header__bit_mask__high_4_bit   byte = 240 // 1111 0000
	DEF_b1_header__bit_mask__low_4_bit    byte = 15  // 0000 1111
	DEF_b1_header__value__sign__plus      byte = DEF_b1_header__bit_mask__sign
	DEF_b1_header__max_len__standard      int  = 96
	DEF_b1_header__max_len__decimal       int  = 32
)

var (
	Err_not_a_number      error = errors.New("NaN")
	Err_header_not_enongh error = errors.New("header not enough")
	Err_div_by_zero       error = errors.New("division by zero")
)

var Fn_is_digit = regexp.MustCompile(`(^[+-]?\d{0,96}(\.\d{1,32})?$)`).MatchString // 정수 96 자리 , 소수점 32 자리 제한
