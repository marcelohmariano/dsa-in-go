package godsa

func BinarySearch(array []int, value int) int {
	low := 0
	high := len(array) - 1

	for low <= high {
		mid := (low + high) / 2

		if value > array[mid] {
			low = mid + 1
			continue
		}

		if value < array[mid] {
			high = mid - 1
			continue
		}

		return mid
	}

	return -1
}
