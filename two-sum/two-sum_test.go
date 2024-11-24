package twosum

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_TwoSum(t *testing.T) {
	tests := []struct {
		name     string
		nums     []int
		target   int
		expected []int
	}{
		{"example 1", []int{3, 2, 4}, 6, []int{1, 2}},
		{"example 2", []int{2, 7, 11, 15}, 9, []int{0, 1}},
		{"example 3", []int{3, 3}, 6, []int{0, 1}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// ACT
			got := twoSum(tt.nums, tt.target)
			// ASSERT
			assert.Equal(t, tt.expected, got)
		})
	}
}
