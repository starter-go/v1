package platforms

import (
	"fmt"
	"os"
)

type innerCurrentPlatformInfoLoader struct {
	moreLoaders []InfoLoader
}

// Load implements InfoLoader.
func (inst *innerCurrentPlatformInfoLoader) Load() Info {
	return innerLoadWithLoader(inst)
}

// OnLoad implements InfoLoader.
func (inst *innerCurrentPlatformInfoLoader) OnLoad(ib *InfoBuilder) error {

	// inst.onLoadCommonInfo(ib)
	// inst.onLoadSpecialInfo(ib)

	inst.tryInit()
	more := inst.moreLoaders

	for _, loader := range more {
		err := loader.OnLoad(ib)
		if err != nil {
			out := os.Stderr
			fmt.Fprintf(out, "[WARN] (innerCurrentPlatformInfoLoader.OnLoad) %s \n", err.Error())
		}
	}

	return nil
}

func (inst *innerCurrentPlatformInfoLoader) tryInit() InfoLoader {
	list := inst.moreLoaders
	if list == nil {
		list = inst.registerSpecialLoaders(list)
		list = inst.registerCommonLoaders(list)
		list = inst.registerDefaultRevsionLoader(list)
		inst.moreLoaders = list
	}
	return inst
}

// 注册特殊的加载器 (由各个平台的代码分别实现)
// func (inst *innerCurrentPlatformInfoLoader) registerSpecialLoaders(list []InfoLoader) []InfoLoader {
// return list
// }

// 注册通用的加载器
func (inst *innerCurrentPlatformInfoLoader) registerCommonLoaders(list []InfoLoader) []InfoLoader {
	list = append(list, new(innerCommonPlatformInfoLoader))
	return list
}

func (inst *innerCurrentPlatformInfoLoader) registerDefaultRevsionLoader(list []InfoLoader) []InfoLoader {

	// tmp for test
	// unameLoader := new(innerUnameInfoLoader)
	// unameLoader.init()
	// list = append(list, unameLoader)

	list = append(list, new(innerDefaultRevisionLoader))

	return list
}
