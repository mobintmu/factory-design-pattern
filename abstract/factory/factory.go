package factory

import (
	"go-example/abstract/shirt"
	"go-example/abstract/shoe"
)

type ISportFactory interface {
	MakeShoe() shoe.IShoe
	MakeShirt() shirt.IShirt
	GetShoe() shoe.IShoe
	GetShirt() shirt.IShirt
}
