package stack

import (
	"fmt"
)
type I interface{}
type Stack struct{
	len,cap int
	items []I
}

func NewStack() * Stack{
	return &Stack{
		len : 0,
		cap : 0,
	}
}


func (ps * Stack) resize(nSize int){
	if(nSize < ps.cap){
		panic("error resize, new size should be greater than old size")
	}
	tItems := ps.items
	fmt.Println(ps.len,nSize)
	ps.items = make([]I,nSize,nSize)
	for i := 0; i < ps.len; i++ {
		ps.items[i] = tItems[i]
	}
}

func (ps * Stack) size() int {
	return ps.len
}

func (ps * Stack) full() bool {
	return ps.len == ps.cap
}

func (ps * Stack) empty() bool {
	return ps.len == 0
}

func (ps * Stack) push(v I){
	if(ps.full()){
		newCap := (ps.cap + 1) * 2
		ps.resize(newCap)
		ps.cap = newCap
	}
	ps.items[ps.len] = v
	ps.len++
}

func (ps * Stack) pop() I{
	if(ps.empty()){
		panic("stack is empty")
	}
	v := ps.items[ps.len-1]
	ps.len--
	return v
}

func (ps * Stack) top() I{
	if(ps.empty()){
		panic("stack is empty")
	}
	v := ps.items[ps.len-1]
	return v
}