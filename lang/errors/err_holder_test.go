package errors

import (
	"fmt"
	"testing"
)

func TestSimpleErrorHolder(t *testing.T) {

	holder := NewErrorHolder()

	holder.HandleError(fmt.Errorf("err-1"))
	holder.HandleError(fmt.Errorf("err-2"))
	holder.HandleError(fmt.Errorf("err-3"))
	holder.HandleError(fmt.Errorf("err-4"))

	all := holder.All()
	for idx, err := range all {
		fmt.Printf("\t error[%d] = %s \n", idx, err.Error())
	}

}

func TestErrorList(t *testing.T) {

	holder := NewErrorList()

	holder.HandleError(fmt.Errorf("err-1"))
	holder.HandleError(fmt.Errorf("err-2"))
	holder.HandleError(fmt.Errorf("err-3"))
	holder.HandleError(fmt.Errorf("err-4"))

	all := holder.All()
	for idx, err := range all {
		fmt.Printf("\t error[%d] = %s \n", idx, err.Error())
	}

}
