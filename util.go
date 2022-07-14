package snum

import (
	"github.com/ericlagergren/decimal"
)

func (t *Snum) Is_zero() bool {
	pt_zero := &Snum{}
	pt_zero.Init(0, 0)
	pt_zero.Set__zero()

	n_cmp := t.Cmp(pt_zero)
	if n_cmp != 0 {
		return false
	}
	return true
}

func (t *Snum) Is_zero__under() bool {
	pt_zero := &Snum{}
	pt_zero.Init(0, 0)
	pt_zero.Set__zero()

	n_cmp := t.Cmp(pt_zero)
	if n_cmp != -1 {
		return false
	}
	return true
}

func (t *Snum) Is_zero__over() bool {
	pt_zero := &Snum{}
	pt_zero.Init(0, 0)
	pt_zero.Set__zero()

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

func (t *Snum) Round_down(_step_size int) {
	t.decimal.Context.RoundingMode = decimal.ToZero
	t.decimal.Round(_step_size)
}

func (t *Snum) Round_up(_step_size int) {
	t.decimal.Context.RoundingMode = decimal.AwayFromZero
	t.decimal.Round(_step_size)
	t.decimal.Context.RoundingMode = decimal.ToZero // round_down 기준으로 복구
}

func (t *Snum) Pow(_num int64) {
	t.decimal.Context.Pow(t.decimal, t.decimal, decimal.New(_num, 0))
}

// Output:
//  x         step_size      Group_down      Group_up
//  123.321          -4         123.321       123.321
//  123.321          -3         123.321       123.321
//  123.321          -2         123.32        123.33
//  123.321          -1         123.3         123.4
//  123.321           0         123           124
//  123.321           1         120           130
//  123.321           2         100           200
//  123.321           3         0             1000
//  123.321           4         0             10000
func (t *Snum) Group_down(_n_step_size int) {
	n_len_decimal := t.decimal.Scale()
	n_len_integer := t.decimal.Precision() - n_len_decimal

	if n_len_integer <= _n_step_size {
		// step_size 가 snum 자릿수를 초과할 경우 0 반환
		t.decimal = decimal.New(0, 0)
	} else {
		t.decimal.Context.RoundingMode = decimal.ToZero
		t.decimal.Quantize(-_n_step_size)
	}
}

func (t *Snum) Group_up(_n_step_size int) {
	n_len_decimal := t.decimal.Scale()
	n_len_integer := t.decimal.Precision() - n_len_decimal

	if n_len_integer <= _n_step_size {
		// step_size 가 snum 자릿수를 초과할 경우 10^step_size 반환
		t.decimal = decimal.New(10, 0)
		t.Pow(int64(_n_step_size))
	} else {
		t.decimal.Context.RoundingMode = decimal.AwayFromZero
		t.decimal.Quantize(-_n_step_size)
		t.decimal.Context.RoundingMode = decimal.ToZero
	}
}
