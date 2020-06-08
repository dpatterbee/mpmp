package main

import (
	"fmt"
	"math"
	"os"
	"strconv"
)

var n int

func main() {
	start := []int{}
	if len(os.Args) < 2 {
		panic("no size specified")
	}
	n1, err := strconv.Atoi(os.Args[1])
	if err != nil {
		panic(err)
	}
	n = n1
	checkerboard(n*n, start, n)
}

func distanceCheck(board [][2]int) {

	var dists []float64
	for i, v := range board {
		for j := i + 1; j < len(board); j++ {
			dists = append(dists, euclid(v, board[j]))
		}
	}

	if isUnique(dists) {
		fmt.Println(board)
	}
}

func isUnique(dists []float64) bool {
	keys := make(map[float64]bool)

	for _, v := range dists {
		if _, ok := keys[v]; ok {
			return false
		}
		keys[v] = true
	}
	return true
}

func euclid(x, y [2]int) float64 {
	xDiff := float64(x[0] - y[0])
	yDiff := float64(x[1] - y[1])
	return math.Sqrt(math.Pow(xDiff, 2) + math.Pow(yDiff, 2))
}

func checkerboard(setLength int, subset []int, subsetSize int) {
	var iStart int

	if len(subset) != 0 {
		iStart = subset[len(subset)-1] + 1
	}

	for i := iStart; i < setLength; i++ {
		newSubset := append(subset, i)
		if len(newSubset) == subsetSize {

			var board [][2]int
			for _, v := range newSubset {
				board = append(board, [2]int{v / n, v % n})
			}
			go distanceCheck(board)

		} else {
			checkerboard(setLength, newSubset, subsetSize)
		}
	}

}
