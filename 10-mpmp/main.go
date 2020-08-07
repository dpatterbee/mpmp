package main

import (
	"fmt"
	"math"
)

var count int
var squares [][4][2]int

func main() {
	count = 0
	genSquares([]int{}, 4)
	var newSquares [][4][2]int
	for _, v := range squares {
		if i, j := isSqr(v); i {
			newSquares = append(newSquares, v)
			fmt.Println(v, j)
		}
	}
	squares = newSquares
	genBoards(25, []int{}, 13)
	fmt.Println(count)
}

func isSqr(pts [4][2]int) (bool, int) {
	d2 := dist(pts[0], pts[1])
	d3 := dist(pts[0], pts[2])
	d4 := dist(pts[0], pts[3])

	if d2 == 0 || d3 == 0 || d4 == 0 {
		return false, 0
	}

	if d2 == d3 && 2*d2 == d4 && 2*dist(pts[1], pts[3]) == dist(pts[1], pts[2]) {
		return true, 1
	}

	// if d3 == d4 && 2*d3 == d2 && 2*dist(pts[2], pts[1]) == dist(pts[2], pts[3]) {
	// 	return true, 2
	// }

	// if d2 == d4 && 2*d2 == d3 && 2*dist(pts[1], pts[2]) == dist(pts[1], pts[3]) {
	// 	return true, 3
	// }
	return false, 0

}

func dist(a, b [2]int) int {
	dist := math.Pow(float64(a[0]-b[0]), 2) + math.Pow(float64(a[1]-b[1]), 2)
	return int(dist)
}

func checkSquares(board [5][5]bool) {

	for _, v := range squares {
		if (board[v[0][0]][v[0][1]] && board[v[1][0]][v[1][1]] && board[v[2][0]][v[2][1]] && board[v[3][0]][v[3][1]]) || (!board[v[0][0]][v[0][1]] && !board[v[1][0]][v[1][1]] && !board[v[2][0]][v[2][1]] && !board[v[3][0]][v[3][1]]) {
			return
		}
	}

	fmt.Println(board)
	count++

}

func genSquares(subset []int, subsetSize int) {

	var iStart int

	if len(subset) != 0 {
		iStart = subset[len(subset)-1] + 1
	}

	for i := iStart; i < 25; i++ {
		newSubset := append(subset, i)
		if len(newSubset) == subsetSize {
			var coords [4][2]int
			for i, v := range newSubset {
				coords[i] = [2]int{v / 5, v % 5}
			}
			squares = append(squares, coords)
		} else {
			genSquares(newSubset, subsetSize)
		}
	}

}

func makeBoard(set []int) {

	set = removeDupes(set)
	if len(set) < 13 {
		return
	}

	board := [5][5]bool{}

	for _, v := range set {
		board[v/5][v%5] = true
	}

	checkSquares(board)
}

func removeDupes(set []int) []int {
	trueMap := make(map[int]bool)
	for _, v := range set {
		trueMap[v] = true
	}

	var trueSet []int

	for i := range trueMap {
		trueSet = append(trueSet, i)
	}

	return trueSet
}

func genBoards(setLength int, subset []int, subsetSize int) {
	var iStart int

	if len(subset) != 0 {
		iStart = subset[len(subset)-1] + 1
	}

	for i := iStart; i < setLength; i++ {
		newSubset := append(subset, i)
		if len(newSubset) == subsetSize {
			go makeBoard(newSubset)
		} else {
			genBoards(setLength, newSubset, subsetSize)
		}
	}
}
