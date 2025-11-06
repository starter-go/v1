package mock

import (
	"testing"

	"github.com/starter-go/v1/buckets"
)

func TestMockDriverRegistration(t *testing.T) {
	driver := &TheMockDriver{}
	registration := driver.Registration()

	if registration.Name != "mock" {
		t.Errorf("Expected driver name to be 'mock', got '%s'", registration.Name)
	}

	if !registration.Enabled {
		t.Error("Expected driver to be enabled")
	}

	if registration.Priority != 0 {
		t.Errorf("Expected driver priority to be 0, got %d", registration.Priority)
	}

	if registration.Driver != driver {
		t.Error("Expected driver reference to be the same")
	}
}

func TestMockDriverOpenBucket(t *testing.T) {
	driver := &TheMockDriver{}
	ctx := &buckets.Context{
		Config: buckets.Configuration{
			Driver: "mock",
		},
	}

	bucket, err := driver.OpenBucket(ctx)
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	if bucket == nil {
		t.Fatal("Expected bucket to be created")
	}

	// Test that we can get the context back
	if bucket.GetContext() != ctx {
		t.Error("Expected to get the same context back")
	}

	// Test that we can get an object
	obj := bucket.GetObject("test-object")
	if obj == nil {
		t.Fatal("Expected object to be created")
	}

	if obj.Name() != "test-object" {
		t.Errorf("Expected object name to be 'test-object', got '%s'", obj.Name())
	}
}

func TestMockObjectOperations(t *testing.T) {
	driver := &TheMockDriver{}
	ctx := &buckets.Context{
		Config: buckets.Configuration{
			Driver: "mock",
		},
	}

	bucket, _ := driver.OpenBucket(ctx)
	obj := bucket.GetObject("test-object")

	// Test that object doesn't exist initially
	exists, err := obj.Exists()
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}
	if exists {
		t.Error("Expected object to not exist initially")
	}

	// Test fetching non-existent object
	_, _, err = obj.Fetch()
	if err != ErrObjectNotFound {
		t.Errorf("Expected ErrObjectNotFound, got %v", err)
	}

	// Test putting an object
	meta := &buckets.ObjectMeta{
		Name: "test-object",
	}
	data := &buckets.ObjectData{
		Length: 100,
	}
	err = obj.Put(meta, data)
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	// Test that object now exists
	exists, err = obj.Exists()
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}
	if !exists {
		t.Error("Expected object to exist after Put")
	}

	// Test removing an object
	err = obj.Remove()
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	// Test that object no longer exists
	exists, err = obj.Exists()
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}
	if exists {
		t.Error("Expected object to not exist after Remove")
	}
}
