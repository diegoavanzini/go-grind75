package insertionsort

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestInsertionSort(t *testing.T) {
	var numbers = []int{99, 44, 6, 2, 1, 5, 63, 87, 283, 4, 0}
	sorted := insertionSort(numbers)

	fmt.Println(sorted)

	assert.Equal(t, sorted[0], 0)
	assert.Equal(t, sorted[1], 1)
	assert.Equal(t, sorted[2], 2)
	assert.Equal(t, sorted[3], 4)
	assert.Equal(t, sorted[4], 5)
	assert.Equal(t, sorted[5], 6)
	assert.Equal(t, sorted[6], 44)
	assert.Equal(t, sorted[7], 63)
	assert.Equal(t, sorted[8], 87)
}
