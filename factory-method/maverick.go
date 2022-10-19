package factory_method

// Concrete class that implements the IGun interface
// Embed with Gun struct that implements the IGun interface

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
