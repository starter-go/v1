package buckets

import "context"

type PutContext struct {

	// input

	Context context.Context

	Meta *ObjectMeta
	Data *ObjectData

	// output
	Error error
}

func (inst *PutContext) Close() error {

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
