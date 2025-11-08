package main

import (
	"testing"

	"github.com/starter-go/v1/afs"
	"github.com/starter-go/v1/afs/libafs"
)

// TestFile_Interface 测试 afs.File 接口是否被正确实现
func TestFile_Interface(t *testing.T) {
	// Create a filesystem for testing
	fs := libafs.Default()
	afs.Init(libafs.DefaultDriver())

	// Get a file node (using a common system file)
	fileNode := fs.GetNode("/etc/passwd")

	// If the file doesn't exist, try another common file
	if !fileNode.Exists() {
		fileNode = fs.GetNode("/etc/hosts")
	}

	// If that doesn't exist either, try a third option
	if !fileNode.Exists() {
		fileNode = fs.GetNode("/etc/group")
	}

	// If none of the common files exist, we'll skip the interface test
	if !fileNode.Exists() {
		t.Skip("No common system files found for testing")
	}

	// Check that fileNode implements File interface
	file, ok := fileNode.(afs.File)
	if !ok {
		t.Fatal("Node does not implement File interface")
	}

	// Verify that the File interface is properly implemented
	// by checking that all methods exist (this will compile-time error if not)
	_ = file.GetSize(false)
	_ = file.GetNameSuffix()
	_ = file.GetNameSuffixLower()
	_ = file.GetNameSuffixUpper()
}

// TestFile_GetSize 测试 File.GetSize 方法
func TestFile_GetSize(t *testing.T) {
	// Create a filesystem for testing
	fs := libafs.Default()
	afs.Init(libafs.DefaultDriver())

	// Get a file node (using a common system file)
	fileNode := fs.GetNode("/etc/passwd")

	// If the file doesn't exist, try another common file
	if !fileNode.Exists() {
		fileNode = fs.GetNode("/etc/hosts")
	}

	// If that doesn't exist either, try a third option
	if !fileNode.Exists() {
		fileNode = fs.GetNode("/etc/group")
	}

	// If none of the common files exist, we'll skip the test
	if !fileNode.Exists() {
		t.Skip("No common system files found for testing")
	}

	// Cast to File interface
	file, ok := fileNode.(afs.File)
	if !ok {
		t.Fatal("Node does not implement File interface")
	}

	// Test GetSize method
	size := file.GetSize(false)

	// Basic validation
	if size < 0 {
		t.Error("GetSize should not return negative value")
	}
}

// TestFile_GetNameSuffix 测试 File.GetNameSuffix 方法
func TestFile_GetNameSuffix(t *testing.T) {
	// Create a filesystem for testing
	fs := libafs.Default()
	afs.Init(libafs.DefaultDriver())

	// Get a file node (using a common system file)
	fileNode := fs.GetNode("/etc/passwd")

	// If the file doesn't exist, try another common file
	if !fileNode.Exists() {
		fileNode = fs.GetNode("/etc/hosts")
	}

	// If that doesn't exist either, try a third option
	if !fileNode.Exists() {
		fileNode = fs.GetNode("/etc/group")
	}

	// If none of the common files exist, we'll skip the test
	if !fileNode.Exists() {
		t.Skip("No common system files found for testing")
	}

	// Cast to File interface
	file, ok := fileNode.(afs.File)
	if !ok {
		t.Fatal("Node does not implement File interface")
	}

	// Test GetNameSuffix method
	suffix := file.GetNameSuffix()

	// Basic validation - should not panic
	_ = suffix
}

// TestFile_GetNameSuffixLower 测试 File.GetNameSuffixLower 方法
func TestFile_GetNameSuffixLower(t *testing.T) {
	// Create a filesystem for testing
	fs := libafs.Default()
	afs.Init(libafs.DefaultDriver())

	// Get a file node (using a common system file)
	fileNode := fs.GetNode("/etc/passwd")

	// If the file doesn't exist, try another common file
	if !fileNode.Exists() {
		fileNode = fs.GetNode("/etc/hosts")
	}

	// If that doesn't exist either, try a third option
	if !fileNode.Exists() {
		fileNode = fs.GetNode("/etc/group")
	}

	// If none of the common files exist, we'll skip the test
	if !fileNode.Exists() {
		t.Skip("No common system files found for testing")
	}

	// Cast to File interface
	file, ok := fileNode.(afs.File)
	if !ok {
		t.Fatal("Node does not implement File interface")
	}

	// Test GetNameSuffixLower method
	suffix := file.GetNameSuffixLower()

	// Basic validation - should not panic
	_ = suffix
}

// TestFile_GetNameSuffixUpper 测试 File.GetNameSuffixUpper 方法
func TestFile_GetNameSuffixUpper(t *testing.T) {
	// Create a filesystem for testing
	fs := libafs.Default()
	afs.Init(libafs.DefaultDriver())

	// Get a file node (using a common system file)
	fileNode := fs.GetNode("/etc/passwd")

	// If the file doesn't exist, try another common file
	if !fileNode.Exists() {
		fileNode = fs.GetNode("/etc/hosts")
	}

	// If that doesn't exist either, try a third option
	if !fileNode.Exists() {
		fileNode = fs.GetNode("/etc/group")
	}

	// If none of the common files exist, we'll skip the test
	if !fileNode.Exists() {
		t.Skip("No common system files found for testing")
	}

	// Cast to File interface
	file, ok := fileNode.(afs.File)
	if !ok {
		t.Fatal("Node does not implement File interface")
	}

	// Test GetNameSuffixUpper method
	suffix := file.GetNameSuffixUpper()

	// Basic validation - should not panic
	_ = suffix
}

// TestFile_NodeMethods 测试 File 继承自 Node 接口的方法
func TestFile_NodeMethods(t *testing.T) {
	// Create a filesystem for testing
	fs := libafs.Default()
	afs.Init(libafs.DefaultDriver())

	// Get a file node (using a common system file)
	fileNode := fs.GetNode("/etc/passwd")

	// If the file doesn't exist, try another common file
	if !fileNode.Exists() {
		fileNode = fs.GetNode("/etc/hosts")
	}

	// If that doesn't exist either, try a third option
	if !fileNode.Exists() {
		fileNode = fs.GetNode("/etc/group")
	}

	// If none of the common files exist, we'll skip the test
	if !fileNode.Exists() {
		t.Skip("No common system files found for testing")
	}

	// Cast to File interface
	file, ok := fileNode.(afs.File)
	if !ok {
		t.Fatal("Node does not implement File interface")
	}

	// Test Node interface methods
	_ = file.String()
	_ = file.GetName()
	_ = file.GetPath()
	_ = file.GetURI()
	_ = file.IsDir()
	_ = file.IsFile()
	_ = file.IsLink()
	_ = file.Exists()

	// Test GetParent method
	_ = file.GetParent()

	// Test GetFileSystem method
	_ = file.GetFileSystem()

	// Test GetMeta method
	_ = file.GetMeta(nil)

	// Test GetIO method
	_ = file.GetIO()
}
