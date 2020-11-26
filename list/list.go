package main

import "fmt"

type element struct {
	value string
	next  *element
}

type list struct {
	head *element
	tail *element
}

func (list *list) add(newvalues string) {
	elm := element{newvalues, nil}

	if list.head == nil {
		list.head = &elm
	} else {
		elm.isnextto(list.tail)
	}
	list.tail = &elm
}

func (elm *element) isnextto(prevelm *element) {
	prevelm.next = elm
}

func (elm *element) getValue() string {
	return elm.value
}

func (list *list) listall() {
	elm := list.head
	for {
		fmt.Println(elm.getValue())
		if elm.next == nil {
			break
		}
		elm = elm.next
	}
}

func main() {
	myList := list{nil, nil}
	myList.add("pig")
	myList.add("bird")
	myList.add("dog")
	myList.add("oh")
	myList.listall()
}
