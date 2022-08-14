package main

/*
 * @Description:
 * @Author: Liuhongq
 * @Date: 2022-05-31 03:08:31
 * @LastEditTime: 2022-05-31 03:09:07
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
	if (this.tail+1)%this.maxSize == this.head {
		fmt.Println("Queue full")
		return
	}

	this.arr[this.tail] = data
	this.tail = (this.tail + 1) % this.maxSize
	return
}

func (this *Queue) pop() (val int) {
	if this.head == this.tail {
		fmt.Println("Queue null")
		return
	}
	val = this.arr[this.head]
	this.head = (this.head + 1) % this.maxSize
	return
}
func (this *Queue) show() {
	temp := this.head
	for i := 0; i < (this.tail-this.head+this.maxSize)%this.maxSize; i++ {
		fmt.Printf("arr[%d]=%d\n", temp, this.arr[temp])
		temp = (temp + 1) % this.maxSize
	}
	fmt.Println()
}
func main() {
	queue := &Queue{
		maxSize: 5,
		head:    0,
		tail:    0,
	}

	for i := 0; i < 4; i++ {
		queue.push(i)
	}
	queue.show()
	queue.pop()

	queue.pop()
	queue.show()
	queue.push(5)
	queue.push(6)
	queue.push(7)
	queue.show()
	queue.pop()
	queue.push(8)
	queue.show()
}
