package buckets

import "io"

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
	Name() ObjectName

	GetBucket() Bucket

	Fetch() (*ObjectMeta, *ObjectData, error)

	Put(meta *ObjectMeta, data *ObjectData) error

	Remove() error

	Exists() (bool, error)
}
