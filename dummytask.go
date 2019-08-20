package main

import (
	"runtime"
	"time"
)

// Run simulates a %100 CPU load on the number of specified cores
func Run(cores int) {
	runtime.GOMAXPROCS(cores)
	for ; cores > 0; cores-- {
		go func() {
			for {
				time.Now()
			}
		}()
	}
}
