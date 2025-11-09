package afs

import "github.com/starter-go/v1/lang"

const (
	theModuleName     = "github.com/starter-go/v1/afs"
	theModuleVersion  = "v0.9.0"
	theModuleRevision = 1
)

func GetThisModule() lang.Module {
	return lang.NewModule(theModuleName, theModuleVersion, theModuleRevision)
}
