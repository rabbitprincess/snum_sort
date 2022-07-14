package snum

import (
	"fmt"
	"math"
	"math/big"
	"strconv"

	"github.com/ericlagergren/decimal"
)

//----------------------------------------------------------------------------------------//
// uint64

func (t *Snum) Get__uint64() (u8_num uint64, err error) {
	if t.decimal == nil {
		return 0, nil
	}
	err = t.decimal.Context.Err()
	if err != nil {
		return 0, err
	}

	u8_num, is_valid := t.decimal.Uint64()
	if is_valid != true {
		return 0, fmt.Errorf("Failed to convert to uint64 | %s", t.String())
	}
	return u8_num, nil
}

func (t *Snum) Set__uint64(_u8_num uint64) (err error) {
	if t.decimal == nil {
		t.Init(0, 0)
	}
	t.decimal.SetUint64(_u8_num)

	return nil
}

//----------------------------------------------------------------------------------------//
// str

func (t *Snum) Get__str() (s_num string, err error) {
	if t.decimal == nil {
		return "", nil
	}
	err = t.decimal.Context.Err()
	if err != nil {
		return "", err
	}

	s_num = fmt.Sprintf("%f", t.decimal.Reduce())
	return s_num, nil
}

func (t *Snum) Set__str(_s_num string) (err error) {
	if t.decimal == nil {
		t.Init(0, 0)
	}
	_, is_valid := t.decimal.SetString(_s_num)
	if is_valid != true {
		return fmt.Errorf("invalid number | %s", _s_num)
	}

	return nil
}

//----------------------------------------------------------------------------------------//
// raw

func (t *Snum) Get__raw() (s_raw string, n_len_decimal int, is_minus bool) {
	pu8_int, pt_big := decimal.Raw(t.decimal.Reduce())
	if *pu8_int < math.MaxUint64 { // under maxUint64
		s_raw = strconv.FormatUint(*pu8_int, 10)
	} else { // over maxUint64
		s_raw = pt_big.String()
	}

	// - 처리
	if t.decimal.Sign() < 0 {
		is_minus = true
	}

	n_len_decimal = t.decimal.Scale()
	return s_raw, n_len_decimal, is_minus
}

func (t *Snum) Set__raw(_s_raw string, _n_len_decimal int, _is_minus bool) {
	if t.decimal == nil {
		t.Init(0, 0) // 임시
	}
	pt_big := big.NewInt(0)
	pt_big.SetString(_s_raw, 10)
	t.decimal.SetBigMantScale(pt_big, _n_len_decimal)

	// - 처리
	if _is_minus == true {
		t.decimal.Neg(t.decimal)
	}
}

//------------------------------------------------------------------------------------------//
// util

func (t *Snum) Set__zero() {
	if t.decimal == nil {
		t.Init(0, 0)
	}
	t.decimal.SetUint64(0)
}

func (t *Snum) Trim_digit() error {
	n_len_decimal := t.decimal.Scale()
	n_len_integer := t.decimal.Precision() - t.decimal.Scale()

	var err_decimal, err_integer error

	// 후처리 - 소수
	if n_len_decimal > t.len_decimal {
		err_decimal = fmt.Errorf("Decimal limit exceeded | input : %d | limit : %d", n_len_decimal, t.len_decimal) // 에러 처리

		t.decimal.Quantize(t.len_decimal)
	}

	// 후처리 - 정수
	if n_len_integer > t.len_standard {
		err_integer = fmt.Errorf("Integer limit exceeded | input : %d | limit : %d", n_len_integer, t.len_standard) // 에러 처리

		pt_snum := &Snum{}
		pt_snum.Init(0, 0)
		pt_snum.decimal.SetUint64(10)
		pt_snum.Pow(int64(t.len_standard))
		n_len_decimal__before := t.decimal.Scale()

		t.decimal.Rem(t.decimal, pt_snum.decimal)
		t.decimal.SetScale(n_len_decimal__before)
	}

	// 에러일 경우 리턴
	if err_decimal != nil {
		return err_decimal
	}
	if err_integer != nil {
		return err_integer
	}
	return nil
}
