package implementation

import "github.com/starter-go/v1/afs"

type Context struct {
	FS     afs.FS
	Driver afs.Driver
	IO     afs.FileSystemIO

	FullAPI     FileSystemAPI
	PlatformAPI PlatformAPI
}
