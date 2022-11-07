package snum

import (
	"fmt"
	"math"
	"math/big"

	"github.com/ericlagergren/decimal"
)

//----------------------------------------------------------------------------------------//
// uint64

func (t *Snum) GetUint64() (u8 uint64, err error) {
	if t.decimal == nil {
		return 0, nil
	}
	err = t.decimal.Context.Err()
	if err != nil {
		return 0, err
	}

	u8, valid := t.decimal.Uint64()
	if valid != true {
		return 0, fmt.Errorf("Failed to convert to uint64 | %s", t.String())
	}
	return u8, nil
}

func (t *Snum) SetUint64(u8 uint64) (err error) {
	if t.decimal == nil {
		t.Init()
	}
	t.decimal.SetUint64(u8)
	return nil
}

//----------------------------------------------------------------------------------------//
// str

func (t *Snum) GetStr() (sn string, err error) {
	if t.decimal == nil {
		return "", nil
	}
	err = t.decimal.Context.Err()
	if err != nil {
		return "", err
	}

	sn = fmt.Sprintf("%f", t.decimal.Reduce())
	return sn, nil
}

func (t *Snum) SetStr(sn string) (err error) {
	if t.decimal == nil {
		t.Init()
	}
	_, valid := t.decimal.SetString(sn)
	if valid != true {
		return fmt.Errorf("invalid number | %s", sn)
	}
	return nil
}

//----------------------------------------------------------------------------------------//
// raw

func (t *Snum) GetRaw() (bigNum *big.Int, lenDecimal int, isMinus bool) {
	pu8, pbig := decimal.Raw(t.decimal.Reduce())
	if *pu8 < math.MaxUint64 { // under maxUint64
		bigNum = big.NewInt(0).SetUint64(*pu8)
	} else { // over maxUint64
		bigNum = pbig
	}

	// - 처리
	if t.decimal.Sign() < 0 {
		isMinus = true
	}

	lenDecimal = t.decimal.Scale()
	return bigNum, lenDecimal, isMinus
}

func (t *Snum) SetRaw(big *big.Int, lenDecimal int, isMinus bool) {
	if t.decimal == nil {
		t.Init() // 임시
	}

	t.decimal.SetBigMantScale(big, lenDecimal)

	// - 처리
	if isMinus == true {
		t.decimal.Neg(t.decimal)
	}
}

//------------------------------------------------------------------------------------------//
// util

func (t *Snum) SetZero() {
	if t.decimal == nil {
		t.Init()
	}
	t.decimal.SetUint64(0)
}

func (t *Snum) TrimDigit(lenInteger, lenDecimal int) error {
	lenIntegerNow := t.decimal.Precision() - t.decimal.Scale()
	lenDecimalNow := t.decimal.Scale()

	var errDecimal, errInteger error

	// 후처리 - 소수
	if lenDecimalNow > lenDecimal {
		errDecimal = fmt.Errorf("Decimal limit exceeded | input : %d | limit : %d", lenDecimalNow, lenDecimal) // 에러 처리
		t.decimal.Quantize(lenDecimal)
	}

	// 후처리 - 정수
	if lenIntegerNow > lenInteger {
		errInteger = fmt.Errorf("Integer limit exceeded | input : %d | limit : %d", lenIntegerNow, lenInteger) // 에러 처리
		pt_snum := NewSnum(0)

		pt_snum.decimal.SetUint64(10)
		pt_snum.Pow(int64(lenInteger))
		lenDecimalBefore := t.decimal.Scale()
		t.decimal.Rem(t.decimal, pt_snum.decimal)
		t.decimal.SetScale(lenDecimalBefore)
	}

	// 에러일 경우 리턴
	if errDecimal != nil {
		return errDecimal
	}
	if errInteger != nil {
		return errInteger
	}
	return nil
}
