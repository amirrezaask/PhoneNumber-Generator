package main

import (
	"fmt"
	"log"
	"os"
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
func Generate(start int, end int) {

	for i := 0; i+start <= end; i++ {
		writeToFile("db.txt", fmt.Sprintf("%d\n", start+i))
	}
}

func main() {
	counter := 0
	for i := 9121000000; i <= 9129999999; i++ {
		if counter%100 == 0 {
			log.Printf("Goroutine %v Started!!", counter/100)
			go Generate(i, i+100)
		}
		counter += 1
	}
	Generate(9121000000, 9129999999)

}
