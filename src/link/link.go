package link

import (
	"fmt"
)
type T interface{}
type Node struct {
	value T
	next *Node
}

type LinkedList struct {
	root *Node
}


func (p *LinkedList) insert(v T,c * Node) * Node{
	if(c != nil){
		c.next = p.insert(v,c.next)
		return c
	}else{
		return &Node{
			value : v,
		}
	}
}
func (p *LinkedList) Insert(v T){
	p.root = p.insert(v,p.root)
}

func (p *LinkedList) Show(){
	pCNode := p.root
	for pCNode != nil{
		fmt.Println(pCNode.value)
		pCNode = pCNode.next
	}
}

func (p *LinkedList) reverse(l,c * Node) * Node{
	if(c != nil){
		r := p.reverse(c,c.next)
		c.next = l
		return r
	}
	return l
}
func (p *LinkedList) Reverse(){
	p.root = p.reverse(nil,p.root)
}