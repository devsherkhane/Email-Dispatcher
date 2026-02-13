package main

import (
	"encoding/csv"
	"os"
)

func loadRecipent(filepath string, ch chan Recipent) error {
	defer close(ch)
	file, err := os.Open(filepath)
	if err != nil {
		return err
	}
	defer file.Close()

	r := csv.NewReader(file)
	records, err := r.ReadAll()
	if err != nil {
		return err
	}

	for _, record := range records[1:] {
		ch <- Recipent{
			Name:  record[0],
			Email: record[1],
		}
	}

	return nil
}
