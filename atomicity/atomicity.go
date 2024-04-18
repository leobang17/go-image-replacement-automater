package atomicity

import (
	"context"
	"fmt"
	"sync"
)

// TODO: refactor
func NewAtomicSection[T any]() AtomicSection[T] {	
	ctx, cancel := context.WithCancel(context.Background())
	var (
		successChan chan string = make(chan string)
		errorChan chan error = make(chan error)
		wg *sync.WaitGroup = &sync.WaitGroup{}
		mutex sync.Mutex = sync.Mutex{}
		logs []T = []T {}
	)

	var runnable Runnable[T] = func(do Do[T]) {
		wg.Add(1)
		go func() {
			defer wg.Done()
			defer func() {
				if recover := recover(); recover != nil {
					cancel()
					errorChan <- fmt.Errorf("panic occured during 'do'. %v", recover)
				}
			}()

			select {
			case <- ctx.Done():
				return
			default:
				result, err := do()
				if err != nil {
					cancel()
					errorChan <- err
				}

				// critical section
				mutex.Lock()
				logs = append(logs, result)
				mutex.Unlock()
			}
		}()
	}

	var Section AtomicSection[T] = func(execAtomic ExecuteAtomicTask[T]) ([]T, *AtomicityError[T]) {
		defer func() {
			close(errorChan)
			close(successChan)
		}()

		execAtomic(runnable)

		go func() {
			wg.Wait()
			select {
			case <- ctx.Done():
				return 
			default:
				successChan <- "success"
			}
		}()

		select {
		case <- errorChan:
			return []T{}, &AtomicityError[T]{ Logs: logs }
		case <- successChan:
			return logs, nil
		}
	}
	return Section
}