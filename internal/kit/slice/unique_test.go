package slice

import (
	"reflect"
	"testing"
)

func Test(t *testing.T) {
	t.Run("Sorting and removing duplicates", func(t *testing.T) {
		expected := []int{1, 2, 3}

		numbers := []int{3, 2, 1, 1, 2, 3}
		result := SortAndRemoveDuplicatesNumbers(numbers)

		if !reflect.DeepEqual(expected, result) {
			t.Errorf("Expected %v | Actual %v", expected, result)
		}
	})

	t.Run("Sorting and with any to remove", func(t *testing.T) {
		expected := []int{1, 2, 3}

		numbers := []int{3, 2, 1}
		result := SortAndRemoveDuplicatesNumbers(numbers)

		if !reflect.DeepEqual(expected, result) {
			t.Errorf("Expected %v | Actual %v", expected, result)
		}
	})

	t.Run("Removing with no needs to sort", func(t *testing.T) {
		expected := []int{2}

		numbers := []int{2}
		result := SortAndRemoveDuplicatesNumbers(numbers)

		if !reflect.DeepEqual(expected, result) {
			t.Errorf("Expected %v | Actual %v", expected, result)
		}
	})

}
