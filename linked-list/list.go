package main

import "fmt"

type List interface {
    Add(int)
    Remove() (int, error)
}

type LinkedStack struct {
    root *llnode
}

type LinkedList struct {
    root *llnode
}

type ListError string

func (e ListError) Error() string {
    return string(e)
}

type llnode struct {
    data int
    next *llnode
}

func (ll *LinkedStack) Add(data int) {
    ll.root = &llnode{data, ll.root}
}

func (ll *LinkedStack) Remove() (int, error) {
    if ll.root == nil {
        return 0, ListError("Tried to remove from empty stack!")
    }
    data := ll.root.data
    ll.root = ll.root.next
    return data, nil
}
    

func MakeLinkedStack() *LinkedStack {
    return &LinkedStack{nil}
}

func (ll *LinkedList) Add(data int) {
    ll.root.next = &llnode{data, ll.root.next}
}

func (ll *LinkedList) Remove() (int, error) {
    if ll.root.next.next == nil {
        return 0, ListError("Tried to remove from empty list!")
    }
    cur := ll.root
    var prev *llnode
    for cur.next.next != nil {
        cur, prev = cur.next, cur // traverse list
    }
    
    prev.next = cur.next
    return cur.data, nil
}
    

func MakeLinkedList() *LinkedList {
    return &LinkedList{&llnode{0, &llnode{}}}
}

func main() {
    var s List = MakeLinkedStack()
    var l List = MakeLinkedList()
    
    for i := 0; i < 10; i++ {
        s.Add(i)
        l.Add(i)
    }
    for i := 0; i < 11; i++ {
        datas, errs := s.Remove()
        datal, errl := l.Remove()
        if errs != nil {
            fmt.Println(errs)
        } else {
            fmt.Println(datas)
        }
        if errl != nil {
            fmt.Println(errl)
        } else {
            fmt.Println(datal)
        }
    }
}
    
    
