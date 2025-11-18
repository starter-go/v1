package keys

type Service interface {
	GetDriverManager() DriverManager

	GetDriverRegistry() DriverRegistry
}

type ServiceLoader interface {
	Load() Service
}

////////////////////////////////////////////////////////////////////////////////

var theServiceHolder innerServiceHolder

// 获取默认的 keys.Service
func GetService() Service {
	holder := &theServiceHolder
	return holder.getSer()
}

func Init(sl ServiceLoader) {

	if sl == nil {
		return
	}

	holder := &theServiceHolder
	holder.loader = sl
}

////////////////////////////////////////////////////////////////////////////////

type innerServiceHolder struct {
	service Service
	loader  ServiceLoader
}

func (inst *innerServiceHolder) getSer() Service {
	ser := inst.service
	if ser == nil {
		ser = inst.loadSer()
		inst.service = ser
	}
	return ser
}

func (inst *innerServiceHolder) loadSer() Service {
	ldr := inst.loader
	if ldr == nil {
		panic("no keys.ServiceLoader, use 'keys.Init(...)' to setup")
	}
	return ldr.Load()
}
