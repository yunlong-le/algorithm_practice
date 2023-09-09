package main

func bubbleSort(arr []int) {
	if arr == nil || len(arr) == 0 {
		return
	}
	for i := len(arr) - 1; i > 0; i-- {
		for j := 0; j < i; j++ {
			if arr[j] > arr[j+1] {
				swap(arr, j, j+1)
			}
		}
	}
}
