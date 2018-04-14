package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
)

//var prefix = "/User/amirrezaask/Desktop/phones"

type workerData struct {
	start int
	end   int
}

var workerPool = make(chan workerData, 900)
var count = make(chan int, 1)
var exit = make(chan int, 1)

func writeToFile(path, text string) error {
	f, err := os.OpenFile(path, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0777)
	if err != nil {
		return err
	}
	defer f.Close()
	_, err = f.Write([]byte(text))
	if err != nil {
		return err
	}
	return nil
}
func Work() {
	w := <-workerPool
	Generate(fmt.Sprintf("./output/%d_%d.txt", w.start, w.end), w.start, w.end)
	count <- 1 + <-count
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
func Generate(filename string, start int, end int) {
	for i := 0; i+start <= end; i++ {
		writeToFile(filename, fmt.Sprintf("%d\n", start+i))
	}

}

func main() {
	counter := 0
	exit <- 1
	prefix, err := strconv.Atoi(os.Args[1])
	if err != nil {
		log.Fatalf("Argument should be cellphone prefix:%v", err)
	}
	start, err := strconv.Atoi(fmt.Sprintf("%d1000000", prefix))
	if err != nil {
		log.Fatalf("Something went wrong :%v", err)
	}
	end, err := strconv.Atoi(fmt.Sprintf("%d9999999", prefix))
	fmt.Println(end - start)

	if err != nil {
		log.Fatalf("Something went wrong :%v", err)
	}
	go HeadWorker()
	for i := 0; i < 900; i++ {
		go Work()
	}
	for i := 0; i+start <= end; i++ {
		if counter%10000 == 0 {
			log.Printf("Goroutine %v Started!!", counter/10000)
			workerPool <- workerData{start: i + start, end: i + end}
		}
		counter += 1
	}
	if <-exit == 1 {
		return
	}
}
