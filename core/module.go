package core

import (
	"embed"

	"github.com/starter-go/v1/core/components"
	"github.com/starter-go/v1/core/resources"
)

type Module interface {
	Name() string
	Version() string
	Revision() int
	Resources() []*resources.ResourceRegistration
	Components() []*components.ComponentRegistration
	Dependencies() []Module
}

////////////////////////////////////////////////////////////////////////////////

type ModuleBuilder struct {
	name    string
	version string
	rev     int
	rlist   []*resources.ResourceRegistration
	clist   []*components.ComponentRegistration
	deplist []Module
}

func (inst *ModuleBuilder) Build() Module {}

func (inst *ModuleBuilder) SetName(name string) *ModuleBuilder {
	inst.name = name
	return inst
}

func (inst *ModuleBuilder) SetVersion(ver string) *ModuleBuilder {
	inst.version = ver
	return inst
}

func (inst *ModuleBuilder) SetRevision(rev int) *ModuleBuilder {
	inst.rev = rev
	return inst
}

func (inst *ModuleBuilder) AddComponents(list ...*components.ComponentRegistration) *ModuleBuilder {
	if list != nil {
		inst.clist = append(inst.clist, list...)
	}
	return inst
}

func (inst *ModuleBuilder) AddResources(list ...*resources.ResourceRegistration) *ModuleBuilder {
	if list != nil {
		inst.rlist = append(inst.rlist, list...)
	}
	return inst
}

func (inst *ModuleBuilder) EmbedResources(resFS embed.FS, resPath string) *ModuleBuilder {}

func (inst *ModuleBuilder) Depend(dep Module) *ModuleBuilder {
	if dep != nil {
		inst.deplist = append(inst.deplist, dep)
	}
	return inst
}
