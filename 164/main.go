/*
 * @Descripttion: your project
 * @version: 1.0
 * @Author: power_840
 * @Date: 2021-06-08 22:00:54
 * @LastEditors: power_840
 * @LastEditTime: 2021-07-08 21:11:40
 */
package main

import "fmt"

func main() {
	nums := []int{23, 32, 53, 63, 12, 74, 77, 89, 99}
	BubbleSort(&nums)
	fmt.Println(nums)
}

func BubbleSort(arr *[]int) []int {
	for i := 0; i < len(*arr)-1; i++ {
		for j := i + 1; j < len(*arr)-i-1; j++ {
			if (*arr)[i] > (*arr)[j] {
				(*arr)[i], (*arr)[j] = (*arr)[j], (*arr)[i]
			}
		}
	}
	return *arr
}
