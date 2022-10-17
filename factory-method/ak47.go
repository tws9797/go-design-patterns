package factory_method

type Ak47 struct {
	Gun
}

func NewAk47() IGun {
	return &Ak47{
		Gun{
			Name:  "AK 47 Gun",
			Power: 4,
		},
	}
}
