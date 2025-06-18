package main

import "fmt"

type Node struct {
	data int
	next *Node
}

type LinkedList struct {
	head *Node
}

func (list *LinkedList) Insert(data int) {
	newNode := &Node{data: data}

	if list.head == nil {
		list.head = newNode
	} else {
		current := list.head
		for current.next != nil {
			current = current.next
		}
		current.next = newNode
	}
}

func (list *LinkedList) Display() {
	current := list.head

	if current == nil {
		fmt.Println("Linked List is empty")
		return
	}

	fmt.Print("Linked List: ")
	for current != nil {
		fmt.Printf("%d ", current.data)
		current = current.next
	}
	fmt.Println()
}

func main() {
	list := LinkedList{}

	list.Insert(10)
	list.Insert(20)
	list.Insert(30)
	list.Insert(40)

	list.Display()
}
