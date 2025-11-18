package keys

import "encoding/pem"

// PEM 表示一个保存了一组 blocks 的结构
type PEM struct {
	Blocks []*pem.Block
}
