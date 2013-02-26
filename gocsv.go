package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"
)

type csvHead string
type CSVLine map[string]string
type CSVContents struct {
	csvLines    map[int]CSVLine
	csvHeadline []string
}

func main() {
	var delimiter string = ","
	if _, err := ReadCSV("test.csv", delimiter); err != nil {
		fmt.Println("Aborted with err ", err)
	}
}

func buildCSVLine(csvContents *CSVContents, line, delimiter string, rowCount int) {
	csvLine := make(CSVLine)
	lineValues := strings.Split(line, delimiter)
	for key, value := range csvContents.csvHeadline {
		csvLine[value] = lineValues[key]
	}
	csvContents.csvLines[rowCount] = csvLine
	fmt.Println(csvContents.csvLines)
}

func buildCSVHeadline(csvContents *CSVContents, line, delimiter string) {
	csvContents.csvHeadline = strings.Split(line, delimiter)
	csvContents.csvLines = make(map[int]CSVLine)
	fmt.Println(csvContents.csvHeadline)
}

func ReadCSV(filename string, delimiter string) (csvContents CSVContents, err error) {

	fmt.Println("Trying to use file " + filename + " as csv with " + delimiter + " as delimiter")

	//try reading file, defer close
	fileHandle, err := os.Open(filename)
	if err != nil {
		return csvContents, err
	}
	defer fileHandle.Close()

	//read file linewise until nothing is left
	reader := bufio.NewReader(fileHandle)
	var isHeadline bool = true
	var rowCount int = 0
	for {
		line, err := reader.ReadString(10)
		//remove leading and trailing whitespaces/newlines
		line = strings.TrimSpace(line)
		//if this is the first line, use it as header
		if isHeadline {
			buildCSVHeadline(&csvContents, line, delimiter)
			isHeadline = false
		} else {
			buildCSVLine(&csvContents, line, delimiter, rowCount)
			rowCount++
			fmt.Println(line)
		}

		if err == io.EOF {
			break
		} else if err != nil {
			return csvContents, err
		}
	}
	return csvContents, nil

}
