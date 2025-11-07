package buckets

import (
	"context"
	"io"
)

type ObjectName string

type ObjectMeta struct {
	ObjectSum

	Name ObjectName

	Length int64

	ContentType string
}

type ObjectData struct {
	Length int64

	Closer io.Closer

	Reader io.Reader
}

type ObjectSum struct {
	Algorithm HashAlgorithm
	Sum       []byte
}

type Object interface {
	GetContext() context.Context

	WithContext(cc context.Context) Object

	Name() ObjectName

	GetBucket() Bucket

	Fetch(ctx *FetchContext) error

	Put(ctx *PutContext) error

	Remove() error

	Exists() (bool, error)
}
