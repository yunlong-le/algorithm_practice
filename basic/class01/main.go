package main

import (
	"fmt"
	"math/rand"
	"sort"
	"time"
)

func swap(arr []int, i int, j int) {
	temp := arr[i]
	arr[i] = arr[j]
	arr[j] = temp
}
func comparator(arr []int) {
	sort.Ints(arr)
}

func generateRandomArray(maxSize, maxValue int) []int {
	rand.Seed(time.Now().UnixNano())
	length := rand.Intn(maxSize + 1)
	arr := make([]int, length)
	for i := 0; i < length; i++ {
		arr[i] = rand.Intn(2*maxValue+1) - maxValue
	}
	return arr
}

func copyArray(arr []int) []int {
	if arr == nil {
		return nil
	}
	res := make([]int, len(arr))
	copy(res, arr)
	return res
}

func isEqual(arr1, arr2 []int) bool {
	if (arr1 == nil && arr2 != nil) || (arr1 != nil && arr2 == nil) {
		return false
	}
	if arr1 == nil && arr2 == nil {
		return true
	}
	if len(arr1) != len(arr2) {
		return false
	}
	for i := 0; i < len(arr1); i++ {
		if arr1[i] != arr2[i] {
			return false
		}
	}
	return true
}

func printArray(arr []int) {
	if arr == nil {
		return
	}
	for _, val := range arr {
		fmt.Printf("%d ", val)
	}
	fmt.Println()
}

func main() {
	testTime := 100000
	maxSize := 100
	maxValue := 100
	succeed := true
	for i := 0; i < testTime; i++ {
		arr1 := generateRandomArray(maxSize, maxValue)
		arr2 := copyArray(arr1)
		//selectionSort(arr1)
		//bubbleSort(arr1)
		insertionSort(arr1)
		comparator(arr2)
		if !isEqual(arr1, arr2) {
			succeed = false
			printArray(arr1)
			printArray(arr2)
			break
		}
	}
	if succeed {
		fmt.Println("Nice!")
	} else {
		fmt.Println("Fucking fucked!")
	}

	arr := generateRandomArray(maxSize, maxValue)
	printArray(arr)
	//selectionSort(arr)
	//bubbleSort(arr)
	insertionSort(arr)
	printArray(arr)
}
