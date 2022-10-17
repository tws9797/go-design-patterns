package factory_method

// https://golangbyexample.com/golang-factory-design-pattern/
// Provides an interface for creating an objects in a superclass
// but allows subclasses to alter the type of objects that will be created

type IGun interface {
	setName(name string)
	setPower(power int)
	getName() string
	getPower() int
}
