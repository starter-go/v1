package libafs

import (
	"github.com/starter-go/v1/afs"
)

var theDefaultDriver afs.Driver

////////////////////////////////////////////////////////////////////////////////

// Default 函数是 DefaultFS() 简单别名
func Default() afs.FS {
	return DefaultFS()
}

func DefaultFS() afs.FS {
	driver := DefaultDriver()
	return driver.GetFS()
}

func DefaultDriver() afs.Driver {
	driver := theDefaultDriver
	if driver == nil {
		driver = innerLoadDefaultDriver()
		theDefaultDriver = driver
	}
	return driver
}

////////////////////////////////////////////////////////////////////////////////

func innerLoadDefaultDriver() afs.Driver {
	builder := new(DriverBuilder)
	builder.PlatformAPI = createNewPlatformAPI()
	return builder.Create()
}

// func createNewPlatformAPI() implementation.PlatformAPI {}

////////////////////////////////////////////////////////////////////////////////
