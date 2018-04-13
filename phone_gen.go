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
	start, err := strconv.Atoi(os.Args[1])
	if err != nil {
		log.Fatalf("Error First Paramter should be a number")
	}
	end, err := strconv.Atoi(os.Args[2])
	if err != nil {
		log.Fatalf("Error Second Parameter should be a number")
	}
	for i := 0; i+start <= end; i++ {
		if counter%10000 == 0 {
			log.Printf("Goroutine %v Started!!", counter/10000)
			go Generate(fmt.Sprintf("%d_%d.txt", i+start, i+start+10000), i+start, i+start+10000)
		}
		counter += 1
	}

}
