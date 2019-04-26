package main

import (
	"fmt"
	"log"
)

type workerData struct {
	start int
	end   int
}

var workerPool = make(chan workerData, 900)
var count = make(chan int, 1)
var exit = make(chan int, 1)

func Work(id int) {
start:
	w := <-workerPool
	log.Printf("Worker %d Started Working\n", id)
	Generate(fmt.Sprintf("./output/%d_%d.txt", w.start, w.end), w.start, w.end)
	count <- 1 + <-count
	goto start
}
func HeadWorker() {
	counter := <-exit
	if counter == 900 {
		exit <- 1
		return
	}
	count <- counter
	return
}
