package tree

import (
	"reflect"
	"testing"

	"Leetcode/algorithms/kit"
)

type entry94 struct {
	name     string
	input    []int
	expected []int
}

func buildInorderTraversalCase() []entry94 {
	cases := []entry94{
		{
			name:     "x1",
			input:    []int{1, NULL, 2, 3},
			expected: []int{1, 3, 2},
		},
		{
			name:     "x2",
			input:    []int{1, 2, 3, 4, 5, NULL, 7},
			expected: []int{4, 2, 5, 1, 3, 7},
		},
		{
			name:     "x3",
			input:    []int{1, 2, 3, 4, 5, 6, 7},
			expected: []int{4, 2, 5, 1, 6, 3, 7},
		},
	}
	return cases
}

// run: go test -v base.go 94.*
func Test_inorderTraversal(t *testing.T) {
	for _, tt := range buildInorderTraversalCase() {
		t.Run(tt.name, func(t *testing.T) {
			root := kit.Ints2Tree(tt.input)
			if output := inorderTraversal(root); !reflect.DeepEqual(output, tt.expected) {
				t.Errorf("inorderTraversal(%v)=%v, expected=%v", tt.input, output, tt.expected)
			}
		})
	}
}

func Test_Tree2Inorder(t *testing.T) {
	for _, tt := range buildInorderTraversalCase() {
		t.Run(tt.name, func(t *testing.T) {
			root := kit.Ints2Tree(tt.input)
			if output := kit.Tree2Inorder(root); !reflect.DeepEqual(output, tt.expected) {
				t.Errorf("Tree2Inorder(%v)=%v, expected=%v", tt.input, output, tt.expected)
			}
		})
	}
}
