package singleton

import (
	"fmt"
	"sync"
)

// The sync.Once will only perform the operation only once. See below code
var once sync.Once

func GetInstance2() *Single {
	if singleInstance == nil {
		once.Do(func() {
			fmt.Println("Creating Single Instance Now")
			singleInstance = &Single{}
		})
	} else {
		fmt.Println("Single Instance already created-2")
	}

	return singleInstance
}
