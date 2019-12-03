package main
import (
	"fmt"
	"math"
)

type Abser interface{
	Scale(s float64)
}

type Vertex struct{
	X, Y float64
}

func (v Vertex) Abs() float64 {
	return math.Sqrt(v.X * v.X + v.Y * v.Y)
}

func (v *Vertex) Scale(s float64) {
	v.X *= float64(s)
	v.Y *= float64(s)
}

func main(){
	v := Vertex{3,4}
	(&v).Scale(11.11345)
	fmt.Println(v.Abs())
}