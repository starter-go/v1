package buckets

import "io"

type Bucket interface {
	io.Closer

	GetContext() *Context

	GetObject(name ObjectName) Object
}
