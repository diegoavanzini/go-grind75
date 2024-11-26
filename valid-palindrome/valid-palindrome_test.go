package validpalindrome

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func BenchmarkIsPalindrome(b *testing.B) {
	isPalindrome("A man, a plan?////+=, a c!A@nal: Panama")
}

func Test_isPalindrome(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected bool
	}{
		{"Example 1", "A man, a plan, a canal: Panama", true},
		{"Example 2", "race a car", false},
		{"Example 3", "", true},
		{"Example 4", "A man, a plan?////+=, a c!A@nal: Panama", true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			actual := isPalindrome(tt.input)
			assert.Equal(t, tt.expected, actual)
		})
	}
}
