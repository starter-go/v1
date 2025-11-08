package ios

import (
	"io"

	"github.com/starter-go/v1/lang/errors"
)

func Close(cl io.Closer) {
	if cl == nil {
		return
	}
	err := cl.Close()
	errors.HandleError(err)
}
