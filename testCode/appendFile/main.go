package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func main() {
	//os.OpenFile("./test.log")
	f, err := os.OpenFile("./test.log", os.O_RDWR|os.O_CREATE, 0755)
	if err != nil {
		log.Fatal(err)
	}
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
	if err := f.Close(); err != nil {
		log.Fatal(err)
	}
}
