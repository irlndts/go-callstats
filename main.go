package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

var (
	slider int
	input  string
	output string
)

func init() {
	flag.IntVar(&slider, "slider", 0, "slider's size")
	flag.StringVar(&input, "input", "", "path to input file")
	flag.StringVar(&output, "output", "", "path to output file")
	flag.Parse()
}

func main() {
	// check input parameters
	if slider == 0 || input == "" || output == "" {
		flag.Usage()
		return
	}

	// open input file
	inputFile, err := os.Open(input)
	if err != nil {
		fmt.Fprintf(os.Stderr, "failed to open input file: %s", err)
		return
	}
	defer inputFile.Close()

	// init output file
	outputFile, err := os.OpenFile(output, os.O_CREATE|os.O_RDWR, 0660)
	if err != nil {
		fmt.Fprintf(os.Stderr, "failed to create output file: %s", err)
		return
	}
	defer outputFile.Close()

	median := InitMedianQueue(slider)

	reader := csv.NewReader(inputFile)
	for {
		// read file line by line
		line, err := reader.Read()
		if err != nil {
			if err == io.EOF {
				// end of file
				break
			}
			fmt.Fprintf(os.Stderr, "failed to read csv: %s", err)
			return
		}

		for c, l := range line {
			// read data from the input file
			data, err := strconv.Atoi(strings.Trim(l, "\r "))
			if err != nil {
				fmt.Fprintf(os.Stderr, "failed to read column: %d line: %d: %s", c, l, err)
			}

			// add delay to the slider
			median.AddDelay(data)

			// write result to the output file
			_, err = outputFile.WriteString(fmt.Sprintln(median.GetMedian()))
			if err != nil {
				fmt.Fprintf(os.Stderr, "faile to write to file: %s", err)
			}
		}
	}
}
