package linkedList

import (
	"reflect"
	"testing"

	"Leetcode/algorithms/kit"
)

// run: go test -v base.go 2.*
func Test_addTwoNumbers(t *testing.T) {
	cases := []struct {
		name     string
		input    [][]int
		expected []int
	}{
		{
			name:     "x1",
			input:    [][]int{{2, 4, 3}, {5, 6, 4}},
			expected: []int{7, 0, 8},
		},
	}

	for _, tt := range cases {
		t.Run(tt.name, func(t *testing.T) {
			l1 := kit.Ints2List(tt.input[0])
			l2 := kit.Ints2List(tt.input[1])
			output := addTwoNumbers(l1, l2)
			out2ints := kit.List2Ints(output)
			if !reflect.DeepEqual(out2ints, tt.expected) {
				t.Errorf("addTwoNumbers(%v, %v)=%v, expected=%v", tt.input[0], tt.input[1], out2ints, tt.expected)
			}
		})
	}
}
