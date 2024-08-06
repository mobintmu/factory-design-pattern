package main

import (
	"go-example/factory/manufacture"
)

func main() {

	markCompany := manufacture.NewMarkCompany()
	for i := 0; i <= 10; i++ {
		markCompany.AddGun()
	}
	markCompany.PrintGunList()

}
