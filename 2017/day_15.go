package main

import (
	"fmt"
	"github.com/johncgriffin/overflow"
)

func run(value, factor, div int64) chan int64 {
	ch := make(chan int64, 10)

	go func() {
		for {
			value = overflow.Mul64p(value, factor)
			value = value % 2147483647

			if value % div == 0 {
				ch <- value
			}
		}
	}()

	return ch
}

func main() {
	a := run(512, 16807, 4)
	b := run(191, 48271, 8)
	//a := run(65, 16807, 4)
	//b := run(8921, 48271, 8)

	equal := 0
	for i := 0; i < 5*1000*1000; i ++ {
		av := <- a
		bv :=  <- b

		if av&0xffff == bv&0xffff {
			equal += 1
		}

		if i % 100000 == 0{
			fmt.Println(i)
		}
	}

	fmt.Println("equal", equal)
}
