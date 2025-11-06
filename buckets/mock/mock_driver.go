package mock

import (
	"errors"

	"github.com/starter-go/v1/buckets"
)

var ErrObjectNotFound = errors.New("object not found")

// TheMockDriver is a mock implementation of the buckets.Driver interface
type TheMockDriver struct {
}

// _impl ensures TheMockDriver implements the buckets.Driver interface
func (inst *TheMockDriver) _impl() buckets.Driver {
	return inst
}

// Registration returns the driver registration information
func (inst *TheMockDriver) Registration() *buckets.DriverRegistration {
	return &buckets.DriverRegistration{
		Name:     "mock",
		Enabled:  true,
		Priority: 0,
		Driver:   inst,
	}
}

// OpenBucket opens and returns a mock bucket
func (inst *TheMockDriver) OpenBucket(ctx *buckets.Context) (buckets.Bucket, error) {
	return &mockBucket{
		context: ctx,
		objects: make(map[buckets.ObjectName]*mockObject),
	}, nil
}

// mockBucket is a mock implementation of the buckets.Bucket interface
type mockBucket struct {
	context *buckets.Context
	objects map[buckets.ObjectName]*mockObject
}

// Close closes the bucket
func (b *mockBucket) Close() error {
	return nil
}

// GetContext returns the bucket context
func (b *mockBucket) GetContext() *buckets.Context {
	return b.context
}

// GetObject returns a mock object
func (b *mockBucket) GetObject(name buckets.ObjectName) buckets.Object {
	if obj, exists := b.objects[name]; exists {
		return obj
	}

	obj := &mockObject{
		bucket: b,
		name:   name,
		data:   nil,
	}
	b.objects[name] = obj
	return obj
}

// mockObject is a mock implementation of the buckets.Object interface
type mockObject struct {
	bucket *mockBucket
	name   buckets.ObjectName
	data   []byte
}

// Name returns the object name
func (o *mockObject) Name() buckets.ObjectName {
	return o.name
}

// GetBucket returns the parent bucket
func (o *mockObject) GetBucket() buckets.Bucket {
	return o.bucket
}

// Fetch retrieves the object data
func (o *mockObject) Fetch() (*buckets.ObjectMeta, *buckets.ObjectData, error) {
	if o.data == nil {
		return nil, nil, ErrObjectNotFound
	}

	meta := &buckets.ObjectMeta{
		Name:        o.name,
		Length:      int64(len(o.data)),
		ContentType: "application/octet-stream",
	}

	data := &buckets.ObjectData{
		Length: int64(len(o.data)),
		Closer: nil,
		Reader: nil, // In a real implementation, this would be a reader for the data
	}

	return meta, data, nil
}

// Put stores the object data
func (o *mockObject) Put(meta *buckets.ObjectMeta, data *buckets.ObjectData) error {
	// In a real implementation, we would read the data from the reader
	// For this mock, we'll just store a placeholder
	o.data = make([]byte, data.Length)
	return nil
}

// Remove deletes the object
func (o *mockObject) Remove() error {
	o.data = nil
	return nil
}

// Exists checks if the object exists
func (o *mockObject) Exists() (bool, error) {
	return o.data != nil, nil
}
