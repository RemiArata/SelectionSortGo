package ss

func SelectionSort(data []int) []int {

	var smallestVal, smallestIdx int

	for j := 0; j < len(data); j++ {
		smallestIdx, smallestVal = findMin(data, j)
		// fmt.Println("The smallest value is", smallestVal, "at idx", smallestIdx)
		data[j], data[smallestIdx] = smallestVal, data[j]
	}
	return data
}

func findMin(data []int, startIdx int) (int, int) { // return is loc, val pair
	minVal := data[startIdx]
	minIdx := startIdx
	for i := startIdx; i < len(data); i++ {
		if data[i] < minVal {
			minVal = data[i]
			minIdx = i
		}
	}
	return minIdx, minVal
}
