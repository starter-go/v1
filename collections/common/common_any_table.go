package common

import (
	"sort"

	"github.com/starter-go/v1/collections"
	"github.com/starter-go/v1/lang/threads"
)

type AnyTable struct {
	tag      string
	table    map[string]any
	locker   threads.Locker
	strategy threads.Strategy
}

func (inst *AnyTable) innerGetTable() map[string]any {
	tab := inst.table
	if tab == nil {
		tab = make(map[string]any)
		inst.table = tab
	}
	return tab
}

func (inst *AnyTable) Init(tag string, st threads.Strategy) {
	if st == nil {
		st = threads.GetStrategy(threads.Safe)
	}
	inst.tag = tag
	inst.locker = st.NewLocker()
	inst.strategy = st
	inst.innerGetTable()
}

func (inst *AnyTable) Getter() collections.Getter {}

func (inst *AnyTable) Setter() collections.Setter {}

func (inst *AnyTable) Keys() []string {

	l := inst.locker
	l.Lock()
	defer l.Unlock()

	dst := make([]string, 0)
	src := inst.innerGetTable()
	for k, v := range src {
		if v == nil {
			continue
		}
		dst = append(dst, k)
	}
	sort.Strings(dst)
	return dst
}

func (inst *AnyTable) Export(dst map[string]any) map[string]any {

	l := inst.locker
	l.Lock()
	defer l.Unlock()

	if dst == nil {
		dst = make(map[string]any)
	}
	src := inst.innerGetTable()
	for k, v := range src {
		if v == nil {
			continue
		}
		dst[k] = v
	}
	return dst
}

func (inst *AnyTable) Import(src map[string]any) {

	l := inst.locker
	l.Lock()
	defer l.Unlock()

	dst := inst.innerGetTable()
	if src == nil {
		return
	}
	for k, v := range src {
		if v == nil {
			continue
		}
		dst[k] = v
	}
}

func (inst *AnyTable) SetValue(name string, value any) {

	l := inst.locker
	l.Lock()
	defer l.Unlock()

	tab := inst.innerGetTable()
	tab[name] = value
}

func (inst *AnyTable) GetValue(name string) any {

	l := inst.locker
	l.Lock()
	defer l.Unlock()

	tab := inst.innerGetTable()
	return tab[name]
}
