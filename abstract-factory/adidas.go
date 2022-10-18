package abstract_factory

type Adidas struct{}

func (a *Adidas) MakeShoe() IShoe {
	return &AdidasShoe{
		Shoe{
			Logo: "Adidas",
			Size: 14,
		},
	}
}

func (a *Adidas) MakeShort() IShort {
	return &AdidasShort{
		Short{
			Logo: "Adidas",
			Size: 14,
		},
	}
}
