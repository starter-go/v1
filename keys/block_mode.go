package keys

type BlockMode string

const (
	BlockModeFoo BlockMode = "foo"
)

type BlockModeInfo struct {
	Mode BlockMode

	Size SizeInBytes // the block-size
}
