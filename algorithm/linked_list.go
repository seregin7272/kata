package main

import (
	"fmt"
	"log"
)

type ListNode struct {
	Val  int
	Next *ListNode
}

type LinkedList struct {
	Head *ListNode
	Len  int
}

func (l *LinkedList) Insert(pos int, val int) {
	newNode := &ListNode{Val: val}
	if pos < 0 || pos > l.Len {
		return
	}
	if pos == 0 {
		newNode.Next = l.Head
		l.Head = newNode
		l.Len++
		return
	}

	prevNode := l.IndexOf(pos - 1)
	nextNode := prevNode.Next
	newNode.Next = nextNode
	prevNode.Next = newNode

	l.Len++
}

func (l *LinkedList) Remove(pos int) {
	prevNode := l.IndexOf(pos - 1)
	prevNode.Next = prevNode.Next.Next
	l.Len--
}

func (l *LinkedList) IndexOf(pos int) *ListNode {
	node := l.Head
	for i := 0; i < pos; i++ {
		node = node.Next
	}

	return node
}

func (l *LinkedList) PushFront(val int) {
	if l.Head == nil {
		l.Head = &ListNode{Val: val}
		return
	}
	l.Head.Next = &ListNode{Val: val}
}

func (l *LinkedList) PushBack(val int) {
	if l.Head == nil {
		l.Head = &ListNode{Val: val}
		return
	}

	node := l.Head
	for node.Next != nil {
		node = node.Next
	}

	node.Next = &ListNode{Val: val}
}

func (l *LinkedList) ToString() {
	node := l.Head
	for node.Next != nil {
		log.Println(node.Val)
		node = node.Next
	}

	log.Println(node.Val)
}

// https://leetcode.com/problems/linked-list-cycle/
func hasCycle(head *ListNode) bool {
	fast := head
	slow := head

	for {
		if fast == nil || fast.Next == nil {
			return false
		}
		fast = fast.Next.Next
		slow = slow.Next

		if fast == slow {
			return true
		}
	}

	return false
}

// https://leetcode.com/problems/delete-node-in-a-linked-list/
func deleteNode(node *ListNode) {
	node.Val, node.Next = node.Next.Val, node.Next.Next
}

// https://leetcode.com/problems/remove-linked-list-elements/
func removeElements(head *ListNode, val int) *ListNode {
	if head == nil {
		return head
	}

	tmp := head

	for {

		if head.Val != val {
			break
		}

		if head.Next == nil {
			return nil
		}

		head = head.Next
	}

	for {
		if tmp.Next == nil {
			break
		}

		if tmp.Next.Val == val {
			tmp.Next = tmp.Next.Next
		} else {
			tmp = tmp.Next
		}
	}

	return head
}

// https://leetcode.com/problems/merge-two-sorted-lists/
func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
	head := &ListNode{Val: -1} // фиктивный узел для хранения результата
	tail := head               // хранит последний узел(наибольший) в хвосте head или это временый хвост head
	for {

		// Если какой-либо список становится полностью пустым
		// напрямую соединяем все элементы другого списка
		if list1 == nil {
			tail.Next = list2
			break
		} else if list2 == nil {
			tail.Next = list1
			break
		}

		// Добавлен к последнему из отсортировнного списка через
		// tail.Next т.к. tail - последний отсортированный элемент

		if list1.Val < list2.Val {
			tail.Next = list1 // tail.Next = list1 = head ... node.Next
			list1 = list1.Next

		} else {
			tail.Next = list2 // dummy.Next ссмотрит на list2
			list2 = list2.Next
		}

		tail = tail.Next //  двигаем хвост
	}

	return head.Next
}

// https://leetcode.com/problems/reverse-linked-list/
// Input: head = [1,2,3,4,5]
//Output: [5,4,3,2,1]
func reverseList(head *ListNode) *ListNode {
	var prev *ListNode

	for head != nil {
		next := head.Next
		head.Next = prev

		prev = head
		head = next
	}

	return prev
}

// https://leetcode.com/problems/palindrome-linked-list/
//Input: head = [1,2,2,1]
//Output: true
func isPalindrome(head *ListNode) bool {
	curr := head
	var stack []int

	for curr != nil {
		stack = append(stack, curr.Val)
		curr = curr.Next
	}

	curr = head
	for curr != nil {
		size := len(stack)
		val := stack[size-1]
		if val != curr.Val {
			return false
		}

		stack = stack[:size-1]
		curr = curr.Next

	}

	return true
}

//https://leetcode.com/problems/swapping-nodes-in-a-linked-list/
//Input: head = [1,2,3,4,5], k = 2
//Output: [1,4,3,2,5]
func swapNodes(head *ListNode, k int) *ListNode {
	curr := head
	var list []int
	for curr != nil {
		list = append(list, curr.Val)
		curr = curr.Next
	}

	list[k-1], list[len(list)-k] = list[len(list)-k], list[k-1]

	curr = head
	i := 0
	for curr != nil {
		curr.Val = list[i]
		curr = curr.Next
		i++
	}

	return head
}

func main() {

	ll1 := LinkedList{}
	ll1.PushBack(1)
	ll1.PushBack(2)
	ll1.PushBack(3)
	ll1.PushBack(4)
	ll1.PushBack(5)

	list := swapNodes(ll1.Head, 2)

	fmt.Println(list)

}
