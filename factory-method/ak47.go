package factory_method

// Concrete class that implements the IGun interface
// Embed with Gun struct that implements the IGun interface

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
