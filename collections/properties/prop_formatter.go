package properties

import "github.com/starter-go/v1/collections/common"

////////////////////////////////////////////////////////////////////////////////

func Format(tab Table) (string, error) {
	f := Formatter{}
	return f.Format(tab)
}

func Encode(tab Table) ([]byte, error) {
	f := Formatter{}
	return f.Encode(tab)
}

////////////////////////////////////////////////////////////////////////////////

type Formatter struct {
	codec common.PropertyTableCodec
}

func (inst *Formatter) Encode(tab Table) ([]byte, error) {
	if tab == nil || inst == nil {
		return []byte{}, nil
	}
	tmp := tab.Export(nil)
	return inst.codec.Encode(tmp)
}

func (inst *Formatter) Format(tab Table) (string, error) {
	data, err := inst.Encode(tab)
	if err != nil {
		return "", err
	}
	return string(data), nil
}

////////////////////////////////////////////////////////////////////////////////
// EOF
