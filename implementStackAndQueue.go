package main

import (
    "fmt"
    )
//  STACK
type Stack struct {
    length int
    items []int
}
func (s *Stack) pop() (item int){
    if s.length > 0 {
        item = s.items[s.length-1]
        s.items = s.items[:s.length-1]
        s.length -= 1
    } else {
        item = 0 //panic("My first panic!")
    }
    return
}
func (s *Stack) push(item int){
    s.items = append(s.items, item)
    s.length+=1
    return
}

//  QUEUE
type Queue struct {
    head *Node
    tail *Node
    length int
}
type Node struct {
    data int
    next *Node
}
func (q *Queue) pop() (item *Node){
    if q.length > 0 {
        item = q.head
        if q.head.next != nil {
            q.head = q.head.next
        }
        item.next = nil
        q.length -= 1
    }
    return
}
func (q *Queue) push(item int){
    newNode := Node{data:item}
    if q.length > 0 { // if there are items in the queue
        if q.head != q.tail { // if there is more than one item in the queue
            q.tail.next = &newNode
            q.tail = &newNode
        } else { // if there's only one item in the queue
            q.head.next = &newNode
            q.tail = &newNode
        }
    } else { // if its the first item in the queue
        q.head = &newNode
        q.tail = &newNode
    }
    q.length+=1
    return
}

// how would you implement a stack and queue in Go?
func main() {

    var newStack Stack
    newStack.items = []int{1,2,3,4,5,5,6}
    newStack.length = len(newStack.items)

    item1 := newStack.pop()
    fmt.Println(item1, newStack)
    newStack.push(1)
    fmt.Println(newStack)


    var newQueue Queue

    newQueue.push(1)
    newQueue.push(2)
    newQueue.push(3)
    newQueue.push(4)

    fmt.Println(newQueue)

    item2 := newQueue.pop()
    fmt.Println(newQueue, item2)

}
