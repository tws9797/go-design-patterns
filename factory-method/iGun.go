package factory_method

// https://refactoring.guru/design-patterns/factory-method
// https://golangbyexample.com/golang-factory-design-pattern/
// Provides an interface for creating an objects in a superclass
// but allows subclasses to alter the type of objects that will be created

type IGun interface {
	SetName(name string)
	SetPower(power int)
	GetName() string
	GetPower() int
}
