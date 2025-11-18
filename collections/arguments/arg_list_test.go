package arguments

import (
	"os"
	"testing"
)

func TestArgList(t *testing.T) {

	args := os.Args

	list := NewList(args)

	str := list.String()
	t.Logf("arguments = %s", str)
	t.Logf("detail:")

	list.ListItemsWithFilter(func(a *Argument) bool {
		index := a.Index()
		value := a.Value()
		t.Logf("index:%d value:%s", index, value)
		return false
	})

	// for  ;;  {}

}
