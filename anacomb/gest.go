package anacomb

import (
	"reflect"
	"slices"
	"testing"
)

func AssertEquals[K comparable](expected K, obtained K, t *testing.T) {
	if expected != obtained {
		t.Errorf("Expected %v, but got %v", expected, obtained)
	}
}

func AssertEqualSlice[K comparable](expected, obtained []K, t *testing.T) {
	if !slices.Equal(expected, obtained) {
		t.Errorf("Expected %v, but got %v", expected, obtained)
	}
}

func AssertNil(value any, t *testing.T) {
	v := reflect.ValueOf(value)
	if value != nil || (v.Kind() == reflect.Ptr || v.Kind() == reflect.Interface || v.Kind() == reflect.Slice || v.Kind() == reflect.Map) && !v.IsNil() {
		t.Errorf("Value is not nil")
	}
}

func AssertNotNil(value any, t *testing.T) {
	v := reflect.ValueOf(value)
	if value == nil || (v.Kind() == reflect.Ptr || v.Kind() == reflect.Interface || v.Kind() == reflect.Slice || v.Kind() == reflect.Map) && v.IsNil() {
		t.Errorf("Value is nil")
	}
}

func AssertTrue(value bool, t *testing.T) {
	if !value {
		t.Error("Expected true, but got false.")
	}
}

func AssertFalse(value bool, t *testing.T) {
	if value {
		t.Error("Expected false, but got true")
	}
}
