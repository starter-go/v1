package buckets

import "context"

type BucketContext struct {
	Context context.Context // the global-context

	Driver Driver

	DriverManager DriverManager

	Configuration Configuration

	Bucket Bucket
}
