package factory_method

type Maverick struct {
	Gun
}

func NewMaverick() IGun {
	return &Maverick{
		Gun{
			Name:  "Maverick gun",
			Power: 5,
		},
	}
}
