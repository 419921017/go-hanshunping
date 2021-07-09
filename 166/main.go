/*
 * @Descripttion: your project
 * @version: 1.0
 * @Author: power_840
 * @Date: 2021-07-08 21:53:17
 * @LastEditors: power_840
 * @LastEditTime: 2021-07-08 22:11:10
 */
package main

import "fmt"

func main() {
	arr := []int{1, 8, 10, 89, 1000, 1234, 3321}
	fmt.Println(binarySearch(arr, 10))
}

func binarySearch(arr []int, target int) int {
	index := -1
	low := 0
	high := len(arr) - 1

	for low <= high {
		middle := (low + high) / 2
		if arr[middle] == target {
			index = middle
			return index
		} else if target < arr[middle] {
			high = middle - 1
		} else {
			low = middle + 1
		}
	}

	return index
}
