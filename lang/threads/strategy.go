package threads

// Strategy 提供一种处理并发的策略
type Strategy interface {

	// 获取并发模式
	Mode() Mode

	// 新建一个锁
	NewLocker() Locker
}

func GetStrategy(mode Mode) Strategy {
	if mode == Fast {
		return getFastStrategy()
	}
	return getSafeStrategy()
}
