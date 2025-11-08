package libafs

import (
	"github.com/starter-go/v1/afs"
	"github.com/starter-go/v1/afs/implementation"
)

type DriverBuilder struct {
	PlatformAPI implementation.PlatformAPI
}

func (inst *DriverBuilder) Create() afs.Driver {

	papi := inst.PlatformAPI
	if papi == nil {
		panic("libafs.DriverBuilder: PlatformAPI is nil")
	}

	ctx := new(implementation.Context)
	driver := new(innerCommonDriver)
	fs := new(innerCommonFS)
	fapi := new(innerCommonFullAPIImpl)
	fsio := new(innerCommonFSIO)

	driver.context = ctx
	fs.context = ctx
	fapi.context = ctx

	ctx.Driver = driver
	ctx.FS = fs
	ctx.PlatformAPI = papi
	ctx.FullAPI = fapi
	ctx.IO = fsio

	return ctx.Driver
}
