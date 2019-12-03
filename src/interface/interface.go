package main

import (
	"fmt"
	"math"
)

type I interface {
	M()
}

type T struct {
	S string
}