package builder

// https://refactoring.guru/design-patterns/builder/go/example
// https://golangbyexample.com/builder-pattern-golang/
// Used for constructing complex object

type IBuilder interface {
	SetWindowType()
	SetDoorType()
	SetNumFloor()
	GetHouse() House
}

func GetBuilder(builderType string) IBuilder {
	if builderType == "normal" {
		return NewNormalBuilder()
	}

	if builderType == "igloo" {
		return NewIglooBuilder()
	}
	return nil
}
