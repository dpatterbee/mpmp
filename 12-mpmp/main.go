package main

import (
	"fmt"
	"sync"
)

func main() {

	var wg sync.WaitGroup

	for i := 2; i < 10000; i++ {

		wg.Add(1)
		go func(i int) {
			defer wg.Done()

			var count int

			for j := 1; j*j < i; j++ {
				if i%j == 0 {
					if j == i/j {
						count++
					} else {
						count += 2
					}
				}
			}

			if count == 64 {
				fmt.Println(i)
				for j := 1; j*j < i; j++ {
					if i%j == 0 {
						fmt.Println(j, i/j)
					}
				}
			}
		}(i)
		// fmt.Println(i, count)
	}

	wg.Wait()
}
