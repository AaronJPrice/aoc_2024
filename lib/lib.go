package lib

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

// Reads file 'filename' into a slice of strings, each element in the slice is a line in the file
func ReadLines(filename string) ([]string, error) {
	f, err := os.Open(filename)
	if err != nil {
		return nil, fmt.Errorf("reading lines from file: %w", err)
	}
	defer f.Close()

	var data []string
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		data = append(data, scanner.Text())
	}

	err = scanner.Err()
	if err != nil {
		return nil, fmt.Errorf("reading lines from file (filename='%v'): %w", filename, err)
	}

	return data, nil
}

func IfErrLogFatal(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
