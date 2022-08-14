package main

import "fmt"

/*
 * @Description:
 * @Author: Liuhongq
 * @Date: 2022-05-31 03:52:22
 * @LastEditTime: 2022-05-31 04:32:23
 * @LastEditors: Liuhongq
 * @Reference:
 */
func selectSort(arr []int) {
	for i := 0; i < len(arr)-1; i++ {
		maxIndex := i
		for j := i + 1; j < len(arr); j++ {
			if arr[j] > arr[maxIndex] {
				maxIndex = j
			}
		}
		if maxIndex != i {
			arr[i], arr[maxIndex] = arr[maxIndex], arr[i]
		}
	}
}
func main() {
	arr := []int{10, 34, 19, 100, 80}
	selectSort(arr)
	fmt.Println(arr)
}
