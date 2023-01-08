package main

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"
)

func init() {
	// Creat db connection

}

var ()

const ()

func main() {
	colIndex := map[string]int{
		"merchant_contact": -1,
		"merchant_name":    -1,
		"title":            -1,
		"amount":           -1,
		"start_time":       -1,
		"end_time":         -1,
		"detail":           -1,
		"image_url":        -1,
	}

	// 1. Load csv file
	records := load()
	header := records[0]
	records = records[1:]
	for i, col := range header {
		if _, ok := colIndex[col]; ok {
			colIndex[col] = i
		}
	}
	fmt.Println(colIndex)

	for _, record := range records {
		fmt.Println(record)
	}
	// 2. Data format

	// 3. Insert
	// 4. Show result
	// 5. 確認是否會有換行問題
}

func createProject() {

}
func insertPrize() {
	chainID := "poylgon"
	contractAddress := "prizes"
	ownerAddress := "0x1b0b3e5e9f1f7c7c3c6b1b0b3e5e9f1f7c7c3c6b"
	projectID := "1"     //
	projectName := "1"   //
	publistTime := "now" //
	insertSQL := "INSERT INTO prizes (chain_id, contract_address, owner_address, project_id, project_name, publist_time) VALUES ($1, $2, $3, $4, $5, $6)"
}

func load() [][]string {
	filePath := "data.csv"
	f, err := os.Open(filePath)
	if err != nil {
		log.Fatal("Unable to read input file "+filePath, err)
	}
	defer f.Close()

	csvReader := csv.NewReader(f)
	records, err := csvReader.ReadAll()
	if err != nil {
		log.Fatal("Unable to parse file as CSV for "+filePath, err)
	}

	return records
}
