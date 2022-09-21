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

	u8, is_valid := t.decimal.Uint64()
	if is_valid != true {
		return 0, fmt.Errorf("Failed to convert to uint64 | %s", t.String())
	}
	return u8, nil
}

func (t *Snum) SetUint64(u8 uint64) (err error) {
	if t.decimal == nil {
		t.Init(0, 0)
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

func (t *Snum) SetStr(_sn string) (err error) {
	if t.decimal == nil {
		t.Init(0, 0)
	}
	_, is_valid := t.decimal.SetString(_sn)
	if is_valid != true {
		return fmt.Errorf("invalid number | %s", _sn)
	}

	return nil
}

//----------------------------------------------------------------------------------------//
// raw

func (t *Snum) GetRaw() (big *big.Int, lenDecimal int, isMinus bool) {
	pu8, pbig := decimal.Raw(t.decimal.Reduce())
	if *pu8 < math.MaxUint64 { // under maxUint64
		big = big.SetUint64(*pu8)
	} else { // over maxUint64
		big = pbig
	}

	// - 처리
	if t.decimal.Sign() < 0 {
		isMinus = true
	}

	lenDecimal = t.decimal.Scale()
	return big, lenDecimal, isMinus
}

func (t *Snum) SetRaw(big *big.Int, lenDecimal int, isMinus bool) {
	if t.decimal == nil {
		t.Init(0, 0) // 임시
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
		t.Init(0, 0)
	}
	t.decimal.SetUint64(0)
}

func (t *Snum) TrimDigit() error {
	lenDecimal := t.decimal.Scale()
	lenInteger := t.decimal.Precision() - t.decimal.Scale()

	var errDecimal, errInteger error

	// 후처리 - 소수
	if lenDecimal > t.lenDecimal {
		errDecimal = fmt.Errorf("Decimal limit exceeded | input : %d | limit : %d", lenDecimal, t.lenDecimal) // 에러 처리

		t.decimal.Quantize(t.lenDecimal)
	}

	// 후처리 - 정수
	if lenInteger > t.lenStandard {
		errInteger = fmt.Errorf("Integer limit exceeded | input : %d | limit : %d", lenInteger, t.lenStandard) // 에러 처리

		pt_snum := &Snum{}
		pt_snum.Init(0, 0)
		pt_snum.decimal.SetUint64(10)
		pt_snum.Pow(int64(t.lenStandard))
		lenDecimal_before := t.decimal.Scale()

		t.decimal.Rem(t.decimal, pt_snum.decimal)
		t.decimal.SetScale(lenDecimal_before)
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
