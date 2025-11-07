package ossdriver

import (
	"context"
	"fmt"
	"net/url"
	"strings"

	"github.com/aliyun/alibabacloud-oss-go-sdk-v2/oss"
	"github.com/starter-go/v1/buckets"
)

////////////////////////////////////////////////////////////////////////////////
// 参考:
//  https://help.aliyun.com/zh/oss/developer-reference/

////////////////////////////////////////////////////////////////////////////////

type innerOSSObject struct {
	context context.Context
	bucket  *innerOSSBucket
	name    buckets.ObjectName // the object-name (key)
}

func (inst *innerOSSObject) _impl() buckets.Object {
	return inst
}

func (inst *innerOSSObject) Name() buckets.ObjectName {
	return inst.name
}

func (inst *innerOSSObject) GetContext() context.Context {
	return inst.context
}

func (inst *innerOSSObject) WithContext(cc context.Context) buckets.Object {
	if cc == nil {
		cc = context.Background()
	}
	inst.context = cc
	return inst
}

func (inst *innerOSSObject) GetBucket() buckets.Bucket {
	return inst.bucket
}

func (inst *innerOSSObject) Exists() (bool, error) {

	bucketName := inst.bucket.name
	client := inst.bucket.client
	ctx := inst.context
	objectKey := string(inst.name)

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

	client := inst.bucket.client
	ctx := inst.context
	objectKey := string(inst.name)
	bucketName := inst.bucket.name

	req := new(oss.DeleteObjectRequest)
	req.Bucket = &bucketName
	req.Key = &objectKey

	_, err := client.DeleteObject(ctx, req)
	return err
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
	dr.Name = "aliyun/oss"
	dr.Priority = 0
	return dr
}

func (inst *Driver) OpenBucket(ctx *buckets.BucketContext) (buckets.Bucket, error) {

	if ctx == nil {
		return nil, fmt.Errorf("param: context is nil")
	}
	if ctx.Context == nil {
		ctx.Context = context.Background()
	}

	cfg, err := inst.prepareConfig(ctx)
	if err != nil {
		return nil, err
	}

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

func (inst *Driver) prepareConfig(ctx *buckets.BucketContext) (*innerOSSClientConfig, error) {

	src := &ctx.Configuration
	dst := new(innerOSSClientConfig)

	uriResolver := new(innerBucketUriResolver)
	err := uriResolver.Resolve(src.URI)
	if err != nil {
		return nil, err
	}

	dst.Region = uriResolver.region
	dst.Endpoint = uriResolver.endpoint
	dst.BucketName = uriResolver.bucket
	dst.AccessKeyID = src.KeyID
	dst.AccessKeySecret = src.KeySecret

	return dst, nil
}

func (inst *Driver) openAliyunOSSClient(config *innerOSSClientConfig) (*oss.Client, error) {
	// Load the default configuration for OSS
	cfg := oss.LoadDefaultConfig()

	// Create and return a new OSS client
	client := oss.NewClient(cfg)
	return client, nil
}

////////////////////////////////////////////////////////////////////////////////

type innerBucketUriResolver struct {
	region   string
	endpoint string
	bucket   string // the bucket name

}

func (inst *innerBucketUriResolver) Resolve(uri string) error {

	u, err := url.Parse(uri)
	if err != nil {
		return err
	}

	parts := strings.SplitN(uri, "?", 2)

	query := u.Query()
	inst.bucket = query.Get("bucket")
	inst.region = query.Get("region")
	inst.endpoint = parts[0]

	if inst.bucket == "" {
		return fmt.Errorf("buckets: oss.config [bucket] is empty")
	}
	if inst.region == "" {
		return fmt.Errorf("buckets: oss.config [region] is empty")
	}
	if inst.endpoint == "" {
		return fmt.Errorf("buckets: oss.config [endpoint] is empty")
	}

	return nil
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
	context *buckets.BucketContext
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

func (inst *innerOSSBucket) GetContext() context.Context {
	return inst.context.Context
}

func (inst *innerOSSBucket) GetBucketContext() *buckets.BucketContext {
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
