package domain

type FullPurchase struct {
	First    *Purchase
	Second   *Purchase
	Increase Money
}

func (f *FullPurchase) Purchase(s Structure, finance Finance) {
	if s.BuiltOn.RetainAlways {
		f.First = NewPurchase(s.Name, s.Cost, s.Income, finance)
		f.First.Calculate()

		f.Second = NewPurchase(s.BuiltOn.Name, s.BuiltOn.Cost, s.BuiltOn.Income quantity, f.First.FinanceOut}
		f.Second.Calculate()

		f.Increase = f.First.Increase
	} else {
		f.First := NewPurchase(s.BuiltOn.Name, s.BuiltOn.Cost, s.BuiltOn.Income, finance)
		f.First.Calculate()

		f.Second = NewPurchase(s.Name, s.Cost, s.Income, land.Finance)
		f.Second.Calculate()

		f.Increase = f.Second.Increase
	}
}
