package domain

type FullPurchase struct {
	First  *Purchase
	Second *Purchase
	//	Structure *Purchase
	Increase Money
	Hours    int
}

func (f *FullPurchase) Purchase(s Structure, finance Finance) {
	if s.BuiltOn.RetainAlways {
		f.First = NewPurchase(s.Name, s.Cost, s.Income, finance)
		f.First.Calculate()

		f.Second = NewPurchase(s.BuiltOn.Name, s.BuiltOn.Cost, s.BuiltOn.Income, f.First.FinanceOut)
		f.Second.Calculate()

		f.Increase = f.First.Increase
		f.Hours = f.First.Hours
	} else {
		f.First = NewPurchase(s.BuiltOn.Name, s.BuiltOn.Cost, s.BuiltOn.Income, finance)
		f.First.Calculate()

		f.Second = NewPurchase(s.Name, s.Cost, s.Income, f.First.FinanceOut)
		f.Second.Calculate()

		f.Increase = f.Second.Increase
		f.Hours = f.Second.Hours
	}
}
