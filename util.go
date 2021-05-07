package main

import (
	"os"
	"strconv"
	"encoding/csv"
)

func GetData(path string, dim int) *[]float32 {
	file, _ := os.Open(path)
	reader := csv.NewReader(file)
	records, _ := reader.ReadAll()
	results := make([]float32, dim * len(records))
	for i, record := range records {
		for j, _ := range record {
			val, _ := strconv.ParseFloat(record[j], 32)
			results[i*dim+j] = float32(val)
		}
	}
	return &results;
}
