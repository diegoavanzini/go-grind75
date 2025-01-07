package insertionsort

func insertionSort(numbers []int) []int {
	if len(numbers) < 2 {
		return numbers
	}
	for current := 0; current < len(numbers)-1; current++ {
		next := current + 1
		nextNumber := numbers[next]
		if numbers[current] > nextNumber {
			numbers = remove(numbers, next)
			numbers = insertIn(numbers, nextNumber)
			return insertionSort(numbers)
		}
	}
	return numbers
}

func insertIn(slice []int, toInsert int) (result []int) {
	result = make([]int, len(slice))
	for index, number := range slice {
		if number > toInsert {
			var after = make([]int, len(slice[index:]))
			copy(after, slice[index:])
			before := slice[:index]
			atFirstPosition := append(before, toInsert)
			result = append(atFirstPosition, after...)
			return
		}
	}
	return
}

func remove(slice []int, index int) []int {
	return append(slice[:index], slice[index+1:]...)
}
