package main

import (
	"fmt"
)

func main() {
	ReadCSV("test", "x")
}

func ReadCSV(filename string, delimiter string) {
	fmt.Println("Trying to use file " + filename + " as csv with " + delimiter + " as delimiter")

}
