// golang
package main

import (
	"fmt"
	"time"
)

func main() {
	start := time.Njj
	for i := 0; i < 990999099; i++ {
		j += i
	}
	fmt.Println(j)

	cost := time.Since(start)
	fmt.Printf("elapsed time：%s", cost)
}
