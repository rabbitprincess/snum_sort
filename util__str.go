package snum

//-------------------------------------------------------------------------------------------//
// T_Snum

func (t *Snum) Cmp__str(_sn string) (n_cmp int) {
	pt_snum := &Snum{}
	pt_snum.Init(0, 0)
	pt_snum.Set__str(_sn)

	return t.Cmp(pt_snum)
}

//-------------------------------------------------------------------------------------------//
// global

func Cmp__str(_sn_a, _sn_b string) (n_cmp int, err error) {
	pt_a := &Snum{}
	pt_a.Init(0, 0)
	err = pt_a.Set__str(_sn_a)
	if err != nil {
		return 0, err
	}

	n_cmp = pt_a.Cmp__str(_sn_b)
	return n_cmp, nil
}

func Abs__str(_sn string) (sn string, err error) {
	pt := &Snum{}
	pt.Init(0, 0)
	err = pt.Set__str(_sn)
	if err != nil {
		return "", err
	}

	pt.Abs()
	sn = pt.String()
	return sn, nil
}

func Neg__str(_sn string) (sn string, err error) {
	pt := &Snum{}
	pt.Init(0, 0)
	err = pt.Set__str(_sn)
	if err != nil {
		return "", err
	}

	pt.Neg()
	sn = pt.String()
	return sn, nil
}

func Round__str(_sn string, _n_step_size int) (sn string, err error) {
	pt := &Snum{}
	pt.Init(0, 0)
	err = pt.Set__str(_sn)
	if err != nil {
		return "", err
	}

	pt.Round(_n_step_size)
	sn = pt.String()
	return sn, nil
}

func Round_down__str(_sn string, _n_step_size int) (sn string, err error) {
	pt := &Snum{}
	pt.Init(0, 0)
	err = pt.Set__str(_sn)
	if err != nil {
		return "", err
	}
	pt.Round_down(_n_step_size)
	sn = pt.String()
	return sn, nil
}

func Round_up__str(_sn string, _n_step_size int) (sn string, err error) {
	pt := &Snum{}
	pt.Init(0, 0)
	err = pt.Set__str(_sn)
	if err != nil {
		return "", err
	}
	pt.Round_up(_n_step_size)
	sn = pt.String()
	return sn, nil
}

func Pow(_sn string, _n8_num int64) (sn string, err error) {
	pt := &Snum{}
	pt.Init(0, 0)
	err = pt.Set__str(_sn)
	if err != nil {
		return "", err
	}
	pt.Pow(_n8_num)

	return pt.String(), nil
}

func Pow10(_n8_num int64) (sn string, err error) {
	return Pow("10", _n8_num)
}
