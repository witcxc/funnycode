package main

import (
	"fmt"
	"strconv"
	"sync"
)

func main() {

	var wg sync.WaitGroup
	for i := 0; i < 50; i++ {
		wg.Add(1)
		go func(si int) {
			defer wg.Done()
			t := strconv.Itoa(si)
			for i := 0; i < 90; i++ {
				t += strconv.Itoa(si)
			}

			fmt.Printf("%s\n", t)
		}(i)
	}
	wg.Wait()
}
