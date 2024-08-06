package factory

import (
	"go-example/abstract/shirt"
	"go-example/abstract/shoe"
)

type Nike struct {
	shoe  shoe.IShoe
	shirt shirt.IShirt
}

type NikeShoe struct {
	shoe.Shoe
}

type NikeShirt struct {
	shirt.Shirt
}

func NewNike() *Nike {
	nike := Nike{}
	nike.shirt = nike.MakeShirt()
	nike.shoe = nike.MakeShoe()
	return &nike
}

func (n *Nike) MakeShoe() shoe.IShoe {
	return &NikeShoe{
		shoe.NewShoe("nike", 20),
	}
}

func (n *Nike) MakeShirt() shirt.IShirt {
	return &NikeShirt{
		shirt.NewShirt("nike"),
	}
}

func (n *Nike) GetShoe() shoe.IShoe {
	return n.shoe
}

func (n *Nike) GetShirt() shirt.IShirt {
	return n.shirt
}
