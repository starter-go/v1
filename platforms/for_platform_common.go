package platforms

import (
	"runtime"
)

type innerCommonPlatformInfoLoader struct {
}

// Load implements InfoLoader.
func (inst *innerCommonPlatformInfoLoader) Load() Info {
	ib := new(InfoBuilder)
	inst.OnLoad(ib)
	return ib.Info()
}

// OnLoad implements InfoLoader.
func (inst *innerCommonPlatformInfoLoader) OnLoad(ib *InfoBuilder) error {

	arch := runtime.GOARCH
	os := runtime.GOOS

	ib.Arch = ArchName(arch)
	ib.OST = OperatingSystemType(os)

	// ib.OSV = ""
	// ib.OSR = 0

	return nil
}

func (inst *innerCommonPlatformInfoLoader) init() InfoLoader {
	return inst
}
