package main

/*
 * @Description:
 * @Author: Liuhongq
 * @Date: 2022-05-31 05:00:37
 * @LastEditTime: 2022-05-31 05:14:21
 * @LastEditors: Liuhongq
 * @Reference:
 */

import "fmt"

func insertSort(arr []int) {
	for i := 1; i <= len(arr)-1; i++ {
		curVal := arr[i]
		compareIndex := i - 1
		for compareIndex >= 0 && curVal > arr[compareIndex] {
			arr[compareIndex+1] = arr[compareIndex]
			compareIndex--
		}
		if compareIndex+1 != i {
			arr[compareIndex+1] = curVal
		}
	}
}
func main() {
	arr := []int{10, 34, 19, 100, 80}
	insertSort(arr)
	fmt.Println(arr)
}
