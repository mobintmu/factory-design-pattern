package factory

import (
	"go-example/abstract/shirt"
	"go-example/abstract/shoe"
)

type Adidas struct {
	shoe  shoe.IShoe
	shirt shirt.IShirt
}

type AdidasShoe struct {
	shoe.Shoe
}

type AdidasShirt struct {
	shirt.Shirt
}

func NewAdidas() *Adidas {
	adidas := Adidas{}
	adidas.shoe = adidas.MakeShoe()
	adidas.shirt = adidas.MakeShirt()
	return &adidas
}

func (a *Adidas) MakeShoe() shoe.IShoe {
	return &AdidasShoe{
		shoe.NewShoe("adidas", 20),
	}
}

func (a *Adidas) MakeShirt() shirt.IShirt {
	return &AdidasShirt{
		shirt.NewShirt("adidas"),
	}
}

func (a *Adidas) GetShoe() shoe.IShoe {
	return a.shoe
}

func (a *Adidas) GetShirt() shirt.IShirt {
	return a.shirt
}
