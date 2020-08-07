package main

import "fmt"

func main() {
	run()
	// n := 123123123
	// fmt.Println(count(0, 7, 7*n))
}

func run() {
	n := 1000
	repeat := 4

	greatestCount := 1
	ttt := [][3]int{{0, 0, 7}}

	for x := 0; x < n; x++ {
		for y := x; y < n; y++ {
			for z := y; z < n; z++ {
				i, j, k := x, y, z
				c := 0
				v := i + j + k
				for {
					i, j, k = nextTriangle(i, j, k)
					if i+j+k == v {
						c++
					} else {
						v = i + j + k
						c = 0
					}
					if c == repeat {
						break
					}
				}
				if i+j+k == 14 {
					currCount := count(x, y, z)
					if currCount > greatestCount {
						greatestCount = currCount
						ttt = [][3]int{{x, y, z}}

					} else if currCount == greatestCount {
						ttt = append(ttt, [3]int{x, y, z})
					}
				}
			}
		}
	}
	fmt.Println(ttt, greatestCount)
}

func count(i, j, k int) int {
	sum := i + j + k
	count := 0

	for sum != 14 {
		i, j, k = nextTriangle(i, j, k)
		sum = i + j + k
		count++
	}

	return count
}

func nextTriangle(i, j, k int) (x, y, z int) {

	x = abs(i - j)
	y = abs(j - k)
	z = abs(k - i)
	return

}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
