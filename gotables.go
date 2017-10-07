package main

import (
	"fmt"
	"flag"
	"math/rand"
	"time"
)

func main() {
	//seeding random to the current time ensures it will be reasonably random
	rand.Seed( time.Now().UTC().UnixNano())
	//commandline arguments
	uptoInt := flag.Int("i", 12, "The highest integer in the times table")
	flag.Parse()
	//int i will store the user input
	var i int
	//main loop. Continues until Ctrl-C
	for {
		//set the random factors
		x := rand.Intn(*uptoInt)
		y := rand.Intn(*uptoInt)
		//print the equation
		fmt.Printf("%dx%d = ",x ,y )
		//get user's answer
		fmt.Scan(&i)
		//check user's answer
		if i != x*y {
			fmt.Printf("Nope. %dx%d = %d\n", x, y, x*y)
		}
	}
}
