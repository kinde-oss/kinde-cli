package reflectvisitor

import (
	"reflect"
	"testing"
)

// Helper types for testing
type TestStruct struct {
	Field1 string
	Field2 int
}

type OptTestStruct struct {
	Set   bool
	Value TestStruct
}

func TestVisitor_Visit_Struct(t *testing.T) {
	root := TestStruct{Field1: "foo", Field2: 42}
	val := reflect.ValueOf(root)
	visitor := NewVisitor(reflect.TypeOf(root))

	called := false
	visitor.Visit(&val, func(w Walker[reflect.Type]) []Walker[reflect.Type] {
		called = true
		if w.T1.Name() != "TestStruct" {
			t.Errorf("Expected type TestStruct, got %s", w.T1.Name())
		}
		return nil
	})

	if !called {
		t.Error("Visitor did not call function for struct")
	}
}

func TestVisitor_Visit_OptStruct(t *testing.T) {
	root := OptTestStruct{
		Set: true,
		Value: TestStruct{
			Field1: "bar",
			Field2: 99,
		},
	}
	val := reflect.ValueOf(root)
	visitor := NewVisitor(reflect.TypeOf(root))

	var visitedTypes []string
	visitor.Visit(&val, func(w Walker[reflect.Type]) []Walker[reflect.Type] {
		visitedTypes = append(visitedTypes, w.T1.Name())
		return nil
	})

	if len(visitedTypes) == 0 {
		t.Error("Expected at least one type to be visited")
	}
	if visitedTypes[0] != "TestStruct" {
		t.Errorf("Expected first visited type to be TestStruct, got %s", visitedTypes[0])
	}
}

func TestVisitor_Visit_NonStruct(t *testing.T) {
	root := 123
	val := reflect.ValueOf(root)
	visitor := NewVisitor(reflect.TypeOf(root))

	called := false
	visitor.Visit(&val, func(w Walker[reflect.Type]) []Walker[reflect.Type] {
		called = true
		return nil
	})

	if called {
		t.Error("Visitor should not call function for non-struct type")
	}
}

func TestVisitor_Visit_PointerToStruct(t *testing.T) {
	root := &TestStruct{Field1: "baz", Field2: 7}
	val := reflect.ValueOf(root)
	visitor := NewVisitor(reflect.TypeOf(root))

	called := false
	visitor.Visit(&val, func(w Walker[reflect.Type]) []Walker[reflect.Type] {
		called = true
		if w.T1.Name() != "TestStruct" {
			t.Errorf("Expected type TestStruct, got %s", w.T1.Name())
		}
		return nil
	})

	if !called {
		t.Error("Visitor did not call function for pointer to struct")
	}
}
