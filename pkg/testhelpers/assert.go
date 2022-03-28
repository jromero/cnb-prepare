package testhelpers

import (
	"reflect"
	"testing"
)

func Equals(t *testing.T, val, expected interface{}) {
	if !reflect.DeepEqual(val, expected) {
		t.Fatalf(
			"Expected\n-----\n%#v\n-----\nbut got\n-----\n%#v\n",
			expected, val,
		)
	}
}
