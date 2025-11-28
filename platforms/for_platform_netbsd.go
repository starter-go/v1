package platforms

// 注册特殊的加载器 (由各个平台的代码分别实现)
func (inst *innerCurrentPlatformInfoLoader) registerSpecialLoaders(list []InfoLoader) []InfoLoader {

	unameLoader := new(innerUnameInfoLoader)
	unameLoader.init()

	list = append(list, unameLoader)
	list = append(list, new(innerNetBSDInfoLoader))

	return list
}

////////////////////////////////////////////////////////////////////////////////

type innerNetBSDInfoLoader struct{}

// Load implements InfoLoader.
func (inst *innerNetBSDInfoLoader) Load() Info {

	return innerLoadWithLoader(inst)
}

// OnLoad implements InfoLoader.
func (i *innerNetBSDInfoLoader) OnLoad(ib *InfoBuilder) error {

	osn := ib.GetProperty("uname.kernel-name")
	ver := ib.GetProperty("uname.kernel-release")

	ib.OSN = OperatingSystemName(osn)
	ib.OSV = OperatingSystemVersion(ver)

	return nil
}
