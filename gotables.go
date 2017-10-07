package main

import (
	"fmt"
	"flag"
	"math/rand"
	"time"
)

func main() {
	rand.Seed( time.Now().UTC().UnixNano())
	uptoInt := flag.Int("i", 12, "The highest integer in the times table")
	flag.Parse()
	var i int
	for {
		x := rand.Intn(*uptoInt)
		y := rand.Intn(*uptoInt)
		fmt.Printf("%dx%d = ",x ,y )
		fmt.Scan(&i)
		if i != x*y {
			fmt.Printf("Nope. %dx%d = %d\n", x, y, x*y)
		}
	}
}
