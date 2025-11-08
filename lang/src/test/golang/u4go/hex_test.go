package u4go

import (
	"testing"

	"github.com/starter-go/v1/lang"
)

func TestHexFromString(t *testing.T) {
	// Test with a simple hex string
	hexStr := "abcdef123456"
	hex := lang.HexFromString(hexStr)

	// Since Hex is a type alias, we can compare it directly
	if string(hex) != hexStr {
		t.Errorf("Expected %s, got %s", hexStr, string(hex))
	}
}

func TestHexFromBytes(t *testing.T) {
	// Test with some bytes
	bytes := []byte{0xab, 0xcd, 0xef, 0x12, 0x34, 0x56}
	hex := lang.HexFromBytes(bytes)

	// Convert back to string to verify
	expected := "abcdef123456"
	if string(hex) != expected {
		t.Errorf("Expected %s, got %s", expected, string(hex))
	}
}
