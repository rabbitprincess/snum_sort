package snum

import (
	"errors"
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
	DEF_dot                     string = "."
	DEF_plus                    string = "+"
	DEF_minus                   string = "-"
	DEF_zero                    string = "0"
	DEF_lenExtendDecimalForCalc int    = 20
	DEF_base10                  int    = 10

	DEF_lenDataMin      int = 1 // 숫자가 zero (0) 인 경우도 1 byte 를 사용 함 - big.int 스팩이 아니라 자체 스팩
	DEF_lenDataMinTotal int = DEF_headerSize + DEF_lenDataMin

	DEF_headerSize               int  = 1
	DEF_headerBitMaskSign        byte = 128 // 1000 0000
	DEF_headerBitMaskStandardLen byte = 127 // 0111 1111
	DEF_headerBitMask4bitHigh    byte = 240 // 1111 0000
	DEF_headerBitMask4bitLow     byte = 15  // 0000 1111
	DEF_headerValueSignPlus      byte = DEF_headerBitMaskSign
	DEF_headerLenInteger         int  = 96
	DEF_headerLenDecimal         int  = 32
)

var (
	ErrNotNumber       error = errors.New("NaN")
	ErrHeaderNotEnough error = errors.New("header not enough")
	ErrDivByZero       error = errors.New("division by zero")
)
