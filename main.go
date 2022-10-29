package main

import (
	"fmt"
	"strconv"
)

func main() {
	i := "12"

	j, er := strconv.Atoi(i)

	fmt.Printf("%T, %v", er, j)
}
