package afs

import "fmt"

var theDefaultFS FS
var theDefaultDriver Driver

func Init(driver Driver) {
	if driver != nil {
		theDefaultDriver = driver
	}
}

func Default() FS {
	fs := theDefaultFS
	if fs == nil {
		fs = innerLoadDefaultFS()
		theDefaultFS = fs
	}
	return fs
}

func innerLoadDefaultFS() FS {
	driver := theDefaultDriver
	if driver == nil {
		err := fmt.Errorf("AFS: no default driver, use afs.Init() to initial")
		panic(err)
	}
	return driver.GetFS()
}
