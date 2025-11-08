package libafs

import (
	"github.com/starter-go/v1/afs"
	"github.com/starter-go/v1/afs/implementation"
)

type innerCommonDriver struct {
	context *implementation.Context
}

func (inst *innerCommonDriver) _impl() afs.Driver {
	return inst
}

func (inst *innerCommonDriver) GetFS() afs.FS {
	return inst.context.FS
}

func (inst *innerCommonDriver) CreateNewFS() afs.FS {

	ctx1 := inst.context
	ctx2 := new(implementation.Context)
	fs2 := new(innerCommonFS)
	cio := new(innerCommonFSIO)

	fs2.context = ctx2

	ctx2.Driver = ctx1.Driver
	ctx2.FS = fs2
	ctx2.FullAPI = ctx1.FullAPI
	ctx2.IO = cio
	ctx2.PlatformAPI = ctx1.PlatformAPI

	return ctx2.FS
}
