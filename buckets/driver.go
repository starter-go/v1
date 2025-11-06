package buckets

type Driver interface {
	OpenBucket(ctx *Context) (Bucket, error)
}

type DriverManager interface {
	OpenBucket(ctx *Context) (Bucket, error)

	GetDriver(name string) (Driver, error)
}
