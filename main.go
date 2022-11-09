package main

import (
	"fmt"
	"sync"
)

var wg = sync.WaitGroup{}

func main() {
	rundomNumbers := []int{}
	for i := 1; i <= 1000; i++ {
		rundomNumbers = append(rundomNumbers, i)
	}

	inputChan := generatePipeline(rundomNumbers)

	//fan-out
	c1 := fanOut(inputChan)
	c2 := fanOut(inputChan)
	c3 := fanOut(inputChan)

	//fan-in
	c := fanIn(c1, c2, c3)
	sum := 0
	for i := 0; i < len(rundomNumbers); i++ {
		sum += <-c
	}
	fmt.Println(sum)
}

func generatePipeline(numbers []int) <-chan int {
	out := make(chan int)
	go func() {
		for _, n := range numbers {
			out <- n
		}
		close(out)
	}()
	return out
}

func fanOut(in <-chan int) <-chan int {
	out := make(chan int)
	go func() {
		for n := range in {
			out <- n * n
		}
		close(out)
	}()
	return out
}

func fanIn(inputChannel ...<-chan int) <-chan int {
	in := make(chan int)
	go func() {
		for _, c := range inputChannel {
			for n := range c {
				in <- n
			}
		}
		close(in)
	}()
	return in
}
