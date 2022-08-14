package main

import "fmt"

/*
 * @Description:
 * @Author: Liuhongq
 * @Date: 2022-05-31 03:52:22
 * @LastEditTime: 2022-05-31 04:05:56
 * @LastEditors: Liuhongq
 * @Reference:
 */
func bubbleSort(arr []int) {
	for i := 0; i < len(arr)-1; i++ {
		for j := 0; j < len(arr)-1-i; j++ {
			if arr[j] > arr[j+1] {
				arr[j], arr[j+1] = arr[j+1], arr[j]
			}
		}
	}
}
func main() {
	arr := []int{10, 34, 19, 100, 80}
	bubbleSort(arr)
	fmt.Println(arr)
}
