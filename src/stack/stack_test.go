package stack
import (
	"testing"
	"fmt"
)

func TestStack( t *testing.T){
	var a,b interface {}
	a,b = "sss","sss"
	fmt.Println(a==b)
	intStack := NewStack()
	fmt.Println(intStack.empty())
	intStack.push("sss")
	intStack.push(1)
	fmt.Println(intStack.empty())
	fmt.Println(intStack.pop())
	fmt.Println(intStack.pop())
	fmt.Println(intStack.empty())
}