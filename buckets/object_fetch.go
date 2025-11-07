package buckets

import "context"

type FetchContext struct {

	// input

	Context context.Context

	WantMeta bool
	WantData bool

	// output

	Error error

	Meta *ObjectMeta
	Data *ObjectData
}

func (inst *FetchContext) Close() error {

	if inst == nil {
		return nil
	}

	data := inst.Data
	if data == nil {
		return nil
	}

	cl := data.Closer
	data.Closer = nil
	if cl == nil {
		return nil
	}

	return cl.Close()
}
