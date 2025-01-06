package bubblesort

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestBubbleSort(t *testing.T) {
	var numbers = []int{99, 44, 6, 2, 1, 5, 63, 87, 283, 4, 0}
	sorted := bubbleSort(numbers)

	fmt.Println(sorted)

	assert.Equal(t, sorted[0], 0)
}
