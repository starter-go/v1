package collections

import "github.com/starter-go/v1/lang"

const (
	theModuleName     = "github.com/starter-go/v1/collections"
	theModuleVersion  = "v0.0.0"
	theModuleRevision = 0
)

func GetThisModule() lang.Module {
	return lang.NewModule(theModuleName, theModuleVersion, theModuleRevision)
}
