package logger

import (
	"fmt"
	"testing"
	"time"
)


func Test_ExecutionTime_success(t *testing.T) {
	StartTimer()
	fmt.Println("print1")
	fmt.Println("print2")
	fmt.Println("print3")
	fmt.Println("print4")
	fmt.Println("print5")
	time.Sleep(3 * time.Second)
	EndTimer()
}
func Test_ExecutionTime_failure(t *testing.T) {
	defer func() {
		if r := recover(); r == nil {
			t.Errorf("should call panic but didn't")
		} 
	}()
	EndTimer()
}