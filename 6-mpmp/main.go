package main

import (
	"fmt"
	"sort"
)

const goal int = 1000000

type solution struct {
	dep1  int
	dep2  int
	steps int
}

func main() {

	fib := []int{0, 1}
	for len(fib) < 40 {
		fib = append(fib, fib[len(fib)-1]+fib[len(fib)-2])
	}

	var solutions []solution

	for left, right := 0, len(fib)-1; left < right; left, right = left+1, right-1 {
		fib[left], fib[right] = fib[right], fib[left]
	}

	for i := range fib {

		fn := fib[i+1]
		fn1 := fib[i]

		for x := 0; x <= goal/fn1; x++ {
			p := fn1 * x
			if p > goal {
				break
			}

			g := goal - p
			var y int
			if g%fn == 0 {
				y = g / fn

				solutions = append(solutions, solution{
					dep1:  x,
					dep2:  y,
					steps: findSteps(x, y),
				})
			}

		}

		if i == len(fib)-3 {
			break
		}

	}

	sort.Sort(bySteps(solutions))
	fmt.Println(solutions[0:10])

}

func findSteps(x, y int) int {
	days := 2

	prevBalance := x
	balance := x + y

	for balance < goal {
		prevBalance, balance = balance, prevBalance+balance
		days++
	}

	if balance > goal {
		panic(x)
	}

	return days
}

type bySteps []solution

func (s bySteps) Len() int {
	return len(s)
}
func (s bySteps) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}
func (s bySteps) Less(i, j int) bool {
	return s[i].steps > s[j].steps
}

func test(x, y int) {
	b := x + y
	b1 := x
	for b < 1000000 {
		b, b1 = b+b1, b
		fmt.Println(b, b1)
	}

}
