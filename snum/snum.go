package snum

import (
	"github.com/ericlagergren/decimal"
)

func NewSnum[T *Snum | int | int32 | int64 | uint | uint32 | uint64 | string | float32 | float64](num T) *Snum {
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
		snum.SetStr(*data)
	case *float32:
		if *data < 0 {
			snum.decimal.SetFloat64(-float64(*data))
			snum.decimal.Neg(snum.decimal)
		} else {
			snum.decimal.SetFloat64(float64(*data))
		}
	case *float64:
		if *data < 0 {
			snum.decimal.SetFloat64(-*data)
			snum.decimal.Neg(snum.decimal)
		} else {
			snum.decimal.SetFloat64(*data)
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
	ret := NewSnum(0)
	ret.decimal.Copy(t.decimal)
	return ret
}

func (t *Snum) UnmarshalJSON(bt []byte) error {
	err := t.SetStr(string(bt))
	return err
}

func (t *Snum) MarshalJSON() ([]byte, error) {
	num, err := t.GetStr()
	return []byte(num), err
}

func (t *Snum) String() string {
	num, _ := t.GetStr()
	return num
}