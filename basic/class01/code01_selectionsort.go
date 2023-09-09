package main

func selectionSort(arr []int) {
	if arr == nil || len(arr) == 0 {
		return
	}
	for i := 0; i < len(arr)-1; i++ {
		minIndex := i
		for j := i + 1; j < len(arr); j++ {
			if arr[j] < arr[minIndex] {
				minIndex = j
			}
		}
		swap(arr, minIndex, i)
	}
}
