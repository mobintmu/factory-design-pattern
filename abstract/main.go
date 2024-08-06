package main

import (
	"go-example/abstract/factory"
	"math/rand"
)

func randRange(min, max int) int {
	return rand.Intn(max-min) + min
}

func main() {

	var listFactory []factory.ISportFactory = make([]factory.ISportFactory, 5)

	for range listFactory {

		random := randRange(0, 10)
		if random < 5 {
			adidas := factory.NewAdidas()
			listFactory = append(listFactory, adidas)
		} else {
			nike := factory.NewNike()
			listFactory = append(listFactory, nike)
		}
	}

	for _, factory_value := range listFactory {
		if factory_value != nil {
			println("___________>>>>>>>>>_________")
			println("logo shirt : ", factory_value.GetShirt().GetLogo())
			println("logo shoe", factory_value.GetShoe().GetLogo())
			println("___________<<<<<<<<<<<_________")
		}

	}
}
