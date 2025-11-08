package errors

type Holder interface {
	ErrorHandler

	HasError() bool

	First() error

	Last() error

	All() []error
}

////////////////////////////////////////////////////////////////////////////////

type ErrorList interface {
	Holder
}

func NewErrorList() ErrorList {
	return new(innerErrorListHolder)
}

type innerErrorListHolder struct {
	all []error
}

func (inst *innerErrorListHolder) _impl() ErrorList {
	return inst
}

func (inst *innerErrorListHolder) First() error {
	list := inst.all
	count := len(list)
	if count > 0 {
		return list[0]
	}
	return nil
}

func (inst *innerErrorListHolder) HandleError(err error) {
	if err == nil {
		return
	}
	inst.all = append(inst.all, err)
}

func (inst *innerErrorListHolder) HasError() bool {
	list := inst.all
	count := len(list)
	return (count > 0)
}

func (inst *innerErrorListHolder) Last() error {
	list := inst.all
	count := len(list)
	if count > 0 {
		return list[count-1]
	}
	return nil
}

func (inst *innerErrorListHolder) All() []error {
	return inst.all
}

////////////////////////////////////////////////////////////////////////////////

type SimpleErrorHolder interface {
	Holder
}

func NewErrorHolder() SimpleErrorHolder {
	return new(innerSimpleErrorHolder)
}

type innerSimpleErrorHolder struct {
	first error
	last  error
	count int
}

func (inst *innerSimpleErrorHolder) _impl() SimpleErrorHolder {
	return inst
}

func (inst *innerSimpleErrorHolder) First() error {
	return inst.first
}

func (inst *innerSimpleErrorHolder) HandleError(err error) {
	if err == nil {
		return
	}
	if inst.first == nil {
		inst.first = err
	}
	inst.last = err
	inst.count++
}

func (inst *innerSimpleErrorHolder) HasError() bool {
	return (inst.count > 0)
}

func (inst *innerSimpleErrorHolder) Last() error {
	return inst.last
}

func (inst *innerSimpleErrorHolder) All() []error {
	list := make([]error, 0)
	count := inst.count
	if count == 1 {
		list = append(list, inst.first)
	} else if count > 1 {
		list = append(list, inst.first)
		list = append(list, inst.last)
	}
	return list
}

////////////////////////////////////////////////////////////////////////////////
