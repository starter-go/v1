package u4go

import (
	"testing"

	"github.com/starter-go/v1/lang/threads"
)

func TestGetStrategyWithFastMode(t *testing.T) {
	strategy := threads.GetStrategy(threads.Fast)
	if strategy == nil {
		t.Error("Expected strategy to be not nil")
	}
}

func TestGetStrategyWithSafeMode(t *testing.T) {
	strategy := threads.GetStrategy(threads.Safe)
	if strategy == nil {
		t.Error("Expected strategy to be not nil")
	}
}

func TestFastStrategyMode(t *testing.T) {
	strategy := threads.GetStrategy(threads.Fast)
	mode := strategy.Mode()
	if mode != threads.Fast {
		t.Errorf("Expected mode to be Fast, got %v", mode)
	}
}

func TestSafeStrategyMode(t *testing.T) {
	strategy := threads.GetStrategy(threads.Safe)
	mode := strategy.Mode()
	if mode != threads.Safe {
		t.Errorf("Expected mode to be Safe, got %v", mode)
	}
}

func TestFastStrategyNewLocker(t *testing.T) {
	strategy := threads.GetStrategy(threads.Fast)
	locker := strategy.NewLocker()
	if locker == nil {
		t.Error("Expected locker to be not nil")
	}
}

func TestSafeStrategyNewLocker(t *testing.T) {
	strategy := threads.GetStrategy(threads.Safe)
	locker := strategy.NewLocker()
	if locker == nil {
		t.Error("Expected locker to be not nil")
	}
}

func TestFastLockerMethods(t *testing.T) {
	strategy := threads.GetStrategy(threads.Fast)
	locker := strategy.NewLocker()

	// These should not panic
	locker.Lock()
	locker.Unlock()
}

func TestSafeLockerMethods(t *testing.T) {
	strategy := threads.GetStrategy(threads.Safe)
	locker := strategy.NewLocker()

	// These should not panic
	locker.Lock()
	locker.Unlock()
}
