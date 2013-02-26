package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

type csvHead string
type csvLine map[csvHead]string
type csvContents struct {
	csvLines    []csvLine
	csvHeadline []string
}

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
	var isHeadline bool = true
	for {
		line, err := reader.ReadString(10)

		//if this is the first line, use it as header
		if isHeadline {
			BuildCSVHeadline(line)
			isHeadline = false
		} else {
			fmt.Print(line)
		}

		if err == io.EOF {
			break
		} else if err != nil {
			return err
		}
	}
	return nil

}
