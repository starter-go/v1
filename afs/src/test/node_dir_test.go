package main

import (
	"testing"

	"github.com/starter-go/v1/afs"
	"github.com/starter-go/v1/afs/libafs"
)

// TestDirectory_Interface 测试 afs.Directory 接口是否被正确实现
func TestDirectory_Interface(t *testing.T) {
	// Create a filesystem for testing
	fs := libafs.Default()
	afs.Init(libafs.DefaultDriver())

	// Get root directory node
	rootNode := fs.GetNode("/")
	if rootNode == nil {
		t.Fatal("Failed to get root node")
	}

	// Check that rootNode implements Directory interface
	dir, ok := rootNode.(afs.Directory)
	if !ok {
		t.Fatal("Root node does not implement Directory interface")
	}

	// Verify that the Directory interface is properly implemented
	// by checking that all methods exist (this will compile-time error if not)
	_ = dir.ListNames()
	_ = dir.ListPaths()
	_ = dir.ListNodes()
	_ = dir.GetChild("")
	_ = dir.GetHref("")
}

// TestDirectory_ListNames 测试 Directory.ListNames 方法
func TestDirectory_ListNames(t *testing.T) {
	// Create a filesystem for testing
	fs := libafs.Default()
	afs.Init(libafs.DefaultDriver())

	// Get root directory node
	rootNode := fs.GetNode("/")
	if rootNode == nil {
		t.Fatal("Failed to get root node")
	}

	// Cast to Directory interface
	dir, ok := rootNode.(afs.Directory)
	if !ok {
		t.Fatal("Root node does not implement Directory interface")
	}

	// Test ListNames method
	names := dir.ListNames()

	// Basic validation - should not panic
	if names == nil {
		t.Error("ListNames should not return nil")
	}
}

// TestDirectory_ListPaths 测试 Directory.ListPaths 方法
func TestDirectory_ListPaths(t *testing.T) {
	// Create a filesystem for testing
	fs := libafs.Default()
	afs.Init(libafs.DefaultDriver())

	// Get root directory node
	rootNode := fs.GetNode("/")
	if rootNode == nil {
		t.Fatal("Failed to get root node")
	}

	// Cast to Directory interface
	dir, ok := rootNode.(afs.Directory)
	if !ok {
		t.Fatal("Root node does not implement Directory interface")
	}

	// Test ListPaths method
	paths := dir.ListPaths()

	// Basic validation - should not panic
	if paths == nil {
		t.Error("ListPaths should not return nil")
	}
}

// TestDirectory_ListNodes 测试 Directory.ListNodes 方法
func TestDirectory_ListNodes(t *testing.T) {
	// Create a filesystem for testing
	fs := libafs.Default()
	afs.Init(libafs.DefaultDriver())

	// Get root directory node
	rootNode := fs.GetNode("/")
	if rootNode == nil {
		t.Fatal("Failed to get root node")
	}

	// Cast to Directory interface
	dir, ok := rootNode.(afs.Directory)
	if !ok {
		t.Fatal("Root node does not implement Directory interface")
	}

	// Test ListNodes method
	nodes := dir.ListNodes()

	// Basic validation - should not panic
	if nodes == nil {
		t.Error("ListNodes should not return nil")
	}
}

// TestDirectory_GetChild 测试 Directory.GetChild 方法
func TestDirectory_GetChild(t *testing.T) {
	// Create a filesystem for testing
	fs := libafs.Default()
	afs.Init(libafs.DefaultDriver())

	// Get root directory node
	rootNode := fs.GetNode("/")
	if rootNode == nil {
		t.Fatal("Failed to get root node")
	}

	// Cast to Directory interface
	dir, ok := rootNode.(afs.Directory)
	if !ok {
		t.Fatal("Root node does not implement Directory interface")
	}

	// Test GetChild method with a common directory name
	// Note: This test may not find an actual child depending on the system
	_ = dir.GetChild("etc") // Common directory on Unix systems

	// Test with another common directory
	_ = dir.GetChild("usr")

	// Test with non-existent directory
	_ = dir.GetChild("non-existent-directory")
}

// TestDirectory_GetHref 测试 Directory.GetHref 方法
func TestDirectory_GetHref(t *testing.T) {
	// Create a filesystem for testing
	fs := libafs.Default()
	afs.Init(libafs.DefaultDriver())

	// Get root directory node
	rootNode := fs.GetNode("/")
	if rootNode == nil {
		t.Fatal("Failed to get root node")
	}

	// Cast to Directory interface
	dir, ok := rootNode.(afs.Directory)
	if !ok {
		t.Fatal("Root node does not implement Directory interface")
	}

	// Test GetHref method with relative path
	_ = dir.GetHref("etc")

	// Test with another relative path
	_ = dir.GetHref("usr")

	// Test with absolute path
	_ = dir.GetHref("/")

	// Test with non-existent path
	_ = dir.GetHref("non-existent-path")
}

// TestDirectory_NodeMethods 测试 Directory 继承自 Node 接口的方法
func TestDirectory_NodeMethods(t *testing.T) {
	// Create a filesystem for testing
	fs := libafs.Default()
	afs.Init(libafs.DefaultDriver())

	// Get root directory node
	rootNode := fs.GetNode("/")
	if rootNode == nil {
		t.Fatal("Failed to get root node")
	}

	// Cast to Directory interface
	dir, ok := rootNode.(afs.Directory)
	if !ok {
		t.Fatal("Root node does not implement Directory interface")
	}

	// Test Node interface methods
	_ = dir.String()
	_ = dir.GetName()
	_ = dir.GetPath()
	_ = dir.GetURI()
	_ = dir.IsDir()
	_ = dir.IsFile()
	_ = dir.IsLink()
	_ = dir.Exists()

	// Test GetParent method
	_ = dir.GetParent()

	// Test GetFileSystem method
	_ = dir.GetFileSystem()

	// Test GetMeta method
	_ = dir.GetMeta(nil)

	// Test GetIO method
	_ = dir.GetIO()
}
