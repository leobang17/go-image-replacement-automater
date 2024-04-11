package atomicity

import "fmt"

type AtomicityError[T any] struct {
	Logs []T
}

func (e *AtomicityError[T]) Error() string {
	return fmt.Sprintln("Error occured during atomic section")
}