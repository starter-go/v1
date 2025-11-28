package platforms

import (
	"fmt"
	"os/exec"
	"strings"
)

// 通过执行 'uname' 命令, 实现加载 OS 信息
type innerUnameInfoLoader struct {
	propertiesMapping map[string]int // map[name] index
}

// Load implements InfoLoader.
func (inst *innerUnameInfoLoader) Load() Info {
	return innerLoadWithLoader(inst)
}

// OnLoad implements InfoLoader.
func (inst *innerUnameInfoLoader) OnLoad(ib *InfoBuilder) error {

	out := new(strings.Builder)
	args := []string{"-a"}
	cmd := exec.Command("uname", args...)
	cmd.Stdout = out

	err := cmd.Run()
	if err != nil {
		return err
	}

	inst.handleOutputString(out.String(), ib)
	return nil
}

func (inst *innerUnameInfoLoader) init() {
	inst.addMapping("uname.kernel-name", 0)
	inst.addMapping("uname.nodename", 1)
	inst.addMapping("uname.kernel-release", 2)
	inst.addMapping("uname.kernel-version", 3)
	inst.addMapping("uname.machine", 4)
}

func (inst *innerUnameInfoLoader) addMapping(name string, index int) {
	table := inst.propertiesMapping
	if table == nil {
		table = make(map[string]int)
		inst.propertiesMapping = table
	}
	table[name] = index
}

func (inst *innerUnameInfoLoader) handleOutputString(str string, ib *InfoBuilder) {

	const (
		sep1 = ' '
		sep2 = '\n'
		sep  = string(sep2)
	)

	fmt.Println("uname output string:  " + str)

	str = strings.ReplaceAll(str, string(sep1), sep)
	parts := strings.Split(str, sep)

	fmt.Println()
	for i, p := range parts {
		fmt.Printf("\n  part[%d] = %s", i, p)
	}
	fmt.Println()

	mapping := inst.propertiesMapping
	for name, index := range mapping {
		value, ok := inst.getValueFromArray(parts, index)
		if !ok {
			continue
		}
		ib.SetProperty(name, value)
	}

}

func (inst *innerUnameInfoLoader) getValueFromArray(array []string, index int) (value string, ok bool) {
	count := len(array)
	if 0 <= index && index < count {
		return array[index], true
	}
	return "", false
}
