/*

https://exercism.io/my/solutions/e457f4fface144dfa7b93d90be565702

Write a simple linked list implementation that uses Elements and a List.

The linked list is a fundamental data structure in computer science,
often used in the implementation of other data structures.
They're pervasive in functional programming languages, such as
Clojure, Erlang, or Haskell, but far less common
in imperative languages such as Ruby or Python.

The simplest kind of linked list is a singly linked list.
Each element in the list contains data and a "next"
field pointing to the next element in the list of elements.

This variant of linked lists is often used to represent
sequences or push-down stacks (also called a LIFO stack; Last In, First Out).

As a first take, lets create a singly linked list to contain the range (1..10),
and provide functions to reverse a linked list and convert to and from arrays.

When implementing this in a language with built-in linked lists,
implement your own abstract data type.
*/


package main

import (
    "fmt"
    _"reflect"
    _"math"
)

type Node struct {
    data int
    next *Node
}

type LinkedList struct {
    length int
    head *Node
    tail *Node
}


func (ll *LinkedList) reverse() bool{

    if ll.head != nil {

        current := ll.head.next
        next := ll.head.next.next
        previous := ll.head
        for {

            if next == nil {
                fmt.Println("hey",current.data, current.next, next == nil)
                break
            }

            current.next = previous
            fmt.Println(current.data, current.next)
            previous = current
            current = next
            next = next.next
            fmt.Println(current.data, current.next, next == nil)
        }

        // set the last node's next pointer
        current.next = previous

        // swap the head and the tail of the list
        ll.head, ll.tail = ll.tail, ll.head

        // remove the pointer from the old head
        ll.tail.next = nil
    }
    return true
}

func (q *LinkedList) pop() (item *Node){
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
func (q *LinkedList) push(item int){
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

func (ll *LinkedList) convertToArray() (result []int){
    current := ll.head
    for {
        if current == nil { break }
        result = append(result, current.data)
        current = current.next
    }
    return result
}
//
func convertFrom(slice []int) (result LinkedList){
    for _,v := range(slice){
        result.push(v)
    }
    return
}

func main(){

    var newQueue LinkedList

    newQueue.push(1)
    newQueue.push(2)
    newQueue.push(3)
    newQueue.push(4)

    fmt.Println(newQueue)

    item2 := newQueue.pop()
    fmt.Println(newQueue, item2)
    array := newQueue.convertToArray()
    fmt.Println(array,newQueue, item2)

    var newSlice []int = []int{9,87,6,5,44,3,1,2,3}
    newLL := convertFrom(newSlice)
    fmt.Println(newLL)
    newarray := newLL.convertToArray()
    fmt.Println(newarray)

    ok:=newLL.reverse()
    if ok {
        newarray = newLL.convertToArray()
        fmt.Println(newarray)
    }
}
