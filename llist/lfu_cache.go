package llist

import "sync"

// Разработайте и внедрите структуру данных для наименее часто используемого кэша (LFU)
//Input
//["LFUCache", "put", "put", "get", "put", "get", "get", "put", "get", "get", "get"]
//[[2], [1, 1], [2, 2], [1], [3, 3], [2], [3], [4, 4], [1], [3], [4]]
//Output
//[null, null, null, 1, null, -1, 3, null, -1, 3, 4]
//
//Explanation
//// cnt(x) = the use counter for key x
//// cache=[] will show the last used order for tiebreakers (leftmost element is  most recent)
//LFUCache lfu = new LFUCache(2);
//lfu.put(1, 1);   // cache=[1,_], cnt(1)=1
//lfu.put(2, 2);   // cache=[2,1], cnt(2)=1, cnt(1)=1
//lfu.get(1);      // return 1
//// cache=[1,2], cnt(2)=1, cnt(1)=2
//lfu.put(3, 3);   // 2 is the LFU key because cnt(2)=1 is the smallest, invalidate 2.
//// cache=[3,1], cnt(3)=1, cnt(1)=2
//lfu.get(2);      // return -1 (not found)
//lfu.get(3);      // return 3
//// cache=[3,1], cnt(3)=2, cnt(1)=2
//lfu.put(4, 4);   // Both 1 and 3 have the same cnt, but 1 is LRU, invalidate 1.
//// cache=[4,3], cnt(4)=1, cnt(3)=2
//lfu.get(1);      // return -1 (not found)
//lfu.get(3);      // return 3
//// cache=[3,4], cnt(4)=1, cnt(3)=3
//lfu.get(4);      // return 4
//// cache=[4,3], cnt(4)=2, cnt(3)=3

type DLListNode struct {
	Next *DLListNode
	Prev *DLListNode
	Key  int
	Val  int
	Cnt  int
}

type DLList struct {
	Head *DLListNode
	Tail *DLListNode
}

func (dll *DLList) Push(key int, value int) {
	node := &DLListNode{Val: value, Key: key, Cnt: 1}
	if dll.Head == nil {
		dll.Tail = node
		dll.Head = node
		return
	}

	dll.Head.Prev = node
	node.Next = dll.Head
	dll.Head = node
}

func (dll *DLList) Pop() {
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

// Перемещает ноду из одного листа в голову другого
func (dll *DLList) Move(node *DLListNode, listFrom *DLList, listTo *DLList) {
	// Удаляем из старого листа
	next := node.Next
	prev := node.Prev

	if prev == nil {
		listFrom.Head = next
	} else {
		prev.Next = next
	}

	if next == nil {
		listFrom.Tail = prev
	} else {
		next.Prev = prev
	}

	// Добавляем в новый
	node.Next = nil
	node.Prev = nil

	if listTo.Head == nil {
		listTo.Head = node
		listTo.Tail = node
		return
	}

	listTo.Head.Prev = node
	node.Next = listTo.Head
	listTo.Head = node
}

type LFUCache struct {
	capacity int
	data     map[int]*DLListNode
	freq     map[int]*DLList
	min      int
	sync.RWMutex
}

func Constructor(capacity int) LFUCache {
	return LFUCache{
		capacity: capacity,
		data:     map[int]*DLListNode{},
		freq:     map[int]*DLList{},
	}
}

func (lfu *LFUCache) runCache(node *DLListNode) {

	// переместить ноду в другой бакет +1
	cnt := node.Cnt
	node.Cnt += 1
	dll, ok := lfu.freq[node.Cnt]
	if !ok {
		dll = &DLList{}
		lfu.freq[node.Cnt] = dll
	}

	dll.Move(node, lfu.freq[cnt], dll)

	// удалить если пустой бакет
	dll = lfu.freq[cnt]
	if dll.Head == nil {
		delete(lfu.freq, cnt)
		if lfu.min == cnt {
			lfu.min += 1
		}
	}
}

func (lfu *LFUCache) Get(key int) int {
	node, ok := lfu.data[key]

	if !ok {
		return -1
	}

	lfu.runCache(node)

	return node.Val
}

func (lfu *LFUCache) Put(key int, value int) {
	lfu.Lock()
	defer lfu.Unlock()

	if lfu.capacity <= 0 {
		return
	}

	// если есть такой ключ , +1 и перекидваем в начало списка
	if node, ok := lfu.data[key]; ok {
		node.Val = value
		lfu.runCache(node)

		return
	}

	if lfu.capacity == len(lfu.data) {
		dll := lfu.freq[lfu.min]
		delete(lfu.data, dll.Tail.Key)
		dll.Pop()
		if dll.Head == nil {
			delete(lfu.freq, lfu.min)
		}
	}

	lfu.min = 1
	dll, ok := lfu.freq[lfu.min]
	if !ok {
		dll = &DLList{}
		lfu.freq[lfu.min] = dll
	}

	dll.Push(key, value)
	lfu.data[key] = dll.Head
}
