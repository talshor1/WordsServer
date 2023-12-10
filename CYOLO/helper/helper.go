package helper

func QuickSort(arr []int) {
	if len(arr) < 2 {
		return
	}

	pivotIndex := len(arr) / 2
	pivot := arr[pivotIndex]

	var left, right []int

	for i, value := range arr {
		if i == pivotIndex {
			continue
		}

		if value <= pivot {
			left = append(left, value)
		} else {
			right = append(right, value)
		}
	}

	QuickSort(left)
	QuickSort(right)

	copy(arr, append(append(left, pivot), right...))
}
