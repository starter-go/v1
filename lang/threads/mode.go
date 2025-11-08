package threads

// 表示并发的模式
type Mode int

const (
	Fast Mode = 0 // 快速模式 (无锁)
	Safe Mode = 1 // 安全模式 (加锁🔓)
)
