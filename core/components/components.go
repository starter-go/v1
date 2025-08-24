package components

type ComponentRegistration struct {
	ID        string
	Classes   []string
	Aliases   []string
	Prototype any
	Scope     Scope
	Factory   func() any
	Injector  func(inst any) error
}
