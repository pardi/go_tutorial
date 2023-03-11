package main

import (
	"fmt"
	"os"
	"strings"
	"time"
)

func main() {
	fmt.Println("Ch1 example: performance tests on different methods of updating variables in a loop")

	var s string
	sep := " "

	// Method 1
	start_timer := time.Now()

	for _, arg := range os.Args {
		s += arg + sep
	}

	end_timer := time.Now()

	fmt.Println("For loop time: " + end_timer.Sub(start_timer).String())

	start_timer = time.Now()

	// Method 2
	strings.Join(os.Args, sep)

	end_timer = time.Now()

	fmt.Println("String join time: " + end_timer.Sub(start_timer).String())

}
