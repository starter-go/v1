package platforms

// 注册特殊的加载器 (由各个平台的代码分别实现)
func (inst *innerCurrentPlatformInfoLoader) registerSpecialLoaders(list []InfoLoader) []InfoLoader {
	return list
}
