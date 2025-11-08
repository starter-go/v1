package threads

var theGlobalFastStrategy Strategy

func getFastStrategy() Strategy {
	st := theGlobalFastStrategy
	if st == nil {
		st = new(innerFastStrategy)
		theGlobalFastStrategy = st
	}
	return st
}

////////////////////////////////////////////////////////////////////////////////

type innerFastStrategy struct {
}

func (inst *innerFastStrategy) Mode() Mode {
	return Fast
}

func (inst *innerFastStrategy) NewLocker() Locker {
	return inst
}

func (inst *innerFastStrategy) Lock() {

}

func (inst *innerFastStrategy) Unlock() {

}
