package main

import (
	"fmt"
	"strings"
    "math/rand"
)

type Node[T any] struct {
    data  T
	next *Node[T]
}

type DiyList[T any] struct {
    head *Node[T]
}

func (l *DiyList[T]) Append(val T) {
    newNode := &Node[T]{data: val, next: nil}

    if l.head == nil {
        l.head = newNode
        return
    }

    ptr := l.head
    for ; ptr.next != nil; ptr = ptr.next {}

    ptr.next = newNode
}

func (l DiyList[T]) String() (output string) {
    for ptr := l.head; ptr != nil; ptr = ptr.next {
        output += fmt.Sprintf("%v ", ptr.data)
    }

    if len(output) == 0 {
        return "<Empty DiyList>"
    }

    return strings.TrimSpace(output)
}

func main() {
    myList := DiyList[int]{}
    for i := 1; i <= 5; i++ {
        myList.Append(i << (rand.Int() % 32))
    }
    fmt.Println(myList)

    emptyList := DiyList[float32]{}

    fmt.Println(emptyList)
}