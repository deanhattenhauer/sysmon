package main

import (
	"fmt"
	"os"

	"github.com/deanhattenhauer/sysmon/internal/collect"
)

func main() {
	load, err := collect.ReadLoad()
	if err != nil {
		fmt.Fprintf(os.Stderr, "could not read load: %v\n", err)
		os.Exit(1)
	}
	fmt.Printf("Load (1m, 5m, 15m) = %.2f, %.2f, %.2f\n", load.One, load.Five, load.Fifteen)
}
