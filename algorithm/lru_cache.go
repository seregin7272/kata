package main

import (
	"fmt"
	"sync"
)

// Наименее недавно использованный (LRU)
//Input
//["LRUCache", "put", "put", "get", "put", "get", "put", "get", "get", "get"]
//[[2], [1, 1], [2, 2], [1], [3, 3], [2], [4, 4], [1], [3], [4]]
//Output
//[null, null, null, 1, null, -1, null, -1, 3, 4]
//
//Explanation
//LRUCache lRUCache = new LRUCache(2);
//lRUCache.put(1, 1); // cache is {1=1}
//lRUCache.put(2, 2); // cache is {1=1, 2=2}
//lRUCache.get(1);    // return 1
//lRUCache.put(3, 3); // LRU key was 2, evicts key 2, cache is {1=1, 3=3}

//lRUCache.get(2);    // returns -1 (not found)
//lRUCache.put(4, 4); // LRU key was 1, evicts key 1, cache is {4=4, 3=3}
//lRUCache.get(1);    // return -1 (not found)
//lRUCache.get(3);    // return 3
//lRUCache.get(4);    // return 4

type LRUList struct {
	Head *LRUNode
	Tail *LRUNode
}

func (dll *LRUList) Push(key int, value int) {
	node := &LRUNode{Val: value, Key: key}
	if dll.Head == nil {
		dll.Tail = node
		dll.Head = node
		return
	}

	dll.Head.Prev = node
	node.Next = dll.Head
	dll.Head = node
}

func (dll *LRUList) Pop() {
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

func (dll *LRUList) MoveToHead(node *LRUNode) {
	if node == dll.Head {
		return
	}

	prev := node.Prev
	next := node.Next

	if next != nil {
		next.Prev = prev
	}
	prev.Next = next

	if prev.Next == nil {
		dll.Tail = prev
	}

	dll.Head.Prev = node
	node.Next = dll.Head
	dll.Head = node
	dll.Head.Prev = nil
}

type LRUNode struct {
	Next *LRUNode
	Prev *LRUNode
	Key  int
	Val  int
}

type LRUCache struct {
	capacity int
	data     map[int]*LRUNode
	dll      LRUList
	sync.RWMutex
}

func Constructor(capacity int) LRUCache {
	return LRUCache{
		capacity: capacity,
		data:     make(map[int]*LRUNode, 2),
		dll:      LRUList{},
	}
}

func (lru *LRUCache) Get(key int) int {
	node, ok := lru.data[key]

	if !ok {
		return -1
	}

	lru.dll.MoveToHead(node)

	return node.Val
}

func (lru *LRUCache) Put(key int, value int) {
	lru.Lock()
	defer lru.Unlock()
	if node, ok := lru.data[key]; ok {

		node.Val = value
		lru.dll.MoveToHead(node)
		return
	}

	if len(lru.data) == lru.capacity {
		node := lru.dll.Tail
		delete(lru.data, node.Key)
		lru.dll.Pop()
	}

	lru.dll.Push(key, value)
	lru.data[key] = lru.dll.Head
}

func main() {

	lru := Constructor(2)

	fmt.Println(lru.Get(2)) // -1
	lru.Put(2, 6)           // {2,6}
	fmt.Println(lru.Get(1)) // -1
	lru.Put(1, 5)           // [{1,5} {2,6}]
	lru.Put(1, 2)           // [{1,2} {2,6}]
	fmt.Println(lru.Get(1)) // 2
	lru.Get(2)              //6

}
