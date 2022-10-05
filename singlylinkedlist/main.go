package main

import (
	"fmt"
	//package name and folder name are different
	singlylinkedlist "example.com/linkedlist"
	"example.com/list"
	"github.com/americanstone/singlylist"
)

// testing different ways to import a package

func main() {
	remote := singlylist.NewSinglyLinkedList[uint]()
	remote.InsertAtEnd(2)
	test := list.NewSinglyLinkedList[uint]()
	test.InsertAt(0, 1)
	// this because the module
	l := singlylinkedlist.NewSinglyLinkedList[uint]()
	l.InsertAtHead(1)
	l.InsertAtHead(2)
	l.InsertAtHead(3)
	l.InsertAtEnd(4)
	fmt.Println("size: ", l.Size())
	l.Print()
	v, _ := l.Get(1)

	fmt.Println("get 1: ", v)

	l.DeleteAt(0)
	fmt.Println("size: ", l.Size())
	l.Print()
	l.DeleteAt(2)
	fmt.Println("size: ", l.Size())
	l.Print()
	l.DeleteAt(1)
	fmt.Println("size: ", l.Size())
	l.Print()
	l.DeleteVal(2)
	fmt.Println("size: ", l.Size())
	l.Print()
	l.InsertAtHead(1)
	l.InsertAtHead(2)
	l.InsertAtHead(3)
	l.InsertAtHead(1)
	l.InsertAtHead(2)
	l.InsertAtHead(3)
	l.Print()
	l.DeleteVal(2)
	fmt.Println("delete value 2")
	fmt.Println("size: ", l.Size())
	l.Print()
	l.DeleteVal(1)
	fmt.Println("delete value 1")
	fmt.Println("size: ", l.Size())
	l.Print()
	l.DeleteVal(3)
	fmt.Println("delete value 3")
	fmt.Println("size: ", l.Size())
	l.Print()

	l.DeleteVal(2)
	fmt.Println("delete value 3")
	fmt.Println("size: ", l.Size())
	l.Print()

}
