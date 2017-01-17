package main

import (
	"fmt"
	"strconv"
)

type Int int

func (i Int) String() string {
	return strconv.Itoa(int(i))
}

func main() {
	i := Int(7)
	fmt.Printf("i: %s\n", i)
}
