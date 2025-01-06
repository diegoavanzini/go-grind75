package bubblesort

func bubbleSort(numbers []int) []int {
	if len(numbers) == 0 {
		return numbers
	}
	for i := 0; i < len(numbers)-1; i++ {
		j := i + 1
		if numbers[i] > numbers[j] {
			numbers[i], numbers[j] = numbers[j], numbers[i]
			continue
		}
	}
	lastElement := len(numbers) - 1
	return append(bubbleSort(numbers[:lastElement]), numbers[lastElement])
}
