package buckets

type BucketHolder interface {
	GetBucket() (Bucket, error)
}

type BucketHolderRegistration struct {
	Name string

	Enabled bool

	Priority int

	Holder BucketHolder
}

type BucketHolderRegistry interface {
	ListRegistrations() []*BucketHolderRegistration
}
