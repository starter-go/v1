package buckets

import (
	"context"
	"io"
)

type Bucket interface {
	io.Closer

	GetContext() context.Context

	GetBucketContext() *BucketContext

	GetObject(name ObjectName) Object
}
