package snum

import (
	"github.com/ericlagergren/decimal"
)

func (t *Snum) IsZero() bool {
	pt_zero := &Snum{}
	pt_zero.Init(0, 0)
	pt_zero.SetZero()

	n_cmp := t.Cmp(pt_zero)
	if n_cmp != 0 {
		return false
	}
	return true
}

func (t *Snum) IsZeroUnder() bool {
	pt_zero := &Snum{}
	pt_zero.Init(0, 0)
	pt_zero.SetZero()

	n_cmp := t.Cmp(pt_zero)
	if n_cmp != -1 {
		return false
	}
	return true
}

func (t *Snum) IsZeroOver() bool {
	pt_zero := &Snum{}
	pt_zero.Init(0, 0)
	pt_zero.SetZero()

	n_cmp := t.Cmp(pt_zero)
	if n_cmp != 1 {
		return false
	}
	return true
}

// if( t  <  _pt ) -> -1
// if( t  == _pt ) -> 0
// if( t  >  _pt ) -> +1
func (t *Snum) Cmp(_pt *Snum) int {
	return t.decimal.Cmp(_pt.decimal)
}

func (t *Snum) Abs() {
	t.decimal.Abs(t.decimal)
}

func (t *Snum) Neg() {
	t.decimal.Neg(t.decimal)
}

// Output:
//    x      Round      Round_down      Round_up
//  2.6          3               2             3
//  2.5          3               2             3
//  2.1          2               2             3
//  2            2               2			   2
// -2.1         -2              -2            -3
// -2.5         -3              -2            -3
// -2.6         -3              -2            -3
func (t *Snum) Round(_step_size int) {
	t.decimal.Context.RoundingMode = decimal.ToNearestAway
	t.decimal.Round(_step_size)
	t.decimal.Context.RoundingMode = decimal.ToZero // round_down 기준으로 복구
}

func (t *Snum) RoundDown(_stepSize int) {
	t.decimal.Context.RoundingMode = decimal.ToZero
	t.decimal.Round(_stepSize)
}

func (t *Snum) RoundUp(_stepSize int) {
	t.decimal.Context.RoundingMode = decimal.AwayFromZero
	t.decimal.Round(_stepSize)
	t.decimal.Context.RoundingMode = decimal.ToZero // round_down 기준으로 복구
}

func (t *Snum) Pow(_num int64) {
	t.decimal.Context.Pow(t.decimal, t.decimal, decimal.New(_num, 0))
}

// Output:
//  x         step_size      GroupDown      Group_up
//  123.321          -4         123.321       123.321
//  123.321          -3         123.321       123.321
//  123.321          -2         123.32        123.33
//  123.321          -1         123.3         123.4
//  123.321           0         123           124
//  123.321           1         120           130
//  123.321           2         100           200
//  123.321           3         0             1000
//  123.321           4         0             10000
func (t *Snum) GroupDown(_stepSize int) {
	lenDecimal := t.decimal.Scale()
	lenInteger := t.decimal.Precision() - lenDecimal

	if lenInteger <= _stepSize {
		// step_size 가 snum 자릿수를 초과할 경우 0 반환
		t.decimal = decimal.New(0, 0)
	} else {
		t.decimal.Context.RoundingMode = decimal.ToZero
		t.decimal.Quantize(-_stepSize)
	}
}

func (t *Snum) GroupUp(_stepSize int) {
	lenDecimal := t.decimal.Scale()
	lenInteger := t.decimal.Precision() - lenDecimal

	if lenInteger <= _stepSize {
		// step_size 가 snum 자릿수를 초과할 경우 10^step_size 반환
		t.decimal = decimal.New(10, 0)
		t.Pow(int64(_stepSize))
	} else {
		t.decimal.Context.RoundingMode = decimal.AwayFromZero
		t.decimal.Quantize(-_stepSize)
		t.decimal.Context.RoundingMode = decimal.ToZero
	}
}
