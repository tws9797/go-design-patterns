package factory_method

import "fmt"

// Factory method - returns an object of that class type (Products)
// Factory method should have its return type declared as the common interface

func GetGun(gunType string) (IGun, error) {
	if gunType == "ak47" {
		return NewAk47(), nil
	}
	if gunType == "maverick" {
		return NewMaverick(), nil
	}
	return nil, fmt.Errorf("wrong gun type passed")
}
