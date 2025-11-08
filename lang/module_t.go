package lang

const (
	theModuleName     = "github.com/starter-go/v1/lang"
	theModuleVersion  = "v0.9.1"
	theModuleRevision = 1
)

func GetThisModule() Module {
	return NewModule(theModuleName, theModuleVersion, theModuleRevision)
}
