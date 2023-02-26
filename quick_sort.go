package godsa

func QuickSort(array []int) {
	if len(array) <= 1 {
		return
	}

	left, right := partition(array)

	QuickSort(left)
	QuickSort(right)
}

func partition(array []int) ([]int, []int) {
	pi := 0 // pivot final position
	high := len(array) - 1

	pivot := array[high]

	for i := 0; i < high; i++ {
		if array[i] <= pivot {
			array[i], array[pi] = array[pi], array[i]
			pi++
		}
	}

	array[pi], array[high] = array[high], array[pi]

	// ignore the pivot
	return array[:pi], array[pi+1:]
}
