package snum

import (
	"github.com/ericlagergren/decimal"
)

type Snum struct {
	decimal     *decimal.Big
	lenStandard int
	lenDecimal  int
}

func (t *Snum) Init(_len_standard, _len_decimal int) {
	if _len_standard == 0 {
		_len_standard = 100
	}
	if _len_decimal == 0 {
		_len_decimal = 20
	}

	t.decimal = &decimal.Big{
		Context: decimal.Context{
			OperatingMode: decimal.GDA,
			Precision:     DEF_headerLenInteger + DEF_headerLenDecimal,
			RoundingMode:  decimal.ToZero,
		},
	}
	t.lenStandard = _len_standard
	t.lenDecimal = _len_decimal
}

func (t *Snum) Copy() *Snum {
	pt_ret := &Snum{}
	pt_ret.Init(t.lenStandard, t.lenDecimal)
	pt_ret.decimal.Copy(t.decimal)
	return pt_ret
}

// for interface - package json
func (t *Snum) UnmarshalJSON(_bt []byte) error {
	err := t.SetStr(string(_bt))
	return err
}

// for interface - package json
func (t *Snum) MarshalJSON() ([]byte, error) {
	s_num, err := t.GetStr()
	return []byte(s_num), err
}

func (t *Snum) String() string {
	s_num, _ := t.GetStr()
	return s_num
}
