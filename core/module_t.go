package core

import "embed"

const (
	theModuleName    = ""
	theModuleVersion = "v0.0.0"
	theModuleRev     = 0
)

////////////////////////////////////////////////////////////////////////////////

const (
	theSrcMainResPath = ""
	theSrcTestResPath = ""
)

var theSrcMainResFS embed.FS

var theSrcTestResFS embed.FS

////////////////////////////////////////////////////////////////////////////////

func NewModuleBuilder() any {
	return ""
}
