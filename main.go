package main

import "fmt"

func BinarySearch(array []int, val int) bool {
	right := len(array)
	left := 0
	for left < right-1 {
		mid := (right + left) / 2
		if val == array[mid-1] {
			return true
		}
		if val < array[mid-1] {
			right = mid
		}
		if val > array[mid-1] {
			left = mid
		}
	}

	return false
}

func main() {
	arr := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	fmt.Println(BinarySearch(arr, 5))
	fmt.Println(BinarySearch(arr, 11))
	fmt.Println(BinarySearch(arr, 1))
	fmt.Println(BinarySearch(arr, 0))
	fmt.Println(BinarySearch(arr, 8))
	fmt.Println(BinarySearch(arr, 2))
	fmt.Println(BinarySearch(arr, 25))
}
