package handlers

func FindPairs(nums []int, target int) [][]int {
	result := [][]int{}
	if len(nums) == 0 {
		return result
	}
	// keeping track of number and its occurence
	indexMap := make(map[int][]int)

	for i, num := range nums {
		diff := target - num
		// traversing the map array to find all the pairs
		if index, found := indexMap[diff]; found {
			for _, idx := range index {
				result = append(result, []int{idx, i})
			}
		}
		indexMap[num] = append(indexMap[num], i)
	}
	// returning the result
	return result

}
