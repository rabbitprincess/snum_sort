package snum

//-------------------------------------------------------------------------------------------//
// T_Snum

func (t *Snum) CmpStr(num string) (cmp int) {
	sn := &Snum{}
	sn.Init(0, 0)
	sn.SetStr(num)

	return t.Cmp(sn)
}

//-------------------------------------------------------------------------------------------//
// global

func CmpStr(a, b string) (cmp int, err error) {
	snA := &Snum{}
	snA.Init(0, 0)
	err = snA.SetStr(a)
	if err != nil {
		return 0, err
	}
	cmp = snA.CmpStr(b)
	return cmp, nil
}

func AbsStr(num string) (ret string, err error) {
	sn := &Snum{}
	sn.Init(0, 0)
	err = sn.SetStr(num)
	if err != nil {
		return "", err
	}

	sn.Abs()
	ret = sn.String()
	return ret, nil
}

func NegStr(num string) (ret string, err error) {
	pt := &Snum{}
	pt.Init(0, 0)
	err = pt.SetStr(num)
	if err != nil {
		return "", err
	}

	pt.Neg()
	ret = pt.String()
	return ret, nil
}

func RoundStr(num string, stepSize int) (ret string, err error) {
	sn := &Snum{}
	sn.Init(0, 0)
	err = sn.SetStr(num)
	if err != nil {
		return "", err
	}

	sn.Round(stepSize)
	ret = sn.String()
	return ret, nil
}

func RoundDownStr(num string, stepSize int) (ret string, err error) {
	sn := &Snum{}
	sn.Init(0, 0)
	err = sn.SetStr(num)
	if err != nil {
		return "", err
	}
	sn.RoundDown(stepSize)
	ret = sn.String()
	return ret, nil
}

func RoundUpStr(num string, stepSize int) (ret string, err error) {
	sn := &Snum{}
	sn.Init(0, 0)
	err = sn.SetStr(num)
	if err != nil {
		return "", err
	}
	sn.RoundUp(stepSize)
	ret = sn.String()
	return ret, nil
}

func Pow(snum string, nnum int64) (ret string, err error) {
	sn := &Snum{}
	sn.Init(0, 0)
	err = sn.SetStr(snum)
	if err != nil {
		return "", err
	}
	sn.Pow(nnum)

	return sn.String(), nil
}

func Pow10(num int64) (ret string, err error) {
	return Pow("10", num)
}
