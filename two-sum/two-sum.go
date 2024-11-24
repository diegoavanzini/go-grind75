package twosum

func twoSum(nums []int, target int) []int {
	var result = []int{}
	var foundedNums = map[int]int{}
	for i, v := range nums {
		if manca, ok := foundedNums[v]; ok {
			result = append(result, manca, i)
			return result
		}
		manca := target - v
		if manca > 0 {
			foundedNums[manca] = i
		}
	}
	return result
}
