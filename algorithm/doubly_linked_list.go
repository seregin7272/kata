package main

import (
	"fmt"
)

type DoublyListNode struct {
	Val  int
	Next *DoublyListNode
	Prev *DoublyListNode
}

type DoublyLinkedList struct {
	Head *DoublyListNode
	Tail *DoublyListNode
}

func (dll *DoublyLinkedList) PushFront(val int) {
	node := &DoublyListNode{Val: val}
	if dll.Head == nil {
		dll.Head = node
		dll.Tail = node
		return
	}
	node.Next = dll.Head
	dll.Head.Prev = node
	dll.Head = node
}

func (dll *DoublyLinkedList) PopFront() {
	if dll.Head == nil {
		return
	}

	head := dll.Head.Next

	if head == nil {
		dll.Head = nil
		dll.Tail = nil
		return
	}

	head.Prev = nil
	dll.Head = head
}

func (dll *DoublyLinkedList) PushBack(val int) {
	node := &DoublyListNode{Val: val}
	if dll.Head == nil {
		dll.Head = node
		dll.Tail = node
		return
	}

	dll.Tail.Next = node
	node.Prev = dll.Tail
	dll.Tail = node

}

func (dll *DoublyLinkedList) PopBack() {
	if dll.Tail == nil {
		return
	}

	dll.Tail = dll.Tail.Prev

	if dll.Tail == nil {
		dll.Head = nil
		return
	}

	dll.Tail.Next = nil
}

func (dll DoublyLinkedList) Get(idx int) int {
	head := dll.Head
	for i := 0; i < idx; i++ {
		if head == nil {
			return -1
		}

		head = head.Next
	}

	if head == nil {
		return -1
	}

	return head.Val
}

func (dll *DoublyLinkedList) Set(idx int, val int) {
	node := &DoublyListNode{Val: val}

	head := dll.Head
	if head == nil {
		return
	}

	for i := 0; i < idx; i++ {
		if head == nil {
			return
		}

		head = head.Next
	}

	if head == nil {
		return
	}

	prev := head.Prev
	prev.Next = node
	node.Prev = prev
	node.Next = head
	head.Prev = node

	if node.Prev == nil {
		dll.Head = node
	}

	if node.Next == nil {
		dll.Tail = node
	}

}

// Replace
//Ваша функция должна «вырезать» участок с l по r индексы (l — включительно, r — нет) из листа
//и вставить после i-ого элемента (номера элементов считаются до вырезания, начиная с 0).
//Если хотя бы один из индексов выходит за пределы листа, завершите метод не изменяя список.
// l = 2, r = 4 i = 6
// in [0, 1, 2, 3, 4, 5, 6, 7]
// out [0, 1, 4, 5, 6, 2, 3, 7]
func (dll *DoublyLinkedList) Replace(l int, r int, idx int) {
	head := dll.Head
	start := &DoublyListNode{}
	end := &DoublyListNode{}
	if head == nil {
		return
	}

	i := 0
	for head != nil {
		if i == l {
			start = head
		}

		if i == r {
			end = head
		}

		if i == idx {
			break
		}

		head = head.Next
		i++
	}

	endPrev := end.Prev
	// вырезали l по r
	end.Prev = start.Prev
	start.Prev.Next = end

	// переставили
	headNext := head.Next
	start.Prev = head
	head.Next = start
	headNext.Prev = endPrev
	endPrev.Next = headNext
	fmt.Println()

}

func main() {
	dll := &DoublyLinkedList{}

	for i := 0; i < 8; i++ {
		dll.PushBack(i)
	}

	fmt.Println(dll.Get(2))
	dll.Replace(4, 5, 2)
	fmt.Println(dll.Get(2))
}
