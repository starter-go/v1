package units

import "testing"

// 表示一个具体的测试环境
type Environment struct {
	t *testing.T
}

func (inst *Environment) Init(t *testing.T) {
	inst.t = t
}
