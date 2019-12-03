package tree

import (
	"testing"
	"fmt"
	"time"
	"strconv"
)

type IntType int
func (t IntType) Compare( t1 T) int{
	if(t < t1.(IntType)){
		return -1
	}else if(t > t1.(IntType)){
		return 1
	}
	return 0
}

func TestTree(t *testing.T){
	nT := NewTree()

	// nT.Insert(IntType(3))
	// nT.Insert(IntType(2))
	// nT.Insert(IntType(4))
	// nT.Insert(IntType(2))
	// nT.Insert(IntType(9))
	// nT.Insert(IntType(3))
	// nT.Insert(IntType(1))
	nT.BackOrder()
	tsMs := time.Now().UnixNano()/time.Millisecond.Nanoseconds()
	v := strconv.FormatInt(tsMs, 10)
	fmt.Println(v)
}