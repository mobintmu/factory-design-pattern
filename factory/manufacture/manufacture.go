package manufacture

import (
	"fmt"
	"go-example/factory/gun"
	"math/rand"
)

type MarkCompany struct {
	gunList []gun.IGun
}

func NewMarkCompany() *MarkCompany {
	mk := MarkCompany{}
	mk.gunList = make([]gun.IGun, 0)
	return &mk
}

func (m *MarkCompany) AddGun() {
	myRand := rand.Intn(2)
	if myRand == 0 {
		g := gun.NewAK47()
		m.gunList = append(m.gunList, g)
	} else {
		g := gun.NewShotgun()
		m.gunList = append(m.gunList, g)
	}
}

func (m *MarkCompany) PrintGunList() {
	for _, g := range m.gunList {
		fmt.Println(g.GetName(), g.GetPower())
	}
}
