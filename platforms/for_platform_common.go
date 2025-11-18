package platforms

import (
	"runtime"
)

func (inst *innerCurrentPlatformInfoLoader) onLoadCommonInfo(ib *InfoBuilder) {

	arch := runtime.GOARCH
	os := runtime.GOOS

	ib.Arch = ArchName(arch)
	ib.OST = OperatingSystemType(os)

	ib.OSV = ""
	ib.OSR = 0

}
