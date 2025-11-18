package properties

import "github.com/starter-go/v1/collections/common"

////////////////////////////////////////////////////////////////////////////////

func Parse(text string, dst Table) (Table, error) {
	parser := Parser{}
	return parser.Parse(text, dst)
}

func Decode(data []byte, dst Table) (Table, error) {
	parser := Parser{}
	return parser.Decode(data, dst)
}

////////////////////////////////////////////////////////////////////////////////

type Parser struct {
	codec common.PropertyTableCodec
}

func (inst *Parser) Parse(text string, dst Table) (Table, error) {
	bin := []byte(text)
	return inst.Decode(bin, dst)
}

func (inst *Parser) Decode(bin []byte, dst Table) (Table, error) {
	tmp := make(map[string]string)
	tmp, err := inst.codec.Decode(bin, tmp)
	if err != nil {
		return nil, err
	}
	if dst == nil {
		dst = NewTable()
	}
	dst.Import(tmp)
	return dst, nil
}

////////////////////////////////////////////////////////////////////////////////
// EOF
