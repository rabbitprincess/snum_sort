package sort

import (
	"errors"
	"math/big"
)

func (t *SnumSort) Encode() (enc []byte) {
	bigRaw, lenDecimal, isMinus := t.Snum.GetRaw()

	raw := bigRaw.String()
	if raw == "0" { // 0 일 경우 전처리
		lenDecimal = DEF_digitDecimalMax
	}

	header := encodeHeader(len(raw), lenDecimal)
	body := encodeBody(raw)
	enc = encodeMinus(isMinus, append([]byte{header}, body...))
	return enc
}

func (t *SnumSort) Decode(enc []byte) error {
	// Do not use any other value
	if len(enc) < DEF_lenTotalMin {
		return errors.New("too short")
	}
	isMinus, enc := decodeMinus(enc)
	raw := decodeBody(enc[1:])
	lenDecimal := decodeHeader(len(raw), enc[0])

	if len(enc) == 2 && enc[1] == 0 { // 0 일 경우 후처리
		lenDecimal = 0
	}
	bigRaw, _ := big.NewInt(0).SetString(raw, 10)

	t.Snum.SetRaw(bigRaw, lenDecimal, isMinus)
	return nil
}

//------------------------------------------------------------------------------------------//
// util

func encodeHeader(lenRaw, lenDecimal int) (header byte) {
	posStartDot := lenRaw - lenDecimal + DEF_digitDecimalMax - 1
	header = byte(posStartDot) | DEF_headerBitMaskSign
	return header
}

func encodeBody(raw string) (body []byte) {
	body = make([]byte, 0, len(raw)/2+1)
	numOri := []byte(raw)
	for i := 0; i < len(raw); i++ {
		b4 := numOri[i] - DEF_asciiZeroUnderOne
		if i%2 == 0 {
			body = append(body, b4<<4)
		} else {
			body[i/2] += b4
		}
	}
	return body
}

func encodeMinus(isMinus bool, standard []byte) (minus []byte) {
	if isMinus == true {
		for i := 0; i < len(standard); i++ {
			standard[i] = ^standard[i]
		}
		if standard[len(standard)-1]<<4 != 0xF0 {
			standard = append(standard, 0xF0) // append last 0xF0
		}
	}
	return standard
}

func decodeMinus(minus []byte) (isMinus bool, standard []byte) {
	if minus[0]&DEF_headerBitMaskSign == 0 {
		isMinus = true
		if minus[len(minus)-1] == 0xF0 {
			minus = minus[:len(minus)-1] // remove last 0xF0
		}
		for i := 0; i < len(minus); i++ {
			minus[i] = ^minus[i]
		}
	}
	return isMinus, minus
}

func decodeBody(body []byte) (raw string) {
	for i := 0; i < len(body); i++ {
		high4bit := body[i] >> 4
		low4bit := body[i] - (high4bit << 4)

		raw += string(DEF_asciiZeroUnderOne + high4bit)
		if low4bit > 0 {
			raw += string(DEF_asciiZeroUnderOne + low4bit)
		}
	}
	return raw
}

func decodeHeader(lenRaw int, header byte) (lenDecimal int) {
	lenStandard := header & DEF_headerBitMaskStandardLen
	lenDecimal = lenRaw - int(lenStandard) + DEF_digitDecimalMax - 1
	return lenDecimal
}
