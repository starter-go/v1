package sums

import "github.com/starter-go/v1/keys"

type Driver interface {
	keys.Driver

	Size() keys.SizeInBits

	NewBuilder() Builder

	Sum(data []byte) *Sum
}
