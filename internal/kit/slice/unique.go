package slice

import "sort"

type Unique struct{}

func OrderToRemoveDuplicates(numbers []int) (result []int) {
	result = []int{}
	sort.Sort(sort.IntSlice(numbers))
	for i, number := range numbers {
		if i == 0 {
			result = append(result, number)
		}
		if i != 0 && number != numbers[i-1] {
			result = append(result, number)
		}
	}
	return result
}
