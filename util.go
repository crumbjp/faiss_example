package main

import (
	"os"
	"strconv"
	"encoding/csv"
)

func GetData() *[][]float32 {
	file, _ := os.Open("data.csv")
	var result = &[][]float32{}
	reader := csv.NewReader(file)
	records, _ := reader.ReadAll()
	for _, record := range records {
		var vector = make([]float32, len(record))
		for i, _ := range record {
			val, _ := strconv.ParseFloat(record[i], 32)
			vector[i] = float32(val)
		}
		*result = append(*result, vector)
	}
	return result;
}
