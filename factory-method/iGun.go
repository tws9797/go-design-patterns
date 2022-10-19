package factory_method

// https://refactoring.guru/design-patterns/factory-method
// https://golangbyexample.com/golang-factory-design-pattern/
// Factory method is a creational design pattern that provides an interface for creating objects
// Allows subclasses to alter the type of objects that will be created

type IGun interface {
	SetName(name string)
	SetPower(power int)
	GetName() string
	GetPower() int
}
