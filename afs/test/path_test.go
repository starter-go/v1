package main

import (
	"testing"

	"github.com/starter-go/v1/afs"
)

func TestPathString(t *testing.T) {
	path := afs.Path("/home/user/documents")
	expected := "/home/user/documents"
	actual := path.String()
	if actual != expected {
		t.Errorf("Path.String() = %s; expected %s", actual, expected)
	}
}

func TestPathElements(t *testing.T) {
	// Test with Unix-style path
	path := afs.Path("/home/user/documents")
	elements := path.Elements()
	expectedLength := 4 // Empty first element + "home" + "user" + "documents"
	if len(elements) != expectedLength {
		t.Errorf("Path.Elements() length = %d; expected %d", len(elements), expectedLength)
	}

	expectedElements := []string{"", "home", "user", "documents"}
	for i, expected := range expectedElements {
		if string(elements[i]) != expected {
			t.Errorf("Path.Elements()[%d] = %s; expected %s", i, string(elements[i]), expected)
		}
	}

	// Test with Windows-style path
	path = afs.Path("\\home\\user\\documents")
	elements = path.Elements()
	if len(elements) != expectedLength {
		t.Errorf("Path.Elements() length = %d; expected %d", len(elements), expectedLength)
	}

	// Test with mixed separators
	path = afs.Path("/home\\user/documents")
	elements = path.Elements()
	if len(elements) != expectedLength {
		t.Errorf("Path.Elements() length = %d; expected %d", len(elements), expectedLength)
	}
}

func TestPathNormalize(t *testing.T) {
	// Test normal path
	path := afs.Path("/home/user/documents")
	normalized, err := path.Normalize()
	if err != nil {
		t.Errorf("Path.Normalize() returned error: %v", err)
	}
	expected := afs.Path("/home/user/documents")
	if normalized != expected {
		t.Errorf("Path.Normalize() = %s; expected %s", normalized, expected)
	}

	// Test path with double dots
	path = afs.Path("/home/user/../documents")
	normalized, err = path.Normalize()
	if err != nil {
		t.Errorf("Path.Normalize() returned error: %v", err)
	}
	expected = afs.Path("/home/documents")
	if normalized != expected {
		t.Errorf("Path.Normalize() = %s; expected %s", normalized, expected)
	}

	// Test path with dots
	path = afs.Path("/home/./user/documents")
	normalized, err = path.Normalize()
	if err != nil {
		t.Errorf("Path.Normalize() returned error: %v", err)
	}
	expected = afs.Path("/home/user/documents")
	if normalized != expected {
		t.Errorf("Path.Normalize() = %s; expected %s", normalized, expected)
	}
}

func TestPathElementString(t *testing.T) {
	element := afs.PathElement("home")
	expected := "home"
	actual := element.String()
	if actual != expected {
		t.Errorf("PathElement.String() = %s; expected %s", actual, expected)
	}
}

func TestPathElementIsEmpty(t *testing.T) {
	// Test empty element
	element := afs.PathElement("")
	if !element.IsEmpty() {
		t.Error("PathElement.IsEmpty() should return true for empty element")
	}

	// Test non-empty element
	element = afs.PathElement("home")
	if element.IsEmpty() {
		t.Error("PathElement.IsEmpty() should return false for non-empty element")
	}
}

func TestPathElementIsDot(t *testing.T) {
	// Test dot element
	element := afs.PathElement(".")
	if !element.IsDot() {
		t.Error("PathElement.IsDot() should return true for dot element")
	}

	// Test non-dot element
	element = afs.PathElement("home")
	if element.IsDot() {
		t.Error("PathElement.IsDot() should return false for non-dot element")
	}
}

func TestPathElementIsDoubleDot(t *testing.T) {
	// Test double dot element
	element := afs.PathElement("..")
	if !element.IsDoubleDot() {
		t.Error("PathElement.IsDoubleDot() should return true for double dot element")
	}

	// Test non-double dot element
	element = afs.PathElement("home")
	if element.IsDoubleDot() {
		t.Error("PathElement.IsDoubleDot() should return false for non-double dot element")
	}
}

