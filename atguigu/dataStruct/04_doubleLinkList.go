package main

import "fmt"

type Node struct {
	data int
	pre  *Node
	next *Node
}
type List struct {
	head   *Node
	length int
}

func initLinkNode() *List {
	node := new(Node)
	list := new(List)
	list.head = node
	return list
}
func newNode(data int) *Node {
	return &Node{data: data}
}
func (this *List) insertAfter(index int, data int) {
	if index > this.length {
		fmt.Println("位置错误，插入失败")
		return
	}
	p := this.head
	n := newNode(data)
	if index < 0 {
		for {
			if p.next == nil {
				p.next = n
				n.pre = p
				break
			}
			p = p.next
		}
	} else {
		for i := 0; i < index; i++ {
			p = p.next
		}
		if p.next != nil {
			n.next = p.next
			p.next.pre = n
		}
		p.next = n
		n.pre = p
	}

	this.length++
	return
}
func (this *List) del(index int) {
	if index > this.length {
		fmt.Println("位置错误，删除失败")
		return
	}
	if this.length <= 0 {
		fmt.Println("空链表，删除失败")
		return
	}
	p := this.head
	if index < 0 {
		for {
			if p.next == nil {
				p.pre.next = nil
				p.pre = nil
				break
			}
			p = p.next
		}
	} else if index == 0 {
		if p.next.next == nil {
			p.next.pre = nil
			p.next = nil
		} else {
			p.next.next.pre = p
			p.next = p.next.next
		}
	} else {
		for i := 0; i < index; i++ {
			p = p.next
		}
		if p.next != nil {
			p.next.pre = p.pre
			p.pre.next = p.next
		} else {
			p.pre.next = nil
			p.pre = nil
		}
	}

	this.length--
	return
}
func (this *List) show() {
	p := this.head
	for {
		fmt.Printf("%p===>%v\n", p, p)
		if p.next == nil {
			break
		}
		p = p.next
	}
}
func main() {
	list := initLinkNode()
	list.insertAfter(-1, 1)
	list.insertAfter(-1, 2)
	list.insertAfter(-1, 3)
	list.insertAfter(0, 4)
	list.insertAfter(1, 5)
	list.del(0)
	list.del(0)
	list.del(-1)
	list.insertAfter(2, 6)
	list.insertAfter(2, 7)
	list.del(2)
	list.insertAfter(0, 8)
	list.del(0)
	list.del(0)
	list.show()

}
