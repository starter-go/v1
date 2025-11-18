package platforms

import "github.com/starter-go/v1/lang"

const (
	theModuleName     = "github.com/starter-go/v1/platforms"
	theModuleVersion  = "v0.9.1"
	theModuleRevision = 1
)

func GetThisModule() lang.Module {
	return lang.NewModule(theModuleName, theModuleVersion, theModuleRevision)
}
