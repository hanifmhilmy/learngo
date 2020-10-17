package main

import (
	"flag"
	"log"
	"math/rand"
	"runtime"
	"time"
)

var spawn int

func init() {
	flag.IntVar(&spawn, "spawn", 1000, "number of the go routine will be spawned")
	flag.Parse()
}

func main() {
	log.Println(`Welcome to the learning go section ... this binary increase your memory consumption in the expected way`)

	// proofing the medium article
	// go routine will leak if we `Sending to a channel without receiver`
	for j := 1; j < spawn; j++ {
		log.Printf("Start spawning go routine ... spawner count (%d/1000)\n", j)
		go func() int {
			ch := make(chan int)
			go func() { ch <- query() }()
			go func() { ch <- query() }()
			go func() { ch <- query() }()
			return <-ch
		}()

		// Catcher so it will not break your system
		if runtime.NumGoroutine() > 100000 {
			break
		}

		time.Sleep(time.Duration(100) * time.Millisecond)
	}
}

func query() int {
	n := rand.Intn(100)
	time.Sleep(time.Duration(n) * time.Millisecond)
	return n
}
