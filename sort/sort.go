package sort

import (
	"bytes"
	"encoding/hex"
	"errors"

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

func (t *SnumSort) UnmarshalJSON(bt []byte) error {
	if len(bt) < DEF_lenTotalMin+2 {
		return errors.New("too short")
	}
	enc, err := hex.DecodeString(string(bt[1 : len(bt)-1]))
	if err != nil {
		return err
	}
	t.Decode(enc)
	return nil
}

func (t *SnumSort) MarshalJSON() ([]byte, error) {
	enc := t.Encode()

	buf := bytes.NewBuffer(nil)
	buf.WriteByte('"')
	buf.WriteString(hex.EncodeToString(enc))
	buf.WriteByte('"')
	return buf.Bytes(), nil
}
