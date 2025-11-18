package arguments

import "strings"

func NewList(args []string) *List {
	list := new(List)
	list.Init(args)
	return list
}

////////////////////////////////////////////////////////////////////////////////

type List struct {
	all   []*Argument
	flags map[string]*Argument
}

func (inst *List) Init(args []string) {

	list := make([]*Argument, 0)
	table := make(map[string]*Argument)
	count := 0
	var ptr *Argument = nil

	for _, item1 := range args {
		item2 := inst.innerMakeArgItem(count, item1)
		if item2 == nil {
			continue
		}
		count++
		if item2.IsFlag() {
			table[item2.value] = item2
		}
		if ptr != nil {
			ptr.next = item2
		}
		list = append(list, item2)
		ptr = item2
	}

	inst.all = list
	inst.flags = table
}

func (inst *List) innerMakeArgItem(idx int, rawValue string) *Argument {

	value := strings.TrimSpace(rawValue)
	if len(value) == 0 {
		return nil
	}

	item := new(Argument)
	item.index = idx
	item.rawValue = rawValue
	item.value = value
	item.list = inst

	return item
}

func (inst *List) Reset() {
	all := inst.all
	for _, item := range all {
		item.used = false
	}
}

func (inst *List) Items() []*Argument {
	return inst.all
}

func (inst *List) GetFlag(key string) *Argument {
	return inst.flags[key]
}

func (inst *List) ListItemsWithFilter(accept func(*Argument) bool) []*Argument {
	src := inst.all
	dst := make([]*Argument, 0)
	for _, item := range src {
		if accept(item) {
			dst = append(dst, item)
		}
	}
	return dst
}

func (inst *List) String() string {

	builder := new(strings.Builder)
	src := inst.all

	for _, item := range src {
		if builder.Len() > 0 {
			builder.WriteRune(' ')
		}
		builder.WriteString(item.value)
	}

	return builder.String()
}

////////////////////////////////////////////////////////////////////////////////
// EOF
