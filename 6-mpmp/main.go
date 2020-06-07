package main

import (
	"fmt"
)

func main() {

	goal := 1000000

	fib := []int{0, 1}
	for i := 0; i < 30; i++ {
		fib = append(fib, fib[len(fib)-1]+fib[len(fib)-2])
	}

	solutions := make(map[int]int)

	// fmt.Println(fib)

	for i, v := range fib {
		if i < 3 {
			continue
		}
		fn := v
		fn1 := fib[i+1]
		fmt.Println(fn1, fn)

		for x := 0; x <= 1000000/fn1; x++ {
			p := fn1 * x
			if p > goal {
				break
			}
			for y := 0; y <= 1000000-p; y++ {
				q := fn * (y)
				if p+q > goal {
					break
				}
				if p+q == goal {
					// fmt.Println(i, x, y)
					solutions[i] = solutions[i] + 1
				}
			}
		}

		if i == len(fib)-2 {
			break
		}

	}

	fmt.Println(solutions)

	// fmt.Println(max, maxi, maxj)
	// test(144, 154)
}

func test(x, y int) {
	b := x + y
	b1 := x
	for b < 1000000 {
		b, b1 = b+b1, b
		fmt.Println(b, b1)
	}

}
