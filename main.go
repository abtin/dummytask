package main

import (
	"flag"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"
)

var (
	cores    int
	duration int
)

func main() {
	start := time.Now()

	flag.IntVar(&cores, "c", 1, "-c <number-of-cores>")
	flag.IntVar(&duration, "t", 60, "-t <duration-in-seconds>")
	flag.Parse()

	log.Printf("Starting the task on %d cores and running it for %d seconds\n", cores, duration)
	Run(cores)

	c := make(chan os.Signal, 1)
	signal.Notify(c, syscall.SIGINT, syscall.SIGKILL, syscall.SIGTERM)

	select {
	case <-time.After(time.Duration(duration) * time.Second):
		log.Printf("Task timeout after %d seconds!", duration)
	case <-c:
		log.Printf("Task cancelled after %d milliseconds!\n", (time.Now().UnixNano()-start.UnixNano())/1e6)
	}

	log.Printf("Task finished - Worked for %d milliseconds!\n", (time.Now().UnixNano()-start.UnixNano())/1e6)
}
