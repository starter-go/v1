package platforms

type innerCurrentPlatformInfoLoader struct {
}

func (inst *innerCurrentPlatformInfoLoader) load() Info {
	builder := new(InfoBuilder)
	inst.onLoadCommonInfo(builder)
	inst.onLoadSpecialInfo(builder)
	return builder.Info()
}

// func (inst *innerCurrentPlatformInfoLoader) onLoadCommonInfo(ib *InfoBuilder) {

// }

// func (inst *innerCurrentPlatformInfoLoader) onLoadSpecialInfo(ib *InfoBuilder) {

// }
