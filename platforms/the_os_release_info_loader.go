package platforms

import (
	"io"
	"io/fs"
	"os"
	"strings"
)

// 通过读取 '/etc/os-release' 文件, 实现加载 OS 信息
type innerLinuxOSReleaseInfoLoader struct {
}

// Load implements InfoLoader.
func (inst *innerLinuxOSReleaseInfoLoader) Load() Info {
	ib := new(InfoBuilder)
	inst.OnLoad(ib)
	return ib.Info()
}

// OnLoad implements InfoLoader.
func (inst *innerLinuxOSReleaseInfoLoader) OnLoad(ib *InfoBuilder) error {

	path := "/etc/os-release"
	flag := os.O_RDONLY
	perm := fs.ModePerm

	file, err := os.OpenFile(path, flag, perm)
	if err != nil {
		return err
	}

	defer file.Close()

	data, err := io.ReadAll(file)
	if err != nil {
		return err
	}

	table := inst.parseRawData(data)
	inst.handleKeyValuePairs(table, ib)
	return nil
}

func (inst *innerLinuxOSReleaseInfoLoader) parseRawData(raw []byte) map[string]string {

	const (
		sep1 = '\r'
		sep2 = '\n'
		sep  = sep2
		eq   = '='
		n    = 2
	)

	table := make(map[string]string)
	str := string(raw)

	str = strings.ReplaceAll(str, string(sep1), string(sep2))
	rows := strings.Split(str, string(sep))

	for _, row := range rows {
		parts := strings.SplitN(row, string(eq), n)
		if len(parts) != n {
			continue
		}
		name := strings.TrimSpace(parts[0])
		value := strings.TrimSpace(parts[1])
		table[name] = value
	}

	return table
}

func (inst *innerLinuxOSReleaseInfoLoader) getValueByKey(key string, table map[string]string) string {
	const (
		mark1 = '"'
		mark  = string(mark1)
	)
	value := strings.TrimSpace(table[key])
	length := len(value)
	if strings.HasPrefix(value, mark) && strings.HasSuffix(value, mark) && (length >= 2) {
		value = value[1 : length-1]
	}
	return value
}

func (inst *innerLinuxOSReleaseInfoLoader) handleKeyValuePairs(table map[string]string, ib *InfoBuilder) {

	id := inst.getValueByKey("ID", table)
	vid := inst.getValueByKey("VERSION_ID", table)

	ib.props = table
	ib.OSR = 0
	ib.OSV = OperatingSystemVersion(vid)
	ib.OSN = OperatingSystemName(id)

}

func (inst *innerLinuxOSReleaseInfoLoader) init() InfoLoader {
	return inst
}
