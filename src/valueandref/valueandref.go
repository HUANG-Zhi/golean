package main

import (
	"fmt"
)

func mint(v int) int{
	v = 2
	return v
}

func mins(v string) string{
	v = "12"
	return v
}


//reference
func mina(v []int) []int{
	v[0] = 10
	// v = [] int{1,2}
	return v
}

type structa struct {
	i int
	s string
	a []int
}

func minstr(v *structa) *structa{
	v.i = 2
	v.s = "2"
	v.a[2] = 10
	return v
}

func minm(v map[string] int) map[string] int{
	v["ss"] = 10
	// v = [] int{1,2}
	return v
}

func main (){

	// s := "1"
	// rs := mins(s)
	// fmt.Println(s,rs)

	// a := [] int{1}
	// ra := mina(a)
	// ra[0] = 11
	// fmt.Println(a,ra)

	m := make(map[string] int)
	m["aa"] = 1

	mr := minm(m)
	mr["bb"] = 100

	fmt.Println(m,mr)
}