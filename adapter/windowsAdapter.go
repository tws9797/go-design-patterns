package adapter

import "fmt"

// https://refactoring.guru/design-patterns/adapter
// https://golangbyexample.com/adapter-design-pattern-go/
// Allows objects with incompatible interfaces to collaborate

type WindowsAdapter struct {
	windowMachine *Windows
}

func (w *WindowsAdapter) InsertIntoLightningPort() {
	fmt.Println("Adapter converts Lightning signal to USB.")
	w.windowMachine.InsertIntoUSBPort()
}
