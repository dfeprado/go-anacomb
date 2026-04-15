package anacomb

import (
	"reflect"
	"testing"
)

func AssertEquals[K comparable](expected K, obtained K, t *testing.T) {
	if expected != obtained {
		t.Errorf("Expected %v, but got %v", expected, obtained)
	}
}

func AssertNotNil(value any, t *testing.T) {
	v := reflect.ValueOf(value)
	if value == nil || (v.Kind() == reflect.Ptr || v.Kind() == reflect.Interface || v.Kind() == reflect.Slice || v.Kind() == reflect.Map) && v.IsNil() {
		t.Errorf("Value is not nil")
	}
}

