package lang

import (
	"fmt"
	"strconv"
)

type Module interface {
	fmt.Stringer

	Name() string

	Version() string

	Revision() int
}

func NewModule(name string, ver string, rev int) Module {
	m := new(innerModule)
	m.name = name
	m.ver = ver
	m.rev = rev
	return m
}

////////////////////////////////////////////////////////////////////////////////

type innerModule struct {
	name string
	ver  string
	rev  int
}

// Name implements Module.
func (i *innerModule) Name() string {
	return i.name
}

// Revision implements Module.
func (i *innerModule) Revision() int {
	return i.rev
}

// Version implements Module.
func (i *innerModule) Version() string {
	return i.ver
}

// String implements Module.
func (i *innerModule) String() string {
	rev := strconv.Itoa(i.rev)
	return ("module:" + i.name + "@" + i.ver + "-r" + rev)
}
