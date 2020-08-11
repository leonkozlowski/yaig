package in

import (
	"fmt"
	"io/ioutil"
	"os"
	"log"
)

func reader() {
	file, err := os.Open("book.txt")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Print(file)
	defer file.Close()
}