package platforms

import (
	"strconv"
	"strings"
)

type innerDefaultRevisionLoader struct {
}

// Load implements InfoLoader.
func (inst *innerDefaultRevisionLoader) Load() Info {
	return innerLoadWithLoader(inst)
}

// OnLoad implements InfoLoader.
func (inst *innerDefaultRevisionLoader) OnLoad(ib *InfoBuilder) error {

	rev := ib.OSR
	if rev > 0 {
		return nil
	}

	// 默认情况下， 通过使用 version 计算得到 rev
	ib.OSR = inst.computeRev(ib.OSV)
	return nil
}

func (inst *innerDefaultRevisionLoader) computeRev(ver OperatingSystemVersion) OperatingSystemRevision {

	const (
		sep1     = '.'
		sep2     = '-'
		sep3     = ' '
		sep      = '\n'
		strEmpty = ""
	)

	str := ver.String()
	str = strings.ReplaceAll(str, string(sep1), string(sep))
	str = strings.ReplaceAll(str, string(sep2), string(sep))
	str = strings.ReplaceAll(str, string(sep3), string(sep))
	rows := strings.Split(str, string(sep))
	builder := new(strings.Builder)

	for _, part := range rows {
		part = strings.TrimSpace(part)
		if part == strEmpty {
			continue
		}
		if builder.Len() > 0 {
			part = inst.addPadding0(part, 3)
		}
		builder.WriteString(part)
	}

	str = builder.String()
	n, _ := strconv.ParseUint(str, 10, 64)
	return OperatingSystemRevision(n)
}

func (inst *innerDefaultRevisionLoader) addPadding0(str string, width int) string {
	for len(str) < width {
		str = "0" + str
	}
	return str
}

func (inst *innerDefaultRevisionLoader) init() InfoLoader {
	return inst
}
