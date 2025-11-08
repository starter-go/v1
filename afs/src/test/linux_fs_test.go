package main

import (
	"fmt"
	"testing"

	"github.com/starter-go/v1/afs"
	"github.com/starter-go/v1/afs/libafs"
)

func TestMain(t *testing.T) {
	// Create a new Linux file system

	fs := libafs.Default()
	drv := libafs.DefaultDriver()
	afs.Init(drv)

	// Test getting the root node
	rootNode := fs.GetNode("/")
	fmt.Printf("Root node exists: %v\n", rootNode.Exists())
	fmt.Printf("Root node is directory: %v\n", rootNode.IsDir())
	fmt.Printf("Root node path: %s\n", rootNode.Path())

	// Test listing roots
	roots := fs.ListRoots()
	fmt.Printf("Number of roots: %d\n", len(roots))

	// Test getting a node with URI
	uriNode := fs.GetNodeWithURI("file:///")
	fmt.Printf("URI node exists: %v\n", uriNode.Exists())

	// Test file system IO
	io := fs.IO()
	fmt.Printf("File system IO: %v\n", io)

	// Test node IO
	nodeIO := rootNode.IO()
	fmt.Printf("Node IO: %v\n", nodeIO)
}
