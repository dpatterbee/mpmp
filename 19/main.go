package main

import "fmt"

func main() {
	ps := genPrimes(10000000)

	// fmt.Println(ps[:19])
	// fmt.Println(sumSq(ps[:19]) % 19)

	for i := 2; i < len(ps); i++ {
		sum := sumSq(ps[:i])
		if sum%i == 0 {
			fmt.Println(i)
		}
	}
}

func sumSq(l []int) int {
	var tot int
	m := len(l)
	for _, v := range l {
		tot += (v * v) % m
	}
	return tot
}

func genPrimes(n int) []int {
	sieve := make([]bool, n)

	for i := 2; i*i <= n; i++ {
		if sieve[i] {
			continue
		}
		for j := 0; i*i+j*i < n; j++ {
			sieve[i*i+j*i] = true
		}
	}
	var out []int
	for i, v := range sieve {
		if i < 2 {
			continue
		}
		if !v {
			out = append(out, int(i))
		}
	}
	return out
}
