package main

import "fmt"

type Node struct {
	data int
	prev *Node
	next *Node
}

type LinkedList struct {
	head *Node
	tail *Node
}

func (list *LinkedList) AddNode(data int) {
	newNode := &Node{
		data: data,
		prev: nil,
		next: nil,
	}

	if list.head == nil {
		list.head = newNode
		list.tail = newNode
	} else {
		newNode.prev = list.tail
		list.tail.next = newNode
		list.tail = newNode
	}
}

func (list *LinkedList) Print() {
	currentNode := list.head
	for currentNode != nil {
		fmt.Printf("%d ", currentNode.data)
		currentNode = currentNode.next
	}
	fmt.Println()
}

func (list *LinkedList) Reverse() {
	currentNode := list.tail
	for currentNode != nil {
		fmt.Printf("%d ", currentNode.data)
		currentNode = currentNode.prev
	}
	fmt.Println()
}

func main() {
	list := LinkedList{}

	list.AddNode(10)
	list.AddNode(20)
	list.AddNode(30)
	list.AddNode(40)

	fmt.Println("Double Linked List (forward)")
	list.Print()

	fmt.Println("Double Linked List (reverse)")
	list.Reverse()
}
