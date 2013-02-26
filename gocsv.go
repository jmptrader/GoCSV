package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

func main() {
	var delimiter string = ","
	if err := ReadCSV("test.csv", delimiter); err != nil {
		fmt.Println("Aborted with err ", err)
	}
}

func ReadCSV(filename string, delimiter string) (err error) {
	fmt.Println("Trying to use file " + filename + " as csv with " + delimiter + " as delimiter")

	//try reading file, defer close
	fileHandle, err := os.Open(filename)
	if err != nil {
		return err
	}
	defer fileHandle.Close()

	//read file linewise until nothing is left
	reader := bufio.NewReader(fileHandle)
	for {
		line, err := reader.ReadString(10)
		fmt.Println(line)
		if err == io.EOF {
			break
		} else if err != nil {
			return err
		}
	}
	return nil

}