func TestPathElementIsUserHome(t *testing.T) {
	// Test user home element
	element := afs.PathElement("~")
	if !element.IsUserHome() {
		t.Error("PathElement.IsUserHome() should return true for user home element")
	}

	// Test non-user home element
	element = afs.PathElement("home")
	if element.IsUserHome() {
		t.Error("PathElement.IsUserHome() should return false for non-user home element")
	}
}

func TestPathElementListPath(t *testing.T) {
	elements := afs.PathElementList{
		afs.PathElement("home"),
		afs.PathElement("user"),
		afs.PathElement("documents"),
	}
	path := elements.Path()
	expected := afs.Path("/home/user/documents")
	if path != expected {
		t.Errorf("PathElementList.Path() = %s; expected %s", path, expected)
	}

	// Test empty list
	elements = afs.PathElementList{}
	path = elements.Path()
	expected = afs.Path("")
	if path != expected {
		t.Errorf("PathElementList.Path() = %s; expected %s", path, expected)
	}
}

func TestPathElementListGetParent(t *testing.T) {
	elements := afs.PathElementList{
		afs.PathElement("home"),
		afs.PathElement("user"),
		afs.PathElement("documents"),
	}
	parent, err := elements.GetParent()
	if err != nil {
		t.Errorf("PathElementList.GetParent() returned error: %v", err)
	}
	expectedLength := 2
	if len(parent) != expectedLength {
		t.Errorf("PathElementList.GetParent() length = %d; expected %d", len(parent), expectedLength)
	}

	expectedElements := []string{"home", "user"}
	for i, expected := range expectedElements {
		if string(parent[i]) != expected {
			t.Errorf("PathElementList.GetParent()[%d] = %s; expected %s", i, string(parent[i]), expected)
		}
	}

	// Test with single element
	elements = afs.PathElementList{afs.PathElement("home")}
	parent, err = elements.GetParent()
	if err != nil {
		t.Errorf("PathElementList.GetParent() returned error: %v", err)
	}
	expectedLength = 0
	if len(parent) != expectedLength {
		t.Errorf("PathElementList.GetParent() length = %d; expected %d", len(parent), expectedLength)
	}

	// Test with empty list
	elements = afs.PathElementList{}
	_, err = elements.GetParent()
	if err == nil {
		t.Error("PathElementList.GetParent() should return error for empty list")
	}
}

func TestPathElementListGetChild(t *testing.T) {
	elements := afs.PathElementList{
		afs.PathElement("home"),
		afs.PathElement("user"),
	}
	child := elements.GetChild(afs.PathElement("documents"))
	expectedLength := 3
	if len(child) != expectedLength {
		t.Errorf("PathElementList.GetChild() length = %d; expected %d", len(child), expectedLength)
	}

	expectedElements := []string{"home", "user", "documents"}
	for i, expected := range expectedElements {
		if string(child[i]) != expected {
			t.Errorf("PathElementList.GetChild()[%d] = %s; expected %s", i, string(child[i]), expected)
		}
	}
}

func TestPathElementListIsAbsolute(t *testing.T) {
	// Test absolute path (empty first element)
	elements := afs.PathElementList{
		afs.PathElement(""),
		afs.PathElement("home"),
		afs.PathElement("user"),
	}
	if !elements.IsAbsolute() {
		t.Error("PathElementList.IsAbsolute() should return true for absolute path")
	}

	// Test relative path (starts with dot)
	elements = afs.PathElementList{
		afs.PathElement("."),
		afs.PathElement("home"),
		afs.PathElement("user"),
	}
	if elements.IsAbsolute() {
		t.Error("PathElementList.IsAbsolute() should return false for relative path")
	}

	// Test relative path (starts with double dot)
	elements = afs.PathElementList{
		afs.PathElement(".."),
		afs.PathElement("home"),
		afs.PathElement("user"),
	}
	if elements.IsAbsolute() {
		t.Error("PathElementList.IsAbsolute() should return false for relative path")
	}

	// Test that a list starting with a normal element is considered absolute
	// (because IsRelative() returns false when the first non-empty element is not "." or "..")
	elements = afs.PathElementList{
		afs.PathElement("home"),
		afs.PathElement("user"),
	}
	if !elements.IsAbsolute() {
		t.Error("PathElementList.IsAbsolute() should return true for path starting with normal element")
	}

	// Test another absolute path (only empty elements)
	elements = afs.PathElementList{
		afs.PathElement(""),
		afs.PathElement(""),
	}
	if !elements.IsAbsolute() {
		t.Error("PathElementList.IsAbsolute() should return true for absolute path")
	}
}

