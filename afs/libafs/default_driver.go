package libafs

import (
	"github.com/starter-go/v1/afs"
)

////////////////////////////////////////////////////////////////////////////////

// Default 函数是 DefaultFS() 简单别名
func Default() afs.FS {
	h := &theDefaultFSHolder
	return h.getFS()
}

func DefaultFS() afs.FS {
	h := &theDefaultFSHolder
	return h.getFS()
}

func DefaultDriver() afs.Driver {
	h := &theDefaultFSHolder
	return h.getDriver()
}

////////////////////////////////////////////////////////////////////////////////

var theDefaultFSHolder innerDefaultFSHolder

type innerDefaultFSHolder struct {
	fs     afs.FS
	driver afs.Driver
}

func (inst *innerDefaultFSHolder) loadDriver() afs.Driver {
	builder := new(DriverBuilder)
	builder.PlatformAPI = createNewPlatformAPI()
	return builder.Create()
}

func (inst *innerDefaultFSHolder) loadFS() afs.FS {
	dr := inst.getDriver()
	return dr.GetFS()
}

func (inst *innerDefaultFSHolder) getFS() afs.FS {
	fs := inst.fs
	if fs == nil {
		fs = inst.loadFS()
		inst.fs = fs
	}
	return fs
}

func (inst *innerDefaultFSHolder) getDriver() afs.Driver {
	dr := inst.driver
	if dr == nil {
		dr = inst.loadDriver()
		inst.driver = dr
	}
	return dr
}

////////////////////////////////////////////////////////////////////////////////
