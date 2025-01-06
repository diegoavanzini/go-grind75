package sortingsmalldatasets

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSplitArray(t *testing.T) {
	partecipants := []Participant{
		{Score: 7},
		{Score: 2},
		{Score: 1},
		{Score: 8},
		{Score: 6},
		{Score: 3},
		{Score: 5},
		{Score: 4},
	}
	fmt.Println(partecipants)
	la := partecipants[:5]
	ga := partecipants[5 : len(partecipants)-1]
	fmt.Println(la)
	fmt.Println(partecipants[len(partecipants)-1])
	fmt.Println(ga)
}

func TestMyQuickSort(t *testing.T) {
	partecipants := []Participant{
		{Score: 7},
		{Score: 2},
		{Score: 1},
		{Score: 8},
		{Score: 6},
		{Score: 3},
		{Score: 5},
		{Score: 4},
	}
	fmt.Println(partecipants)
	// QuickSort(partecipants, 0, 9)
	orderedPartecipants := MyQuickSort(partecipants)

	fmt.Println(orderedPartecipants)
	assert.Equal(t, 1, orderedPartecipants[0].Score)
	assert.Equal(t, 2, orderedPartecipants[1].Score)
	assert.Equal(t, 3, orderedPartecipants[2].Score)
	assert.Equal(t, 4, orderedPartecipants[3].Score)
	assert.Equal(t, 5, orderedPartecipants[4].Score)
	assert.Equal(t, 6, orderedPartecipants[5].Score)
	assert.Equal(t, 7, orderedPartecipants[6].Score)
	assert.Equal(t, 8, orderedPartecipants[7].Score)
}
