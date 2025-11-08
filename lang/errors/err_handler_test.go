package errors

import (
	"fmt"
	"testing"
)

func TestHandleError(t *testing.T) {

	err := fmt.Errorf("a example of error")

	HandleError(err)
}

func TestHandlePanic(t *testing.T) {

	defer func() {
		x := recover()
		HandlePanic(x)
	}()

	panic("a example of panic")
}
