package main

import (
	"testing"
)

func TestLinkedListDeleteNode(t *testing.T) {
	list := &LinkedList{}

	list.Append(1)
	list.Append(33)
	list.Append(2)

	list.DeleteNode(0)

	if list.Head.Data != 33 {
		t.Errorf("linkedlist deletenode err")
	}

	if list.Head.Next.Data != 2 {
		t.Errorf("linkedlist deletenode err")
	}

	if list.Head.Next.Next != nil {
		t.Errorf("linkedlist deletenode err")
	}
}

func TestLinkedListDeleteNodeEmpty(t *testing.T) {
	list := &LinkedList{}

	list.DeleteNode(0)

	if list.Head != nil {
		t.Errorf("linkedlist deletenode err")
	}
}


func TestLinkedListUpdate(t *testing.T) {
	list := &LinkedList{}

	list.Append(1)
	list.Append(33)
	list.Append(2)

	list.Update(0, 11)

	if list.Head.Data != 11 {
		t.Errorf("linkedlist deletenode err")
	}
}

func TestLinkedListAppend(t *testing.T) {
	list := &LinkedList{}

	list.Append(101)
	list.Append(102)

	if list.Head.Data != 101 {
		t.Errorf("linkedlist deletenode err")
	}

	if list.Head.Next.Data != 102 {
		t.Errorf("linkedlist deletenode err")
	}

	if list.Head.Next.Next != nil {
		t.Errorf("Append: expected nil and found %v", list.Head.Next.Next)
	}
}

func TestLinkedListDeleteValue(t *testing.T) {
	list := &LinkedList{}

	list.Append(1)
	list.Append(33)
	list.Append(33)

	list.DeleteValue(33)

	if list.Head.Data != 1 {
		t.Errorf("linkedlist deletenode err")
	}

	if list.Head.Next.Data != 33 {
		t.Errorf("linkedlist deletenode err")
	}

	if list.Head.Next.Next != nil {
		t.Errorf("Delete Value: expected nil and found %v", list.Head.Next.Next)
	}

}
