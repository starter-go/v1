package sums

import "io"

type Builder interface {
	io.Closer

	io.Writer

	io.StringWriter

	Sum() *Sum
}
