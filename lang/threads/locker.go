package threads

type Locker interface {
	Lock()
	Unlock()
}
