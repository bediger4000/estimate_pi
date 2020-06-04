package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	rand.Seed(time.Now().UTC().UnixNano())

	total := 100000000
	in := 0

	for i := 0; i < total; i++ {
		xr := rand.Float64()
		yr := rand.Float64()
		d := xr*xr + yr*yr
		if d <= 1.0 {
			in++
		}
	}

	pi := 4.0 * float64(in) / float64(total)

	fmt.Printf("pi = %f\n", pi)
}
