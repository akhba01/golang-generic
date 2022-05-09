package golanggeneric

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Length[T any](param T) T {
	fmt.Println(param)
	return param
}

func TestLength(t *testing.T) {
	var result string = Length("Hello")
	assert.Equal(t, "Hello", result)

	var result2 int = Length(10)
	assert.Equal(t, 10, result2)
}


func LengthGeneral(param interface{}) interface{} {
	fmt.Println(param)
	return param
}
func TestSample(t *testing.T) {
	assert.False(t, false)

	assert.Equal(t,10,LengthGeneral(10))
}
