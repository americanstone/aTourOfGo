package main

import (
	"testing"
)

func TestInsert(t *testing.T) {
	l := NewSinglyLinkedList[int]()
	l.InsertAtHead(1)
	l.InsertAtEnd(2)
	l.InsertAt(0, 0)
	l.InsertAt(3, 2)
	l.InsertAt(3, 5)
	if l.Size() != 4 {
		t.Fatalf("size should be 4 but %v", l.Size())
	}

	for i := uint(0); i < 4; i++ {
		v, ok := l.Get(i)
		if !ok {
			t.Fatalf("should be to get value with valid index %v ", i)
		}
		if uint(v) != i {
			t.Fatalf("index %v should be 0 but %v", i, v)
		}
	}
}

func TestDelete(t *testing.T) {
	l := NewSinglyLinkedList[int]()
	l.InsertAtHead(1)
	l.InsertAtEnd(2)
	l.Append(3)
	l.Append(4)
	l.Append(5)

	if l.Size() != 5 {
		t.Fatalf("size should be 5 but %v", l.Size())
	}

	if l.DeleteHead() != true {
		t.Fatal("should be able to delete head")
	}

	if l.head.val != 2 {
		t.Fatal("new head should be 2")

	}
	l.DeleteAt(1)
	v, _ := l.Get(1)
	if v != 4 {
		t.Fatal("value should be 4")

	}

}
