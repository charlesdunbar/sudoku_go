package main

import (
	"bufio"
	"encoding/csv"
	"io"
	"log"
	"os"
)

func loadCsv(toOpen string) []string {
	var retVal []string
	f, err := os.Open(toOpen)
	if err != nil {
		log.Fatal(err)
	}
	r := csv.NewReader(bufio.NewReader(f))
	// Disable checking number of fields
	r.FieldsPerRecord = -1
	for {
		record, err := r.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatal(err)
		}
		// Skip first row
		if record[0] == "Puzzle" {
			continue
		}
		retVal = append(retVal, record[0])
	}

	return retVal
}
