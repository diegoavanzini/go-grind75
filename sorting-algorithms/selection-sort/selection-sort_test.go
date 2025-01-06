package selectionsort

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSelectionSort(t *testing.T) {
	var numbers = []int{99, 44, 6, 2, 1, 5, 63, 87, 283, 4, 0}

	fmt.Printf("%+v\n", numbers)
	got := selectionSort(numbers)
	fmt.Printf("%+v", got)
	assert.Equal(t, 0, got[0])
	assert.Equal(t, 1, got[1])
	assert.Equal(t, 2, got[2])
	assert.Equal(t, 4, got[3])
	assert.Equal(t, 5, got[4])
	assert.Equal(t, 6, got[5])
	assert.Equal(t, 44, got[6])
	assert.Equal(t, 63, got[7])
	assert.Equal(t, 87, got[8])
	assert.Equal(t, 99, got[9])
	assert.Equal(t, 283, got[10])
}
func TestSlice(t *testing.T) {
	var numbers = []int{99, 44, 6, 2, 0, 5, 63, 87, 283, 4, 1}

	fmt.Printf("%+v", numbers)

	got := remove(numbers, 5)
	fmt.Printf("%+v", got)
}
