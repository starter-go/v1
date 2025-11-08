package threads

import "sync"

var theGlobalSafeStrategy Strategy

func getSafeStrategy() Strategy {
	str := theGlobalSafeStrategy
	if str == nil {
		str = new(innerSafeStrategy)
		theGlobalSafeStrategy = str
	}
	return str
}

////////////////////////////////////////////////////////////////////////////////

type innerSafeStrategy struct{}

func (inst *innerSafeStrategy) Mode() Mode {
	return Safe
}

func (inst *innerSafeStrategy) NewLocker() Locker {
	return new(innerSafeLocker)
}

////////////////////////////////////////////////////////////////////////////////

type innerSafeLocker struct {
	mtx sync.Mutex
}

func (inst *innerSafeLocker) Lock() {
	inst.mtx.Lock()
}

func (inst *innerSafeLocker) Unlock() {
	inst.mtx.Unlock()
}

////////////////////////////////////////////////////////////////////////////////
