/*
 * @Descripttion: your project
 * @version: 1.0
 * @Author: power_840
 * @Date: 2021-07-08 21:12:35
 * @LastEditors: power_840
 * @LastEditTime: 2021-07-08 21:50:11
 */

package main

import "fmt"

func main() {
	arr := []int{1, 2, 3, 4, 5, 6}
	findFirst(arr)
}

func findFirst(arr []int) {
	var entry int
	fmt.Println("input name: ...")
	fmt.Scanln(&entry)

	for i := 0; i < len(arr); i++ {
		if entry == arr[i] {
			fmt.Printf("找到%v, 下标%v", entry, i)
			break
		} else if i == len(arr)-1 {
			fmt.Printf("没有找到%v", entry)
		}
	}
}

func findSecond(arr []int) int {
	index := -1
	var entry int
	fmt.Println("input name: ...")
	fmt.Scanln(&entry)

	for i := 0; i < len(arr); i++ {
		if entry == arr[i] {
			index = i
			break
		}
	}
	return index
}
