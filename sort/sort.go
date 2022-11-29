package sort

import (
	"errors"
	"math/big"

	"github.com/gokch/snum_sort/snum"
)

func NewSnumSort[T snum.SnumConst](num T) *SnumSort {
	return &SnumSort{
		Snum: *snum.NewSnum(num),
	}
}

type SnumSort struct {
	snum.Snum
}

func (t *SnumSort) Encode() (enc []byte, err error) {
	bigRaw, lenDecimal, isMinus := t.Snum.GetRaw()

	raw := bigRaw.String()
	if raw == "0" { // 0 일 경우 전처리
		lenDecimal = DEF_digitDecimalMax
	}

	header := encodeHeader(len(raw), lenDecimal)
	body := encodeBody(raw)
	enc = encodeMinus(isMinus, append([]byte{header}, body...))
	return enc, nil
}

func (t *SnumSort) Decode(enc []byte) (err error) {
	if len(enc) < DEF_lenTotalMin {
		return errors.New("too short")
	}
	isMinus, enc := decodeMinus(enc)
	raw := decodeBody(isMinus, enc[1:])
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
	// 소수점 시작 위치 추출 ( 자릿수는 1번이 0번 idx 지만 길이는 1의 자리가 1 이기 때문에 1 감소 필요 )
	posStartDot := lenRaw - lenDecimal + DEF_digitDecimalMax - 1
	header = DEF_headerBitMaskSign | byte(posStartDot)
	return header
}

func encodeBody(raw string) (body []byte) {
	body = make([]byte, 0, len(raw)) // len / 2 + (len%2!=0)?1:0
	numOri := []byte(raw)
	for i := 0; i < len(raw); i++ {
		b4 := numOri[i] - byte('0')
		if i%2 == 0 {
			body = append(body, b4<<4)
		} else {
			body[i/2] += b4
		}
	}
	return body
}

func encodeMinus(isMinus bool, enc []byte) []byte {
	if isMinus == true {
		for i := 0; i < len(enc); i++ {
			enc[i] = ^enc[i]
		}
		enc = append(enc, 0xFF) // append last 0xFF
	}
	return enc
}

func decodeMinus(enc []byte) (isMinus bool, dec []byte) {
	if enc[0]&DEF_headerBitMaskSign == 0 {
		enc = enc[:len(enc)-1] // separate last 0xFF
		for i := 0; i < len(enc); i++ {
			enc[i] = ^enc[i]
		}
		return true, enc
	}
	return false, enc
}

func decodeBody(isMinus bool, body []byte) (raw string) {
	for i := 0; i < len(body); i++ {
		high4bit := body[i] >> 4
		low4bit := body[i] - (high4bit << 4)
		raw += string('0' + high4bit)
		raw += string('0' + low4bit)
	}
	return raw
}

func decodeHeader(lenRaw int, header byte) (lenDecimal int) {
	lenStandard := header & DEF_headerBitMaskStandardLen
	lenDecimal = lenRaw - int(lenStandard) + DEF_digitDecimalMax - 1
	return lenDecimal
}
