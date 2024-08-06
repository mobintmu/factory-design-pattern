package gun

type Shotgun struct {
	Gun
}

func NewShotgun() IGun {
	return &Shotgun{
		Gun: Gun{
			name:  "Shotgun",
			power: 10,
		},
	}
}
