package sort

const (
	DEF_lenHeader   int = 1
	DEF_lenTotalMin int = DEF_lenHeader + 1

	DEF_digitIntegerMax int = 96
	DEF_digitDecimalMax int = 32

	DEF_headerBitMaskSign        byte = 0x80 // 1000 0000
	DEF_headerBitMaskStandardLen byte = 0x7F // 0111 1111

	DEF_asciiZeroUnderOne byte = 0x30 - 1 // 0011 0000
)
