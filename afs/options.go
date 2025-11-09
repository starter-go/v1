package afs

// Options 包含几个通用的选项,对于不同的操作,某些字段可能是无效的
type Options struct {

	// 文件模式
	Mode FileMode

	// 读写标志位 (value like: 'os.O_RDONLY')
	Flag int

	// -1 表示为空
	User UserID

	// -1 表示为空
	Group GroupID

	// 指示操作不要使用缓存
	Reload bool
}
