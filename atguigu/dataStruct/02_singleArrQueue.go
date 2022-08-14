package main

/*
 * @Description:
 * @Author: Liuhongq
 * @Date: 2022-05-31 03:07:04
 * @LastEditTime: 2022-05-31 03:08:15
 * @LastEditors: Liuhongq
 * @Reference:
 */

import (
	"fmt"
)

type Queue struct {
	maxSize int
	arr     [5]int
	head    int
	tail    int
}

func (this *Queue) push(data int) {
	if this.tail == this.maxSize-1 {
		fmt.Println("Queue full")
		return
	}
	this.tail++
	this.arr[this.tail] = data
	return
}

func (this *Queue) pop() (val int) {
	if this.head == this.tail {
		fmt.Println("Queue null")
		return
	}
	this.head++
	val = this.arr[this.head]
	return
}
func (this *Queue) show() {
	for i := this.head + 1; i <= this.tail; i++ {
		fmt.Printf("arr[%d]=%d\n", i, this.arr[i])
	}
}
func main() {
	queue := &Queue{
		maxSize: 5,
		head:    -1,
		tail:    -1,
	}

	for i := 0; i < 5; i++ {
		queue.push(i)
	}

	queue.pop()
	queue.pop()
	queue.push(5)
	queue.show()
}
