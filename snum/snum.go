package snum

import (
	"github.com/ericlagergren/decimal"
)

type SnumConst interface {
	*Snum | int | int32 | int64 | uint | uint32 | uint64 | string
}

func New[T SnumConst](num T) *Snum {
	snum := &Snum{}
	snum.Init()

	switch data := any(&num).(type) {
	case **Snum:
		snum.decimal.Copy((*data).decimal)
	case *int:
		if *data < 0 {
			snum.decimal.SetUint64(uint64(-*data))
			snum.decimal.Neg(snum.decimal)
		} else {
			snum.decimal.SetUint64(uint64(*data))
		}
	case *int32:
		if *data < 0 {
			snum.decimal.SetUint64(uint64(-*data))
			snum.decimal.Neg(snum.decimal)
		} else {
			snum.decimal.SetUint64(uint64(*data))
		}
	case *int64:
		if *data < 0 {
			snum.decimal.SetUint64(uint64(-*data))
			snum.decimal.Neg(snum.decimal)
		} else {
			snum.decimal.SetUint64(uint64(*data))
		}
	case *uint:
		snum.decimal.SetUint64(uint64(*data))
	case *uint32:
		snum.decimal.SetUint64(uint64(*data))
	case *uint64:
		snum.decimal.SetUint64(*data)
	case *string:
		err := snum.SetStr(*data)
		if err != nil {
			return nil
		}
	}
	return snum
}

type Snum struct {
	decimal *decimal.Big
}

func (t *Snum) Init() {
	t.decimal = &decimal.Big{
		Context: decimal.Context{
			OperatingMode: decimal.GDA,
			Precision:     128,
			RoundingMode:  decimal.ToZero,
		},
	}
}

func (t *Snum) Copy() *Snum {
	ret := New(0)
	ret.decimal.Copy(t.decimal)
	return ret
}

func (t *Snum) UnmarshalJSON(bt []byte) error {
	return t.SetStr(string(bt))
}

func (t *Snum) MarshalJSON() ([]byte, error) {
	return []byte(t.String()), nil
}

func (t *Snum) String() string {
	return t.GetStr()
}
