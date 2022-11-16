package sort

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
