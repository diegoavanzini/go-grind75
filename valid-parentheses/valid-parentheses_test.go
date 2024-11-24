package validparentheses

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_ValidParetheses(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected bool
	}{
		{"Example 1", "()", true},
		{"Example 2", "()[]{}", true},
		{"Example 3", "(]", false},
		{"Example 4", "([])", true},
		{"Example 5", "[([{{()}}])]", true},
		{"Example 5", "[({{()}}])]", false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// ARRANGE
			// ACT
			actual := validateParentheses(tt.input)
			// ASSERT
			assert.Equal(t, tt.expected, actual)
		})
	}
}
