/*
 * @Descripttion: your project
 * @version: 1.0
 * @Author: power_840
 * @Date: 2021-06-08 21:59:11
 * @LastEditors: power_840
 * @LastEditTime: 2021-06-08 22:23:02
 */
package main

import "fmt"

func main() {
	nums := []int{23, 32, 53, 63, 12, 74, 77, 89, 99}
	Bubble(&nums)
	fmt.Println(nums)
}

// 交换排序
// 1. 冒泡排序 Bubble sorti
// 2. 快速排序 Quick sort

func Bubble(list *[]int) []int {
	for i := 0; i < len(*list)-1; i++ {
		for j := i + 1; j < len(*list)-1-i; j++ {
			if (*list)[i] > (*list)[j] {
				(*list)[i], (*list)[j] = (*list)[j], (*list)[i]
			}
		}

	}
	return *list
}

func Quick() {

}
