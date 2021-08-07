package main

import (
	"flag"
	"fmt"
	"strconv"
)

func main() {
	n := flag.Int("n", 1, "any positive integer")
	flag.Parse()

	if *n < 1 {
		fmt.Printf("Your choice of %d is invalid, n must be a positive integer\n", n)
		return
	}

	steps, sequence := collatzConjecture(*n)
	fmt.Printf("Reached 1 in %d steps\nWith sequence %s\n\n", steps, sequence)
}

func collatzConjecture(n int) (int, string) {
	steps := 0
	seq := strconv.Itoa(n)

	for n != 1 {
		if n%2 == 0 {
			n = n / 2
		} else {
			n = n*3 + 1
		}
		steps++
		seq = fmt.Sprintf("%s, %s", seq, strconv.Itoa(n))
	}
	return steps, seq
}
