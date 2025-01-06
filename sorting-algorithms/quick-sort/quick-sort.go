package sortingsmalldatasets

type Participant struct {
	Score int
}

// QuickSort sorts the participants slice in descending order of scores.
// It uses a divide-and-conquer approach for efficient sorting.
func QuickSort(participants []Participant, low, high int) {
	if low < high {
		// Find the pivot position
		p := partition(participants, low, high)
		// Recursively sort the left and right subarrays
		QuickSort(participants, low, p-1)
		QuickSort(participants, p+1, high)
	}
}

// partition rearranges the slice such that elements greater than or equal to the pivot
// appear before it, and elements less than the pivot come after it.
func partition(participants []Participant, low, high int) int {
	pivot := participants[high].Score // Select the pivot as the last element
	i := low - 1                      // Initialize pointer for greater elements

	for j := low; j < high; j++ {
		// If the current element's score is greater than or equal to the pivot
		if participants[j].Score >= pivot {
			i++ // Move the pointer
			// Swap the current element with the element at pointer i
			participants[i], participants[j] = participants[j], participants[i]
		}
	}

	// Finally, place the pivot in its correct position
	participants[i+1], participants[high] = participants[high], participants[i+1]
	return i + 1 // Return the index of the pivot
}

func MyQuickSort(partecipants []Participant) (ordered []Participant) {
	if len(partecipants) < 2 {
		return partecipants
	}
	pivot := len(partecipants) - 1
	var i int = -1
	for j := 0; j < len(partecipants); j++ {
		if partecipants[j].Score < partecipants[pivot].Score {
			i++
			if i == j {
				continue
			}
			partecipants[i], partecipants[j] = partecipants[j], partecipants[i]
		}
	}
	var minorPartition, greaterPartition []Participant
	minorPartition = append(minorPartition, partecipants[:i+1]...)
	greaterPartition = append(greaterPartition, partecipants[i+1:len(partecipants)-1]...)
	ordered = append(append(
		MyQuickSort(minorPartition),
		partecipants[pivot]),
		MyQuickSort(greaterPartition)...)
	return
}
