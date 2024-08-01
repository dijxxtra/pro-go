package store

type SpecialDeal struct {
	Name string
	*Product
	price float64
}

func NewSpecialDeal(name string, p *Product, discount float64) *SpecialDeal {
	return &SpecialDeal{name, p, p.price - discount}
}

func (s *SpecialDeal) GetDetails() (string, float64, float64) {
	return s.Name, s.price, s.Price(0)
}
