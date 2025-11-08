package ios

import (
	"fmt"
	"testing"
)

func TestClose(t *testing.T) {

	bb := new(innerTestCloseBuffer)
	Close(bb)

}

type innerTestCloseBuffer struct {
}

func (inst *innerTestCloseBuffer) Close() error {

	return fmt.Errorf("a mock i/o error")
}
