package abstract_factory

import "fmt"

// https://refactoring.guru/design-patterns/abstract-factory
// https://golangbyexample.com/abstract-factory-design-pattern-go/
// Produce families of related objects without specifying their concrete class

type ISportFactory interface {
	MakeShoe() IShoe
	MakeShort() IShort
}

func GetSportsFactory(brand string) (ISportFactory, error) {
	if brand == "adidas" {
		return &Adidas{}, nil
	}

	if brand == "nike" {
		return &Nike{}, nil
	}

	return nil, fmt.Errorf("wrong brand type passed")
}
