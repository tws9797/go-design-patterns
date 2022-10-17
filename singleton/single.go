package singleton

import (
	"fmt"
	"sync"
)

// https://refactoring.guru/design-patterns/singleton/go/example#example-0
// https://golangbyexample.com/singleton-design-pattern-go/
// Singleton ensure that a class can only have one instance
// while providing global access point to this instance

var lock = &sync.Mutex{}

type Single struct{}

var singleInstance *Single

func GetInstance() *Single {

	// There is a nil-check at the start for making sure singleInstance is empty first time around.
	// This is to prevent expensive lock operations every time the getinstance method is called.
	// If this check fails, then it means that the singleInstance field is already populated.
	if singleInstance == nil {
		lock.Lock()
		defer lock.Unlock()

		// There is another nil-check after the lock is acquired.
		// This is to ensure that if more than one goroutine bypasses the first check, only one goroutine can create the singleton instance.
		// Otherwise, all goroutines will create their own instances of the singleton struct.
		if singleInstance == nil {
			fmt.Println("Creating single instance now.")

			// The singleInstance struct is created within the lock.
			singleInstance = &Single{}
		} else {
			fmt.Println("Single instance already created-1.")
		}
	} else {
		fmt.Println("Single instance already created-2.")
	}

	return singleInstance
}
