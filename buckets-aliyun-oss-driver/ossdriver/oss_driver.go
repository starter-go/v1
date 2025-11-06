package ossdriver

import (
	"github.com/aliyun/alibabacloud-oss-go-sdk-v2/oss"
	"github.com/starter-go/v1/buckets"
)

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

	client, err := inst.openAliyunOSSClient()

	//todo ...

	return nil, nil

}

func (inst *Driver) openAliyunOSSClient() (*oss.Client, error) {

	cfg := oss.LoadDefaultConfig()

	return nil, nil
}
