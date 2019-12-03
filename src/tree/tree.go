package tree

import (
	"fmt"
)

type T interface{
	Compare(t1 T) int
}

type Node struct{
	left,right * Node
	value T
	count int
}

func (n *Node) Compare (t *Node) int{
	return n.value.Compare(t.value)
}


type BinarySearachTree struct{
	root * Node
}

func NewTree() * BinarySearachTree{
	return &BinarySearachTree{}
}

func insert(currentNode,newNode *Node) *Node{
	if(currentNode == nil){
		return newNode
	}

	cpResult := currentNode.Compare(newNode)
	if(cpResult == 0){
		currentNode.count ++
	}else if(cpResult == -1){
		currentNode.right = insert(currentNode.right,newNode)
	}else{
		currentNode.left = insert(currentNode.left,newNode)
	}
	return currentNode
}

func (bs * BinarySearachTree) Insert (v T){
	newNode := &Node {
		value : v,
		count : 1,
	}
	bs.root = insert(bs.root,newNode)
}

func search(currentNode,newNode * Node) *Node{
	if(currentNode == nil){
		return nil
	}

	cpResult := currentNode.Compare(newNode)
	if(cpResult == 0){
		return currentNode
	}else if(cpResult == -1){
		return search(currentNode.right,newNode)
	}else{
		return search(currentNode.left,newNode)
	}
	return nil
}

func (bs * BinarySearachTree) Search (v T) int{
	newNode := &Node {
		value : v,
	}
	n := search(bs.root,newNode)
	if(n == nil){
		return 0
	}else{
		return n.count
	}
}

func preOrder(currentNode * Node){
	if(currentNode == nil){
		return
	}

	fmt.Println(currentNode.value,currentNode.count)
	preOrder(currentNode.left)
	preOrder(currentNode.right)
}

func (bs * BinarySearachTree) PreOrder (){
	preOrder(bs.root)
}


func midOrder(currentNode * Node){
	if(currentNode == nil){
		return
	}

	midOrder(currentNode.left)
	fmt.Println(currentNode.value,currentNode.count)
	midOrder(currentNode.right)
}

func (bs * BinarySearachTree) MidOrder (){
	midOrder(bs.root)
}


func backOrder(currentNode * Node){
	if(currentNode == nil){
		return
	}

	backOrder(currentNode.right)
	fmt.Println(currentNode.value,currentNode.count)
	backOrder(currentNode.left)
}

func (bs * BinarySearachTree) BackOrder (){
	backOrder(bs.root)
}