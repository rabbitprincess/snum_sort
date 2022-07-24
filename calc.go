package snum

//------------------------------------------------------------------------//
// global

func Add(_a *Snum, _b *Snum) (ret *Snum) {
	ret = _a.Copy()
	ret.Add(_b)
	return ret
}

func AddStr(_a, _b string) (sn string, err error) {
	pt_a := &Snum{}
	pt_a.Init(0, 0)
	err = pt_a.SetStr(_a)
	if err != nil {
		return "", err
	}

	pt_b := &Snum{}
	pt_b.Init(0, 0)
	err = pt_b.SetStr(_b)
	if err != nil {
		return "", err
	}

	pt_add := Add(pt_a, pt_b)

	sn = pt_add.String()
	return sn, nil
}

func Sub(_a *Snum, _b *Snum) (pt_ret *Snum) {
	pt_ret = _a.Copy()
	pt_ret.Sub(_b)
	return pt_ret
}

func SubStr(_a, _b string) (sn string, err error) {
	pt_a := &Snum{}
	pt_a.Init(0, 0)
	err = pt_a.SetStr(_a)
	if err != nil {
		return "", err
	}

	pt_b := &Snum{}
	pt_b.Init(0, 0)
	err = pt_b.SetStr(_b)
	if err != nil {
		return "", err
	}

	pt_sub := Sub(pt_a, pt_b)
	sn = pt_sub.String()
	return sn, nil
}

func Mul(_a *Snum, _b *Snum) (pt_ret *Snum) {
	pt_ret = _a.Copy()
	pt_ret.Mul(_b)
	return pt_ret
}

func MulStr(_a, _b string) (sn string, err error) {
	pt_a := &Snum{}
	pt_a.Init(0, 0)
	err = pt_a.SetStr(_a)
	if err != nil {
		return "", err
	}

	pt_b := &Snum{}
	pt_b.Init(0, 0)
	err = pt_b.SetStr(_b)
	if err != nil {
		return "", err
	}

	pt_mul := Mul(pt_a, pt_b)
	sn = pt_mul.String()
	return sn, nil
}

func Div(_a *Snum, _b *Snum) (pt_ret *Snum) {
	pt_ret = _a.Copy()
	pt_ret.Div(_b)
	return pt_ret
}

func DivStr(_a, _b string) (sn string, err error) {
	pt_a := &Snum{}
	pt_a.Init(0, 0)
	err = pt_a.SetStr(_a)
	if err != nil {
		return "", err
	}

	pt_b := &Snum{}
	pt_b.Init(0, 0)
	err = pt_b.SetStr(_b)
	if err != nil {
		return "", err
	}

	pt_div := Div(pt_a, pt_b)
	sn = pt_div.String()
	return sn, nil
}

//-------------------------------------------------------------------------------//
// Tum

func (t *Snum) Add(_pt *Snum) {
	t.decimal.Add(t.decimal, _pt.decimal)
}

func (t *Snum) AddStr(_sn string) (err error) {
	ptum := &Snum{}
	ptum.Init(0, 0)
	err = ptum.SetStr(_sn)
	if err != nil {
		return err
	}

	t.Add(ptum)
	return nil
}

func (t *Snum) Sub(_pt *Snum) {
	t.decimal.Sub(t.decimal, _pt.decimal)
}

func (t *Snum) SubStr(_sn string) (err error) {
	ptum := &Snum{}
	ptum.Init(0, 0)
	err = ptum.SetStr(_sn)
	if err != nil {
		return err
	}

	t.Sub(ptum)
	return nil
}

func (t *Snum) Mul(_pt *Snum) {
	t.decimal.Mul(t.decimal, _pt.decimal)
}

func (t *Snum) MulStr(_sn string) (err error) {
	ptum := &Snum{}
	ptum.Init(0, 0)
	err = ptum.SetStr(_sn)
	if err != nil {
		return err
	}

	t.Mul(ptum)
	return nil
}

func (t *Snum) Div(_pt *Snum) {
	t.decimal.Quo(t.decimal, _pt.decimal)

	if t.decimal.IsNormal() == true {
		t.decimal.Quantize(DEF_lenExtendDecimalForCalc)
	}
}

func (t *Snum) DivStr(_sn string) (err error) {
	ptum := &Snum{}
	ptum.Init(0, 0)
	err = ptum.SetStr(_sn)

	t.Div(ptum)
	return nil
}
