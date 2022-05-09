package golanggeneric

import (
	"fmt"
	"testing"
)

func MultipleParameter[T1 any, T2 any](param1 T1, param2 T2) {
	fmt.Println(param1, param2)
	fmt.Println(param1)
}


func TestMultipleParameter(t *testing.T) {
	MultipleParameter("Hello", 10)
	MultipleParameter(100, "Hello")
}