package sort

/*
	encoding 시 최대 용량 ( max 값일 경우 )
		양수 : 55 byte (header 1 bt + body 54 bt)
		음수 : 55 byte (header 1 bt + body 54 bt)
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
	DEF_lenHeader   int = 1
	DEF_lenBodyMin  int = 1 // 숫자가 zero (0) 인 경우도 1 byte 사용
	DEF_lenTotalMin int = DEF_lenHeader + DEF_lenBodyMin

	DEF_digitIntegerMax int = 96
	DEF_digitDecimalMax int = 32

	DEF_headerBitMaskSign        byte = 128 // 1000 0000
	DEF_headerBitMaskStandardLen byte = 127 // 0111 1111
	DEF_headerValueSignPlus      byte = DEF_headerBitMaskSign
)
