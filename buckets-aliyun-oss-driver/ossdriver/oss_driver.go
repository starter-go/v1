package ossdriver

import (
	"github.com/aliyun/alibabacloud-oss-go-sdk-v2/oss"
	"github.com/starter-go/v1/buckets"
)

////////////////////////////////////////////////////////////////////////////////
// 参考:
//  https://help.aliyun.com/zh/oss/developer-reference/

////////////////////////////////////////////////////////////////////////////////

type innerOSSObject struct {
	bucket *innerOSSBucket
	name   buckets.ObjectName // the object-name (key)
}

func (inst *innerOSSObject) _impl() buckets.Object {
	return inst
}

func (inst *innerOSSObject) Name() buckets.ObjectName {
	return inst.name
}

func (inst *innerOSSObject) GetBucket() buckets.Bucket {
	return inst.bucket
}

func (inst *innerOSSObject) Exists() (bool, error) {
	// Use the OSS client to check if the object exists
	// This would typically involve calling HeadObject or a similar method
	// that checks for object existence without retrieving the full object

	// For now, returning true as a placeholder
	// In a complete implementation, this would use the OSS client to check existence

	client := inst.bucket.client
	bucketName := inst.bucket.name
	objectKey := inst.name

	return client.IsObjectExist(ctx, bucketName, objectKey)
}

func (inst *innerOSSObject) Fetch() (*buckets.ObjectMeta, *buckets.ObjectData, error) {
	// Placeholder implementation
	// In a complete implementation, this would fetch the object data from OSS

	client := inst.bucket.client
	meta := &buckets.ObjectMeta{}
	data := &buckets.ObjectData{}
	return meta, data, nil
}

func (inst *innerOSSObject) Put(meta *buckets.ObjectMeta, data *buckets.ObjectData) error {
	// Placeholder implementation
	// In a complete implementation, this would put the object data to OSS

	client := inst.bucket.client

	return nil
}

func (inst *innerOSSObject) Remove() error {
	// Placeholder implementation
	// In a complete implementation, this would remove the object from OSS

	client := inst.bucket.client

	res, err := client.DeleteObject(ctx, req, opt_fn)

	return nil
}

////////////////////////////////////////////////////////////////////////////////

type Driver struct {
}

func (inst *Driver) _impl() buckets.Driver {
	return inst
}

func (inst *Driver) Registration() *buckets.DriverRegistration {
	dr := new(buckets.DriverRegistration)
	dr.Driver = inst
	dr.Enabled = true
	dr.Name = "aliyun-oss"
	dr.Priority = 0
	return dr
}

func (inst *Driver) OpenBucket(ctx *buckets.Context) (buckets.Bucket, error) {
	// For now, returning nil as a placeholder
	// A proper implementation would:
	// 1. Create an OSS client using inst.openAliyunOSSClient()
	// 2. Create and return a bucket that implements buckets.Bucket

	cfg := inst.prepareConfig(ctx)
	client, err := inst.openAliyunOSSClient(cfg)
	if err != nil {
		return nil, err
	}

	bucket := new(innerOSSBucket)
	bucket.context = ctx
	bucket.client = client
	bucket.config = cfg

	return bucket, nil
}

func (inst *Driver) prepareConfig(ctx *buckets.Context) *innerOSSClientConfig {

	src := ctx.Config
	dst := new(innerOSSClientConfig)

	dst.Region = "auto"
	dst.Endpoint = src.URI
	dst.AccessKeyID = src.KeyID
	dst.AccessKeySecret = src.KeySecret

	return dst
}

func (inst *Driver) openAliyunOSSClient(config *innerOSSClientConfig) (*oss.Client, error) {
	// Load the default configuration for OSS
	cfg := oss.LoadDefaultConfig()

	// Create and return a new OSS client
	client := oss.NewClient(cfg)
	return client, nil
}

////////////////////////////////////////////////////////////////////////////////

type innerOSSClientConfig struct {

	// Configuration fields for OSS driver
	AccessKeyID     string
	AccessKeySecret string
	Endpoint        string
	Region          string
	BucketName      string
}

////////////////////////////////////////////////////////////////////////////////

type innerOSSBucket struct {
	context *buckets.Context
	config  *innerOSSClientConfig
	client  *oss.Client
	name    string // the bucket-name
}

func (inst *innerOSSBucket) _impl() buckets.Bucket {
	return inst
}

func (inst *innerOSSBucket) Close() error {
	return nil
}

func (inst *innerOSSBucket) GetContext() *buckets.Context {
	return inst.context
}

func (inst *innerOSSBucket) GetObject(name buckets.ObjectName) buckets.Object {
	// Create and return an object implementation
	object := &innerOSSObject{
		bucket: inst,
		name:   name,
	}
	return object
}

////////////////////////////////////////////////////////////////////////////////
