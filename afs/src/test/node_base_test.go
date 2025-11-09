package main

import (
	"testing"

	"github.com/starter-go/v1/afs"
	"github.com/starter-go/v1/afs/libafs"
)

func TestNodeGetParent(t *testing.T) {

	afs.Init(libafs.DefaultDriver())
	fs := afs.Default()

	child := fs.GetNode("/foo/bar/example/")
	parent := child.GetParent()

	t.Logf("parent: [%v]", parent)

}

func TestNodeCountParents(t *testing.T) {

	afs.Init(libafs.DefaultDriver())
	fs := afs.Default()

	child := fs.GetNode("/foo/bar/example/")
	count := child.CountParents()

	t.Logf("child = [%v]", child)
	t.Logf("count_parent = %v", count)
}

func TestNodeListParents(t *testing.T) {

	afs.Init(libafs.DefaultDriver())
	fs := afs.Default()

	child := fs.GetNode("/foo/bar/example/")
	list := child.ListParents()

	t.Logf("child = [%v]", child)
	for i, it := range list {
		t.Logf("  parent[%d] = [%v]", i, it)
	}

}

func TestFSGetNode(t *testing.T) {

	namelist := []string{
		"",
		"/",
		"/foo",
		"/foo/bar",
	}

	afs.Init(libafs.DefaultDriver())
	fs := afs.Default()

	for idx, name := range namelist {
		node := fs.GetNode(afs.Path(name))
		t.Logf("\n  index=[%d], str=[%s], node=[%v]", idx, name, node)
	}

}
