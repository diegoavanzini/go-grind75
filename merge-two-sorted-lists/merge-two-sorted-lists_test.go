package mergetwosortedlists

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewLists(t *testing.T) {
	tests := []struct {
		name     string
		list     []int
		expected string
	}{
		{"Example 1", []int{1, 2, 4}, "1, 2, 4"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			actual := NewList(tt.list)
			assert.Equal(t, tt.expected, actual.String())
		})
	}
}

func TestMergeTwoSortedLists(t *testing.T) {
	tests := []struct {
		name     string
		list1    *ListNode
		list2    *ListNode
		expected string
	}{
		{"Example 1", NewList([]int{1, 2, 4}), NewList([]int{1, 3, 4}), "1, 1, 2, 3, 4, 4"},
		{"Example 2", NewList([]int{}), NewList([]int{}), ""},
		{"Example 3", NewList([]int{}), NewList([]int{0}), "0"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			actual := mergeTwoLists(tt.list1, tt.list2)
			assert.Equal(t, tt.expected, actual.String())
		})
	}
}
