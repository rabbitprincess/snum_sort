package snum

import (
	"github.com/ericlagergren/decimal"
)

type Snum struct {
	decimal     *decimal.Big
	lenStandard int
	lenDecimal  int
}

func (t *Snum) Init(lenStandard, lenDecimal int) {
	if lenStandard == 0 {
		lenStandard = 96
	}
	if lenDecimal == 0 {
		lenDecimal = 32
	}

	t.decimal = &decimal.Big{
		Context: decimal.Context{
			OperatingMode: decimal.GDA,
			Precision:     DEF_headerLenInteger + DEF_headerLenDecimal,
			RoundingMode:  decimal.ToZero,
		},
	}
	t.lenStandard = lenStandard
	t.lenDecimal = lenDecimal
}

func (t *Snum) Copy() *Snum {
	ret := &Snum{}
	ret.Init(t.lenStandard, t.lenDecimal)
	ret.decimal.Copy(t.decimal)
	return ret
}

// for interface - package json
func (t *Snum) UnmarshalJSON(bt []byte) error {
	err := t.SetStr(string(bt))
	return err
}

// for interface - package json
func (t *Snum) MarshalJSON() ([]byte, error) {
	num, err := t.GetStr()
	return []byte(num), err
}

func (t *Snum) String() string {
	num, _ := t.GetStr()
	return num
}
