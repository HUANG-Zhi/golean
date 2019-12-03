package main
import (
	"fmt"
)


// [] , chan map 是地址传递，结构体是值传递
func sum (s []int, c chan int){
	sum := 0
	for _,v := range s {
		sum += v
	}

	c <- sum
}

func main(){
	s := [] int { 1, 2, 3, -4, 5, 6}
	c := make(chan int)

	go sum(s[:len(s) / 2],c)
	go sum(s[len(s) / 2 :],c)

	x,y := <-c, <-c
	fmt.Println(x,y,x+y)
}