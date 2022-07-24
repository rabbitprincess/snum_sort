package snum

//-------------------------------------------------------------------------------------------//
// T_Snum

func (t *Snum) CmpStr(_sn string) (n_cmp int) {
	pt_snum := &Snum{}
	pt_snum.Init(0, 0)
	pt_snum.SetStr(_sn)

	return t.Cmp(pt_snum)
}

//-------------------------------------------------------------------------------------------//
// global

func CmpStr(_sn_a, _sn_b string) (n_cmp int, err error) {
	pt_a := &Snum{}
	pt_a.Init(0, 0)
	err = pt_a.SetStr(_sn_a)
	if err != nil {
		return 0, err
	}

	n_cmp = pt_a.CmpStr(_sn_b)
	return n_cmp, nil
}

func AbsStr(_sn string) (sn string, err error) {
	pt := &Snum{}
	pt.Init(0, 0)
	err = pt.SetStr(_sn)
	if err != nil {
		return "", err
	}

	pt.Abs()
	sn = pt.String()
	return sn, nil
}

func NegStr(_sn string) (sn string, err error) {
	pt := &Snum{}
	pt.Init(0, 0)
	err = pt.SetStr(_sn)
	if err != nil {
		return "", err
	}

	pt.Neg()
	sn = pt.String()
	return sn, nil
}

func RoundStr(_sn string, _stepSize int) (sn string, err error) {
	pt := &Snum{}
	pt.Init(0, 0)
	err = pt.SetStr(_sn)
	if err != nil {
		return "", err
	}

	pt.Round(_stepSize)
	sn = pt.String()
	return sn, nil
}

func RoundDownStr(_sn string, _stepSize int) (sn string, err error) {
	pt := &Snum{}
	pt.Init(0, 0)
	err = pt.SetStr(_sn)
	if err != nil {
		return "", err
	}
	pt.RoundDown(_stepSize)
	sn = pt.String()
	return sn, nil
}

func RoundUpStr(_sn string, _stepSize int) (sn string, err error) {
	pt := &Snum{}
	pt.Init(0, 0)
	err = pt.SetStr(_sn)
	if err != nil {
		return "", err
	}
	pt.RoundUp(_stepSize)
	sn = pt.String()
	return sn, nil
}

func Pow(_sn string, _n8_num int64) (sn string, err error) {
	pt := &Snum{}
	pt.Init(0, 0)
	err = pt.SetStr(_sn)
	if err != nil {
		return "", err
	}
	pt.Pow(_n8_num)

	return pt.String(), nil
}

func Pow10(_n8_num int64) (sn string, err error) {
	return Pow("10", _n8_num)
}
