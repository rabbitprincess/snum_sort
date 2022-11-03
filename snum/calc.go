package snum

func (t *Snum) Add(sn *Snum) {
	t.decimal.Add(t.decimal, sn.decimal)
}

func (t *Snum) Sub(sn *Snum) {
	t.decimal.Sub(t.decimal, sn.decimal)
}

func (t *Snum) Mul(sn *Snum) {
	t.decimal.Mul(t.decimal, sn.decimal)
}

func (t *Snum) Div(sn *Snum) {
	t.decimal.Quo(t.decimal, sn.decimal)

	if t.decimal.IsNormal() == true {
		t.decimal.Quantize(20)
	}
}
