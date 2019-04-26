package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
)

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
		go Work(i)
	}
	for i := 0; i+start <= end; i++ {
		if counter%10000 == 0 {
			log.Printf("Goroutine %v Started!!", counter/10000)
			workerPool <- workerData{start: i + start, end: i + start + 10000}
		}
		counter += 1
	}
	if <-exit == 1 {
		return
	}
}
