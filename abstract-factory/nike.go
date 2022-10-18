package abstract_factory

type Nike struct{}

func (n *Nike) MakeShoe() IShoe {
	return &NikeShoe{
		Shoe{
			Logo: "Nike",
			Size: 14,
		},
	}
}

func (n *Nike) MakeShort() IShort {
	return &NikeShort{
		Short{
			Logo: "Nike",
			Size: 14,
		},
	}
}
