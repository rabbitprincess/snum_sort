package snum

//------------------------------------------------------------------------//
// global

func Add(a, b *Snum) (ret *Snum) {
	ret = a.Copy()
	ret.Add(b)
	return ret
}

func AddStr(a, b string) (ret string, err error) {
	snA := &Snum{}
	snA.Init(0, 0)
	err = snA.SetStr(a)
	if err != nil {
		return "", err
	}

	snB := &Snum{}
	snB.Init(0, 0)
	err = snB.SetStr(b)
	if err != nil {
		return "", err
	}

	ret = Add(snA, snB).String()
	return ret, nil
}

func Sub(a, b *Snum) (ret *Snum) {
	ret = a.Copy()
	ret.Sub(b)
	return ret
}

func SubStr(a, b string) (ret string, err error) {
	snA := &Snum{}
	snA.Init(0, 0)
	err = snA.SetStr(a)
	if err != nil {
		return "", err
	}

	snB := &Snum{}
	snB.Init(0, 0)
	err = snB.SetStr(b)
	if err != nil {
		return "", err
	}

	pt_sub := Sub(snA, snB)
	ret = pt_sub.String()
	return ret, nil
}

func Mul(a *Snum, b *Snum) (ret *Snum) {
	ret = a.Copy()
	ret.Mul(b)
	return ret
}

func MulStr(a, b string) (ret string, err error) {
	snA := &Snum{}
	snA.Init(0, 0)
	err = snA.SetStr(a)
	if err != nil {
		return "", err
	}

	snB := &Snum{}
	snB.Init(0, 0)
	err = snB.SetStr(b)
	if err != nil {
		return "", err
	}

	pt_mul := Mul(snA, snB)
	ret = pt_mul.String()
	return ret, nil
}

func Div(a *Snum, b *Snum) (ret *Snum) {
	ret = a.Copy()
	ret.Div(b)
	return ret
}

func DivStr(a, b string) (ret string, err error) {
	ptA := &Snum{}
	ptA.Init(0, 0)
	err = ptA.SetStr(a)
	if err != nil {
		return "", err
	}

	snB := &Snum{}
	snB.Init(0, 0)
	err = snB.SetStr(b)
	if err != nil {
		return "", err
	}

	div := Div(ptA, snB)
	ret = div.String()
	return ret, nil
}

//-------------------------------------------------------------------------------//
// Tum

func (t *Snum) Add(sn *Snum) {
	t.decimal.Add(t.decimal, sn.decimal)
}

func (t *Snum) AddStr(sn string) (err error) {
	snA := &Snum{}
	snA.Init(0, 0)
	err = snA.SetStr(sn)
	if err != nil {
		return err
	}

	t.Add(snA)
	return nil
}

func (t *Snum) Sub(sn *Snum) {
	t.decimal.Sub(t.decimal, sn.decimal)
}

func (t *Snum) SubStr(sn string) (err error) {
	ptB := &Snum{}
	ptB.Init(0, 0)
	err = ptB.SetStr(sn)
	if err != nil {
		return err
	}

	t.Sub(ptB)
	return nil
}

func (t *Snum) Mul(sn *Snum) {
	t.decimal.Mul(t.decimal, sn.decimal)
}

func (t *Snum) MulStr(sn string) (err error) {
	snA := &Snum{}
	snA.Init(0, 0)
	err = snA.SetStr(sn)
	if err != nil {
		return err
	}

	t.Mul(snA)
	return nil
}

func (t *Snum) Div(sn *Snum) {
	t.decimal.Quo(t.decimal, sn.decimal)

	if t.decimal.IsNormal() == true {
		t.decimal.Quantize(DEF_lenExtendDecimalForCalc)
	}
}

func (t *Snum) DivStr(sn string) (err error) {
	snB := &Snum{}
	snB.Init(0, 0)
	err = snB.SetStr(sn)

	t.Div(snB)
	return nil
}
