package common

import (
	"fmt"
	"sort"
	"strings"
)

////////////////////////////////////////////////////////////////////////////////

type PropertyTableCodec struct {
}

func (inst *PropertyTableCodec) Decode(src []byte, dst map[string]string) (map[string]string, error) {
	parser := new(innerStrMapParser)
	return parser.parseBytes(src, dst)
}

func (inst *PropertyTableCodec) Encode(src map[string]string) ([]byte, error) {
	formatter := new(innerStrMapFormatter)
	str := formatter.format(src)
	return []byte(str), nil
}

////////////////////////////////////////////////////////////////////////////////

type innerStrMapFormatter struct {

	// config
	enableSegments      bool
	enableNormalizeKeys bool
	enableMakeLowerKeys bool
	enableSortByKey     bool

	// tmp
	currentKeyPrefix string
	builder          strings.Builder
}

func (inst *innerStrMapFormatter) init() {
	inst.builder.Reset()
	inst.currentKeyPrefix = ""
}

func (inst *innerStrMapFormatter) normalizeKey(key string) string {
	key = innerNormalizeStrMapKey(key)
	if inst.enableMakeLowerKeys {
		key = strings.ToLower(key)
	}
	return key
}

func (inst *innerStrMapFormatter) format(src map[string]string) string {

	tmp := make(map[string]string)
	keys := make([]string, 0)

	for k, v := range src {
		k = inst.normalizeKey(k)
		tmp[k] = v
	}

	for k := range tmp {
		keys = append(keys, k)
	}

	sort.Strings(keys)

	for _, key := range keys {
		value := tmp[key]
		inst.appendKeyValue(key, value)
	}

	return inst.builder.String()
}

func (inst *innerStrMapFormatter) appendKeyValue(key, value string) {

	kp1, kp2, kp3 := inst.splitKey(key)

	inst.handleKeyParts(kp1, kp2, kp3)

	inst.builder.WriteRune('\t')
	inst.builder.WriteString(kp3)
	inst.builder.WriteRune('=')
	inst.builder.WriteString(value)
	inst.builder.WriteRune('\n')
}

func (inst *innerStrMapFormatter) handleKeyParts(p1, p2, p3 string) {

	prefix := p1 + "." + p2
	current := inst.currentKeyPrefix

	if prefix == current {
		return // skip
	}
	inst.currentKeyPrefix = prefix

	// make a new segment head
	b := strings.Builder{}
	b.WriteRune('[')
	b.WriteString(p1)
	if p2 != "" {
		b.WriteString(" \"")
		b.WriteString(p2)
		b.WriteRune('"')
	}
	b.WriteString("]\n")
	inst.builder.WriteString(b.String())
}

func (inst *innerStrMapFormatter) splitKey(key string) (p1, p2, p3 string) {

	// p1 = type
	// p2 = name
	// p3 = field

	all := strings.Split(key, ".")
	count := len(all)
	const space = ""

	switch count {
	case 0:
		return space, space, space
	case 1:
		return space, space, all[0]
	case 2:
		return all[0], space, all[1]
	case 3:
		return all[0], all[1], all[2]
	}

	b := strings.Builder{}
	for i := 2; i < count; i++ {
		part := all[i]
		if b.Len() > 0 {
			b.WriteRune('.')
		}
		b.WriteString(part)
	}
	tail := b.String()
	return all[0], all[1], tail
}

////////////////////////////////////////////////////////////////////////////////

type innerStrMapParser struct {

	// config
	enableMakeLowerKeys         bool
	enableErrorForDuplicateKeys bool

	// tmp
	segmentKeyPrefix string
	output           map[string]string
}

func (inst *innerStrMapParser) init(dst map[string]string) {
	if dst == nil {
		dst = make(map[string]string)
	}
	inst.output = dst
}

func (inst *innerStrMapParser) parseBytes(src []byte, dst map[string]string) (map[string]string, error) {
	str := ""
	if src != nil {
		str = string(src)
	}
	return inst.parseString(str, dst)
}

func (inst *innerStrMapParser) parseString(src string, dst map[string]string) (map[string]string, error) {

	const (
		sep1 = "\r"
		sep2 = "\n"
		sep  = sep2
	)

	str := strings.ReplaceAll(src, sep1, sep)
	rows := strings.Split(str, sep)
	inst.init(dst)

	for idx, row := range rows {
		err := inst.handleRow(idx, row)
		if err != nil {
			return nil, err
		}
	}

	return inst.output, nil
}

func (inst *innerStrMapParser) handleRow(index int, row string) error {

	row = strings.TrimSpace(row)

	if row == "" {
		return nil //skip
	} else if strings.HasPrefix(row, "#") {
		return nil // skip
	}

	if strings.HasPrefix(row, "[") && strings.HasSuffix(row, "]") {
		return inst.handleSegmentHead(index, row)
	}

	return inst.handleKeyValue(index, row)
}

func (inst *innerStrMapParser) handleSegmentHead(index int, row string) error {
	const sep = "."
	str := row
	str = strings.ReplaceAll(str, "[", sep)
	str = strings.ReplaceAll(str, "]", sep)
	str = strings.ReplaceAll(str, "'", sep)
	str = strings.ReplaceAll(str, "\"", sep)
	inst.segmentKeyPrefix = str
	return nil
}

func (inst *innerStrMapParser) normalizeKey(key string) string {
	key = innerNormalizeStrMapKey(key)
	if inst.enableMakeLowerKeys {
		key = strings.ToLower(key)
	}
	return key
}

func (inst *innerStrMapParser) handleKeyValue(index int, row string) error {

	i := strings.IndexByte(row, '=')
	if i < 0 {
		return fmt.Errorf("bad key-value row[%d] (like 'key=value') : [%s]", index, row)
	}

	key1 := inst.segmentKeyPrefix
	key2 := row[0:i]
	value := strings.TrimSpace(row[i+1:])
	key := inst.normalizeKey(key1 + key2)

	older := inst.output[key]
	inst.output[key] = value

	if older != "" {
		err := fmt.Errorf("")
		if inst.enableErrorForDuplicateKeys {
			return err
		} else {
			fmt.Println("Warning:", err)
		}
	}

	return nil
}

////////////////////////////////////////////////////////////////////////////////

func innerNormalizeStrMapKey(rawkey string) string {

	const (
		sep1 = "."
		sep2 = "\n"
		sep  = sep2
	)

	key := strings.ReplaceAll(rawkey, sep1, sep)
	list := strings.Split(key, sep)
	b := new(strings.Builder)

	for _, el := range list {
		el = strings.TrimSpace(el)
		if el == "" {
			continue
		}
		if b.Len() > 0 {
			b.WriteRune('.')
		}
		b.WriteString(el)
	}

	return b.String()
}

////////////////////////////////////////////////////////////////////////////////
