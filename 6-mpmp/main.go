package main

import (
	"fmt"
	"sync"
)

const goal int = 1000000

func main() {

	var wg sync.WaitGroup

	fib := []int{0, 1}
	for len(fib) < 31 {
		fib = append(fib, fib[len(fib)-1]+fib[len(fib)-2])
	}

	solutions := make(map[int]int)

	for left, right := 0, len(fib)-1; left < right; left, right = left+1, right-1 {
		fib[left], fib[right] = fib[right], fib[left]
	}

	for i := range fib {

		fn := fib[i+1]
		fn1 := fib[i]

		for x := 0; x <= 1000000/fn1; x++ {
			p := fn1 * x
			if p > goal {
				break
			}

			wg.Add(1)
			go worker(&wg, fn1, fn, p, i)
		}

		if i == len(fib)-2 {
			break
		}

	}

	wg.Wait()

	fmt.Println(solutions)
}

func worker(wg *sync.WaitGroup, fn1, fn, p, i int) {
	defer wg.Done()

	for y := 0; y <= 1000000-p; y++ {
		q := fn * (y)
		if p+q > goal {
			break
		}
		if p+q == goal {
			fmt.Println(i)
		}
	}

}

func test(x, y int) {
	b := x + y
	b1 := x
	for b < 1000000 {
		b, b1 = b+b1, b
		fmt.Println(b, b1)
	}

}
