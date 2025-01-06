package invertbinarytree

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_invertTree(t *testing.T) {

	tests := []struct {
		name  string
		input *TreeNode
		want  *TreeNode
	}{
		{"Example 1", NewTreeNode([]int{4, 2, 7, 1, 3, 6, 9}), NewTreeNode([]int{4, 7, 2, 9, 6, 3, 1})},
		{"Example 2", NewTreeNode([]int{2, 1, 3}), NewTreeNode([]int{2, 3, 1})},
		{"Example 3", NewTreeNode([]int{}), NewTreeNode([]int{})},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := invertTree(tt.input)
			assert.Equal(t, tt.want, got)
		})
	}

}
