package lang

const (
	theModuleName     = "github.com/starter-go/v1/lang"
	theModuleVersion  = "v0.0.0"
	theModuleRevision = 0
)

func GetThisModule() Module {
	return NewModule(theModuleName, theModuleVersion, theModuleRevision)
}
