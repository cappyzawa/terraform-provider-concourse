package concourse

import (
	"reflect"
	"testing"
)

func TestIfs2strs(t *testing.T) {
	t.Parallel()

	cases := map[string]struct {
		ifs    []interface{}
		expect []string
	}{
		"ifs has an index": {
			ifs:    []interface{}{"one"},
			expect: []string{"one"},
		},
		"ifs has two indexs": {
			ifs:    []interface{}{"one", "two"},
			expect: []string{"one", "two"},
		},
	}

	for name, test := range cases {
		test := test
		t.Run(name, func(t *testing.T) {
			t.Parallel()
			actual := ifs2strs(test.ifs)
			if !reflect.DeepEqual(actual, test.expect) {
				t.Errorf("output should be %v, but it is %v", test.expect, actual)
			}
		})
	}
}
