package main

import (
	"fmt"
	"log"
	"math"
	"math/rand"
	"os"
	"strconv"
	"time"
)

func main() {
	a, err := strconv.ParseFloat(os.Args[1], 64)
	if err != nil {
		log.Fatal(err)
	}
	l, err := strconv.ParseFloat(os.Args[2], 64)
	if err != nil {
		log.Fatal(err)
	}
	n, err := strconv.Atoi(os.Args[3])
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("cracks %f apart\n", a)
	fmt.Printf("needle %f long\n", l)
	fmt.Printf("%d trials\n", n)

	if l > a {
		fmt.Println("Cracks have to be farther apart than needle length")
		return
	}

	rand.Seed(time.Now().UTC().UnixNano())
	twoPi := 2.0 * math.Pi

	m := 0 // count of times needle crosses a crack
	for i := 0; i < n; i++ {
		phi := rand.Float64() * twoPi // random number of radians between 0 and 2Pi

		x := a * rand.Float64() // distance of needle center from a crack
		z := (l / 2.0) * math.Sin(phi)

		// needle overlaps crack above, or crack below
		if x+z >= a || x-z <= 0.0 {
			m++
		}
	}
	pi := l * float64(n) / (a * float64(m))
	fmt.Printf("%f\n", pi)
}
