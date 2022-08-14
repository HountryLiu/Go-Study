package main

import "fmt"

type Boy struct {
	no   int
	next *Boy
}

func initGame(num int) *Boy {
	head := &Boy{}
	tail := &Boy{}
	if num < 1 {
		fmt.Println("数字必须大于等于1")
		return head
	}
	for i := 1; i <= num; i++ {
		newBoy := &Boy{no: i}
		if i == 1 {
			head = newBoy
			tail = head
			head.next = head //
		} else {
			tail.next = newBoy
			newBoy.next = head
		}
		tail = tail.next
	}
	return head
}

func show(head *Boy) {
	if head.next == nil {
		fmt.Println("无人")
		return
	}
	tail := head
	for {
		fmt.Printf("Boy[%d] ===> ", tail.no)
		if tail.next == head {
			break
		}
		tail = tail.next
	}
}
func countBoy(head *Boy) int {
	if head.next == nil {
		return 0
	}
	count := 0
	tail := head
	for {
		count++
		if tail.next == head {
			break
		}
		tail = tail.next
	}
	return count
}

/*
设编号为1，2，… n的n个人围坐一圈，约定编号为k（1<=k<=n）
的人从1开始报数，数到m 的那个人出列，它的下一位又从1开始报数，
数到m的那个人又出列，依次类推，直到所有人出列为止，由此产生一个出队编号的序列
*/
func playGame(head *Boy, k int, m int) {
	if head.next == nil {
		fmt.Println("无人")
		return
	}
	if k < 1 || k > countBoy(head) {
		fmt.Println("数字错误")
		return
	}
	tail := head
	//tail移到队尾
	for {
		if tail.next == head {
			break
		}
		tail = tail.next
	}
	//从k开始报数,自己会数一次所以k-1
	for i := 1; i <= k-1; i++ {
		head = head.next
		tail = tail.next
	}
	//循环数，数到m 的那个人出列，自己会数一次所以m-1
	for {
		//只剩一个人推出
		if head == tail {
			break
		}
		for i := 1; i <= m-1; i++ {
			head = head.next
			tail = tail.next
		}
		fmt.Printf("Boy[%d]出队列\n", head.no)
		//删除head结点人
		head = head.next
		tail.next = head
	}
	fmt.Printf("Boy[%d]最后出队列\n", head.no)
}
func main() {
	head := initGame(5)
	show(head)
	fmt.Println()
	playGame(head, 2, 3)
}
