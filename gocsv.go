package main

import (
	"fmt"
	"os"
)

func main() {
	ReadCSV("test.csv", "x")
}

func ReadCSV(filename string, delimiter byte) {
	fmt.Println("Trying to use file " + filename + " as csv with " + delimiter + " as delimiter")

	//try reading file, defer close
	fileHandle, err := os.Open(filename)
	if err != nil {
		fmt.Println("file not found")
		os.Exit(99)
	}
	defer fileHandle.Close()

}