func TestPathElementListIsRelative(t *testing.T) {
	// Test relative path (starts with dot)
	elements := afs.PathElementList{
		afs.PathElement("."),
		afs.PathElement("home"),
		afs.PathElement("user"),
	}
	if !elements.IsRelative() {
		t.Error("PathElementList.IsRelative() should return true for relative path")
	}

	// Test relative path (starts with double dot)
	elements = afs.PathElementList{
		afs.PathElement(".."),
		afs.PathElement("home"),
		afs.PathElement("user"),
	}
	if !elements.IsRelative() {
		t.Error("PathElementList.IsRelative() should return true for relative path")
	}

	// Test that a list starting with a normal element is NOT considered relative
	// (because IsRelative() returns false when the first non-empty element is not "." or "..")
	elements = afs.PathElementList{
		afs.PathElement("home"),
		afs.PathElement("user"),
	}
	if elements.IsRelative() {
		t.Error("PathElementList.IsRelative() should return false for path starting with normal element")
	}

	// Test absolute path (empty first element)
	elements = afs.PathElementList{
		afs.PathElement(""),
		afs.PathElement("home"),
		afs.PathElement("user"),
	}
	if elements.IsRelative() {
		t.Error("PathElementList.IsRelative() should return false for absolute path")
	}
}

func TestPathElementListNormalize(t *testing.T) {
	// Test normal path
	elements := afs.PathElementList{
		afs.PathElement("home"),
		afs.PathElement("user"),
		afs.PathElement("documents"),
	}
	normalized, err := elements.Normalize()
	if err != nil {
		t.Errorf("PathElementList.Normalize() returned error: %v", err)
	}
	expectedLength := 3
	if len(normalized) != expectedLength {
		t.Errorf("PathElementList.Normalize() length = %d; expected %d", len(normalized), expectedLength)
	}

	// Test path with dots
	elements = afs.PathElementList{
		afs.PathElement("home"),
		afs.PathElement("."),
		afs.PathElement("user"),
		afs.PathElement("documents"),
	}
	normalized, err = elements.Normalize()
	if err != nil {
		t.Errorf("PathElementList.Normalize() returned error: %v", err)
	}
	expectedLength = 3
	if len(normalized) != expectedLength {
		t.Errorf("PathElementList.Normalize() length = %d; expected %d", len(normalized), expectedLength)
	}

	// Test path with double dots
	elements = afs.PathElementList{
		afs.PathElement("home"),
		afs.PathElement("user"),
		afs.PathElement(".."),
		afs.PathElement("documents"),
	}
	normalized, err = elements.Normalize()
	if err != nil {
		t.Errorf("PathElementList.Normalize() returned error: %v", err)
	}
	expectedLength = 2
	if len(normalized) != expectedLength {
		t.Errorf("PathElementList.Normalize() length = %d; expected %d", len(normalized), expectedLength)
	}
	expectedElements := []string{"home", "documents"}
	for i, expected := range expectedElements {
		if string(normalized[i]) != expected {
			t.Errorf("PathElementList.Normalize()[%d] = %s; expected %s", i, string(normalized[i]), expected)
		}
	}

	// Test path with user home
	elements = afs.PathElementList{
		afs.PathElement("~"),
		afs.PathElement("documents"),
	}
	normalized, err = elements.Normalize()
	if err != nil {
		t.Errorf("PathElementList.Normalize() returned error: %v", err)
	}
	// We can't predict the exact home directory path, but we can check that it's not empty
	if len(normalized) == 0 {
		t.Error("PathElementList.Normalize() should not return empty list for user home path")
	}
}
