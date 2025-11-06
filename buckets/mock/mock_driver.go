package mock

import "github.com/starter-go/v1/buckets"

type TheMockDriver struct {
}

func (inst *TheMockDriver) _impl() buckets.Driver {
	return inst
}
