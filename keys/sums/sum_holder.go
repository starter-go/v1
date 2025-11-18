package sums

import "github.com/starter-go/v1/keys"

type Sum struct {
	Algorithm keys.Algorithm
	Size      keys.SizeInBits
	Driver    Driver
	Hash      []byte
}
