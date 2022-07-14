package snum

import (
	"github.com/ericlagergren/decimal"
)

type Snum struct {
	decimal      *decimal.Big
	len_standard int
	len_decimal  int
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
			Precision:     DEF_b1_header__max_len__standard + DEF_b1_header__max_len__decimal,
			RoundingMode:  decimal.ToZero,
		},
	}
	t.len_standard = _len_standard
	t.len_decimal = _len_decimal
}

func (t *Snum) Copy() *Snum {
	pt_ret := &Snum{}
	pt_ret.Init(t.len_standard, t.len_decimal)
	pt_ret.decimal.Copy(t.decimal)
	return pt_ret
}

// for interface - package json
func (t *Snum) UnmarshalJSON(_bt []byte) error {
	err := t.Set__str(string(_bt))
	return err
}

// for interface - package json
func (t *Snum) MarshalJSON() ([]byte, error) {
	s_num, err := t.Get__str()
	return []byte(s_num), err
}

func (t *Snum) String() string {
	s_num, _ := t.Get__str()
	return s_num
}
