package main

import "fmt"

// 返回一个“返回int的函数”
func fibonacci() func() int {
	var i,j int = 0,1
	return func() int{
		// i = j;
		// j = i+ j
		i,j = j, i+j
		return j
	}
}


func main() {
	f := fibonacci()
	for i := 0; i < 10; i++ {
		fmt.Println(f())
	}
}
