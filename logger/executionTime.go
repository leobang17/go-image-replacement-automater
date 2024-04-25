package logger

import (
	"context"
	"errors"
	"fmt"
	"time"
)

var startTime time.Time
var cancel context.CancelFunc
var ctx context.Context

func StartTimer() {
	ctx, cancel = context.WithCancel(context.Background()) 
	startTime = time.Now()
	
	go func() {
		for {
			select {
			case <- ctx.Done():
				return
			default:
				timePassed := time.Since(startTime)
				fmt.Printf("\rProgram has been running for: %.2f seconds", timePassed.Seconds())
				time.Sleep(10 * time.Millisecond) 
			}
		}
	}()
}

func EndTimer() {
	if cancel == nil || ctx == nil {
		panic(errors.New("you must call StartTimer() before EndTimer()"))
	}
	cancel() 
	timePassed := time.Since(startTime)
	fmt.Printf("\nProgram ran for a total of: %.2f seconds\n", timePassed.Seconds())
	cleanup()
}

func cleanup() {
	ctx = nil
	cancel = nil
	startTime = time.Time{}
}