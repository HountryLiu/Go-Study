package main

/*
 * @Description:
 * @Author: Liuhongq
 * @Date: 2022-05-31 06:14:26
 * @LastEditTime: 2022-05-31 07:32:24
 * @LastEditors: Liuhongq
 * @Reference:
 */

import "fmt"

func quickSort(arr []int) []int {
	if len(arr) < 2 {
		return arr
	}
	pindex := (len(arr) - 1) / 2
	pivot := arr[pindex]
	var left, right, mid []int
	for _, v := range arr {
		if v < pivot {
			left = append(left, v)
		} else if v == pivot {
			mid = append(mid, v)
		} else {
			right = append(right, v)
		}
	}
	var res []int
	res = append(res, quickSort(left)...)
	res = append(res, mid...)
	res = append(res, quickSort(right)...)
	return res
}
func main() {
	arr := []int{10, 34, 19, 100, 80, -1}
	fmt.Println(arr)
	arr = quickSort(arr)
	fmt.Println(arr)

}
