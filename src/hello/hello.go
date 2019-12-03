package main

import (
	"fmt"
)

func WordCount(s string) map[string]int {
	return map[string]int{"x": 1}
}

func main() {
	m := make(map[string] int)
	m["aa"] = 1
	a,k := m["aa"]
	fmt.Println(a,k)
}
