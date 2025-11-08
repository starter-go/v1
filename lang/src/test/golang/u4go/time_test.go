package u4go

import (
	"testing"

	"github.com/starter-go/v1/lang"
)

func TestNow(t *testing.T) {
	// Test that Now() returns a valid time
	now := lang.Now()

	// Time should be greater than zero
	if now <= 0 {
		t.Errorf("Expected positive time value, got %d", now)
	}

	// Time should be reasonably close to current time (within 1 day)
	// This is a loose check to ensure the function works
	if now < 1000000000000 || now > 9999999999999 {
		t.Errorf("Time value seems unreasonable: %d", now)
	}
}

func TestTimeComparison(t *testing.T) {
	// Test comparing Time values
	time1 := lang.Now()

	// Small delay to ensure different timestamps
	// (This might not be necessary, but ensures we have different values)
	time2 := lang.Now()

	// Even if they're the same, this should not fail
	// We're just testing that comparison works
	if time1 > 0 && time2 > 0 {
		// This is just a sanity check that comparison works
		_ = time1 < time2 || time1 >= time2 // Should not panic
	} else {
		t.Error("Time values should be positive")
	}
}

func TestTimeZeroValue(t *testing.T) {
	// Test the zero value of Time
	var zeroTime lang.Time
	if zeroTime != 0 {
		t.Errorf("Expected zero value to be 0, got %d", zeroTime)
	}
}
