package properties

import (
	"github.com/starter-go/v1/collections"
	"github.com/starter-go/v1/collections/common"
	"github.com/starter-go/v1/lang/threads"
)

type Table interface {
	SetProperty(name string, value string)

	GetProperty(name string) string

	GetKeys(needSort bool) []string

	Keys() []string

	Getter() collections.Getter

	Setter() collections.Setter

	Export(dst map[string]string) map[string]string

	Import(src map[string]string)

	Reset()
}

func NewTable() Table {
	return NewTableWithMode(threads.Fast)
}

func NewTableWithStrategy(strategy threads.Strategy) Table {
	t := new(innerPropertyTable)
	t.init(strategy)
	return t
}

func NewTableWithMode(mode threads.Mode) Table {
	strategy := threads.GetStrategy(mode)
	t := new(innerPropertyTable)
	t.init(strategy)
	return t
}

////////////////////////////////////////////////////////////////////////////////

type innerPropertyTable struct {
	st common.StringTable
}

// GetKeys implements Table.
func (inst *innerPropertyTable) GetKeys(needSort bool) []string {
	return inst.st.Keys(needSort)
}

// Keys implements Table.
func (inst *innerPropertyTable) Keys() []string {
	return inst.st.Keys(true)
}

func (inst *innerPropertyTable) init(strategy threads.Strategy) Table {
	inst.st.Init(strategy)
	return inst
}

// Getter implements Table.
func (inst *innerPropertyTable) Getter() collections.Getter {
	return inst.st.Getter()
}

// Setter implements Table.
func (inst *innerPropertyTable) Setter() collections.Setter {
	return inst.st.Setter()
}

// Export implements Table.
func (inst *innerPropertyTable) Export(dst map[string]string) map[string]string {
	return inst.st.Export(dst)
}

// GetProperty implements Table.
func (inst *innerPropertyTable) GetProperty(name string) string {
	return inst.st.Get(name)
}

// Import implements Table.
func (inst *innerPropertyTable) Import(src map[string]string) {
	inst.st.Import(src)
}

// Reset implements Table.
func (inst *innerPropertyTable) Reset() {
	inst.st.Reset()
}

// SetProperty implements Table.
func (inst *innerPropertyTable) SetProperty(name string, value string) {
	inst.st.Set(name, value)
}
