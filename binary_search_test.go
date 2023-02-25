package godsa

import (
	"testing"
)

func TestBinarySearch(t *testing.T) {
	numbers := []int{1, 12, 23, 34, 45, 56, 67, 78, 89, 99, 100}

	testBinarySearch(t, numbers, 1, 0)
	testBinarySearch(t, numbers, 12, 1)
	testBinarySearch(t, numbers, 23, 2)
	testBinarySearch(t, numbers, 34, 3)
	testBinarySearch(t, numbers, 45, 4)
	testBinarySearch(t, numbers, 56, 5)
	testBinarySearch(t, numbers, 67, 6)
	testBinarySearch(t, numbers, 78, 7)
	testBinarySearch(t, numbers, 89, 8)
	testBinarySearch(t, numbers, 99, 9)
	testBinarySearch(t, numbers, 100, 10)

	testBinarySearch(t, numbers, 0, -1)
	testBinarySearch(t, numbers, 11, -1)
	testBinarySearch(t, numbers, 22, -1)
}

func testBinarySearch(t *testing.T, numbers []int, value int, want int) {
	t.Helper()

	got := BinarySearch(numbers, value)
	if got != want {
		t.Fatalf("BinarySearch(%v, %v) = %v, want %v", numbers, value, got, want)
	}
}
