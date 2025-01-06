package selectionsort

func selectionSort(numbers []int) []int {
	var result = make([]int, len(numbers))
	if len(numbers) < 2 {
		return numbers
	}
	smallest := 0
	for i, number := range numbers {
		if number < numbers[smallest] {
			smallest = i
		}
	}
	copy(result, numbers)
	return append([]int{numbers[smallest]}, selectionSort(remove(result, smallest))...)
}

func remove(slice []int, index int) []int {
	return append(slice[:index], slice[index+1:]...)
}
