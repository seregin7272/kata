package main

import (
	"fmt"
	"log"
)

type Queue struct {
	items []string
}

func (q *Queue) push(item string) {
	q.items = append(q.items, item)
}

func (q *Queue) pop() {
	if q.isEmpty() {
		return
	}
	q.items = q.items[:len(q.items)-1]
}

func (q *Queue) top() string {
	if q.isEmpty() {
		return ""
	}

	return q.items[0]
}

func (q *Queue) isEmpty() bool {
	return len(q.items) == 0
}

func main() {

	queue := Queue{}

	queue.push("1")
	queue.push("2")
	queue.push("3")
	queue.pop()
	fmt.Printf("%#v", queue)
	log.Println(queue.top()) // 1

}
