package main

import "fmt"

type Node struct {
	Data int   // Data stored in the node (can be any type)
	Next *Node // Pointer to the next node
}

type LinkedList struct {
	Head *Node // Pointer to the first node (head)
}

func (ll *LinkedList) DeleteNode(index int) {
	current := ll.Head

	if current == nil {
		return
	}

	if index == 0 {
		ll.Head = ll.Head.Next
		return
	}

	counter := 1
	current = ll.Head.Next
	for current.Next != nil {
		if counter == index-1 {
			fmt.Println(fmt.Sprintf("me pongo a cambiar puntero en index %d", index))
			current.Next = current.Next.Next
			return
		}
		counter++
		current = current.Next
	}
}

func (ll *LinkedList) Update(index int, value int) {
	current := ll.Head
	counter := 0

	if current == nil {
		return
	}

	if index == 0 {
		ll.Head.Data = value
		return
	}
	counter++
	for current.Next != nil {
		current = current.Next
		if counter == index {
			current.Data = value
			return
		}
		counter++
	}
}

func (ll *LinkedList) Append(data int) {
	newNode := &Node{Data: data, Next: nil}

	if ll.Head == nil {
		ll.Head = newNode
	} else {
		current := ll.Head
		for current.Next != nil {
			current = current.Next
		}
		current.Next = newNode
	}
}

func (ll *LinkedList) Display() {
	current := ll.Head
	for current != nil {
		fmt.Printf("%d -> ", current.Data)
		current = current.Next
	}
	fmt.Println("nil")
}

func (ll *LinkedList) DeleteValue(data int) {
	if ll.Head == nil {
		return
	}

	if ll.Head.Data == data {
		ll.Head = ll.Head.Next
		return
	}

	current := ll.Head
	for current.Next != nil {
		if current.Next.Data == data {
			current.Next = current.Next.Next
			return
		}
		current = current.Next
	}
}
