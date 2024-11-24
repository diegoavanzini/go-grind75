package validparentheses

func validateParentheses(input string) bool {
	var counters = map[rune]rune{
		'(': ')',
		'[': ']',
		'{': '}',
	}
	var stackOpens = []rune{}
	var nextToCLose rune
	for _, v := range input {
		if value, ok := counters[v]; ok {
			stackOpens = append(stackOpens, v)
			nextToCLose = value
			continue
		}
		if v != nextToCLose {
			return false
		}
		if len(stackOpens) > 1 {
			stackOpens = stackOpens[:len(stackOpens)-1]
			nextToCLose = counters[stackOpens[len(stackOpens)-1]]
		}
	}
	return true
}
