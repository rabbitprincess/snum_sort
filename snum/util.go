package snum

import (
	"github.com/ericlagergren/decimal"
)

func (t *Snum) IsZero() bool {
	if t.Cmp(NewSnum(0)) != 0 {
		return false
	}
	return true
}

func (t *Snum) IsZeroUnder() bool {
	if t.Cmp(NewSnum(0)) < 0 {
		return false
	}
	return true
}

func (t *Snum) IsZeroOver() bool {
	if t.Cmp(NewSnum(0)) > 0 {
		return false
	}
	return true
}

// if( t  <  _pt ) -> -1
// if( t  == _pt ) -> 0
// if( t  >  _pt ) -> +1
func (t *Snum) Cmp(sn *Snum) int {
	return t.decimal.Cmp(sn.decimal)
}

func (t *Snum) Abs() {
	t.decimal.Abs(t.decimal)
}

func (t *Snum) Neg() {
	t.decimal.Neg(t.decimal)
}

// Output:
//
//	  x      Round      Round_down      Round_up
//	2.6          3               2             3
//	2.5          3               2             3
//	2.1          2               2             3
//	2            2               2			   2
//
// -2.1         -2              -2            -3
// -2.5         -3              -2            -3
// -2.6         -3              -2            -3
func (t *Snum) Round(step int) {
	t.decimal.Context.RoundingMode = decimal.ToNearestAway
	t.decimal.Round(step)
	t.decimal.Context.RoundingMode = decimal.ToZero // round_down 기준으로 복구
}

func (t *Snum) RoundDown(step int) {
	t.decimal.Context.RoundingMode = decimal.ToZero
	t.decimal.Round(step)
}

func (t *Snum) RoundUp(step int) {
	t.decimal.Context.RoundingMode = decimal.AwayFromZero
	t.decimal.Round(step)
	t.decimal.Context.RoundingMode = decimal.ToZero // round_down 기준으로 복구
}

func (t *Snum) Pow(num int64) {
	t.decimal.Context.Pow(t.decimal, t.decimal, decimal.New(num, 0))
}

// Output:
//
//	x              step      GroupDown      Group_up
//	123.321          -4         123.321       123.321
//	123.321          -3         123.321       123.321
//	123.321          -2         123.32        123.33
//	123.321          -1         123.3         123.4
//	123.321           0         123           124
//	123.321           1         120           130
//	123.321           2         100           200
//	123.321           3         0             1000
//	123.321           4         0             10000
func (t *Snum) GroupDown(step int) {
	lenDecimal := t.decimal.Scale()
	lenInteger := t.decimal.Precision() - lenDecimal

	if lenInteger <= step {
		// step_size 가 snum 자릿수를 초과할 경우 0 반환
		t.decimal = decimal.New(0, 0)
	} else {
		t.decimal.Context.RoundingMode = decimal.ToZero
		t.decimal.Quantize(-step)
	}
}

func (t *Snum) GroupUp(step int) {
	lenDecimal := t.decimal.Scale()
	lenInteger := t.decimal.Precision() - lenDecimal

	if lenInteger <= step {
		// step_size 가 snum 자릿수를 초과할 경우 10^step_size 반환
		t.decimal = decimal.New(10, 0)
		t.Pow(int64(step))
	} else {
		t.decimal.Context.RoundingMode = decimal.AwayFromZero
		t.decimal.Quantize(-step)
		t.decimal.Context.RoundingMode = decimal.ToZero
	}
}
