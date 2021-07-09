/*
 * @Descripttion: your project
 * @version: 1.0
 * @Author: power_840
 * @Date: 2021-07-09 22:53:33
 * @LastEditors: power_840
 * @LastEditTime: 2021-07-09 22:59:55
 */
package main

import "fmt"

func main() {
	var multiArr [4][6]int
	multiArr[1][2] = 1
	multiArr[2][1] = 2
	multiArr[2][3] = 3
	// fmt.Println(multiArr)
	for _, v := range multiArr {
		for _, j := range v {
			fmt.Print(j)
		}
		fmt.Println()
	}
}
