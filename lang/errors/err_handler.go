package errors

import (
	"fmt"
	"os"
)

////////////////////////////////////////////////////////////////////////////////
// interfaces

type ErrorHandler interface {
	HandleError(err error)
}

type PanicHandler interface {
	HandlePanic(p any)
}

////////////////////////////////////////////////////////////////////////////////
// vars

var theDefaultErrorHandler ErrorHandler

var theDefaultPanicHandler PanicHandler

////////////////////////////////////////////////////////////////////////////////
// functions

func SetDefaultErrorHandler(h ErrorHandler) {
	if h == nil {
		return
	}
	theDefaultErrorHandler = h
}

func SetDefaultPanicHandler(h PanicHandler) {
	if h == nil {
		return
	}
	theDefaultPanicHandler = h
}

func GetDefaultErrorHandler() ErrorHandler {

	h := theDefaultErrorHandler
	if h == nil {
		h = new(innerDefaultErrorHandler)
		theDefaultErrorHandler = h

	}
	return h
}

func GetDefaultPanicHandler() PanicHandler {
	h := theDefaultPanicHandler
	if h == nil {
		h = new(innerDefaultPanicHandler)
		theDefaultPanicHandler = h
	}
	return h
}

func HandleError(err error) {
	GetDefaultErrorHandler().HandleError(err)
}

func HandlePanic(p any) {
	GetDefaultPanicHandler().HandlePanic(p)
}

////////////////////////////////////////////////////////////////////////////////

type innerDefaultErrorHandler struct{}

func (inst *innerDefaultErrorHandler) HandleError(err error) {
	if err == nil {
		return
	}
	fmt.Fprintf(os.Stderr, "%v", err.Error())
}

////////////////////////////////////////////////////////////////////////////////

type innerDefaultPanicHandler struct{}

func (inst *innerDefaultPanicHandler) HandlePanic(x any) {

	if x == nil {
		return
	}

	err, ok := x.(error)
	if ok && (err != nil) {
		err = fmt.Errorf("panic: %v", err.Error())
		HandleError(err)
		return
	}

	msg, ok := x.(string)
	if ok {
		err = fmt.Errorf("panic: %v", msg)
		HandleError(err)
		return
	}

	err = fmt.Errorf("panic(unknown): %v", x)
	HandleError(err)
}

////////////////////////////////////////////////////////////////////////////////
