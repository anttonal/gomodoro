package main

import (
	"flag"
	"fmt"
	"os"
	"strconv"
)

func main() {
	m := flag.Bool("m", false, "interpret as minutes")
	flag.Parse()
	args := flag.Args()

	// no time variable
	if len(args) != 1 {
		fmt.Fprintf(os.Stderr, "Usage: timer [-m] AMOUNT\n")
		os.Exit(1)
	}

	n, err := strconv.Atoi(args[0])

	// non int given
	if err != nil || n <= 0 {
		fmt.Fprintf(os.Stderr, "Usage: timer [-m] AMOUNT\n")
		os.Exit(1)
	}

	if *m {
		n *= 60
	}

	cfg := defaultConfig()
	_ = cfg

	countdown(n)
}
