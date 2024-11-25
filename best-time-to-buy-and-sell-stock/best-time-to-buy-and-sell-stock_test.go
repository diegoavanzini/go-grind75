package besttimetobuyandsellstock

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_maxProfit(t *testing.T) {
	tests := []struct {
		name     string
		input    []int
		expected int
	}{
		{"Example 1", []int{7, 1, 5, 3, 6, 4}, 5},
		{"Example 2", []int{7, 6, 4, 3, 1}, 0},
		{"Example 3", []int{7, 8, 4, 3, 1}, 1},
		{"Example 4", []int{7, 8, 10, 1}, 3},
		{"Example ", []int{7, 8, 4, 10, 1}, 6},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			actual := maxProfit(tt.input)
			assert.Equal(t, tt.expected, actual)
		})
	}
}
