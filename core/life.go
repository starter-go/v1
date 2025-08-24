package core

type OnLifePhaseFunc func() error

type OnCreateFunc OnLifePhaseFunc
type OnStartFunc OnLifePhaseFunc
type OnResumeFunc OnLifePhaseFunc
type OnLoopFunc OnLifePhaseFunc
type OnPauseFunc OnLifePhaseFunc
type OnStopFunc OnLifePhaseFunc
type OnDestroyFunc OnLifePhaseFunc

type OnStartPreFunc OnLifePhaseFunc
type OnStartPostFunc OnLifePhaseFunc

type OnStopPreFunc OnLifePhaseFunc
type OnStopPostFunc OnLifePhaseFunc

type Life struct {
	Name    string
	Order   int
	Enabled bool

	OnCreate  OnCreateFunc
	OnStart   OnStartFunc
	OnResume  OnResumeFunc
	OnLoop    OnLoopFunc
	OnPause   OnPauseFunc
	OnStop    OnStopFunc
	OnDestroy OnDestroyFunc

	OnStartPre  OnStartPreFunc
	OnStartPost OnStartPostFunc
	OnStopPre   OnStopPreFunc
	OnStopPost  OnStopPostFunc
}

type Lifecycle interface {
	Life() *Life
}
