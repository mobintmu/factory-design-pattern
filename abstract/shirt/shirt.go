package shirt

type IShirt interface {
	SetLogo(logo string)
	GetLogo() string
}

type Shirt struct {
	logo string
}

func NewShirt(logo string) Shirt {
	return Shirt{
		logo: logo,
	}
}

func (s *Shirt) SetLogo(logo string) {
	s.logo = logo
}

func (s *Shirt) GetLogo() string {
	return s.logo
}
