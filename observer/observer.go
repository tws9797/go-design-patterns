package observer

// https://refactoring.guru/design-patterns/observer
// https://golangbyexample.com/observer-design-pattern-golang/
// Allows an instance (called subject) to publish events to multiple instances (called observers)
// Observers subscribe to subject, notified by events of any state change in subject

type Observer interface {
	Update(string)
	GetId() string
}
