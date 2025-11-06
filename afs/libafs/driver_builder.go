package libafs

import (
	"github.com/starter-go/v1/afs"
	"github.com/starter-go/v1/afs/implementation"
)

type DriverBuilder struct {
	PlatformAPI implementation.PlatformAPI
}

func (inst *DriverBuilder) Create() afs.Driver {

	ctx := new(implementation.Context)
	driver := new(innerCommonDriver)
	fs := new(innerCommonFS)
	fapi := new(innerCommonFullAPIImpl)
	papi := inst.PlatformAPI

	if papi == nil {
		panic("DriverBuilder: PlatformAPI is nil")
	}

	ctx.Driver = driver
	ctx.FS = fs
	ctx.PlatformAPI = papi
	ctx.FullAPI = fapi

	return ctx.Driver
}
