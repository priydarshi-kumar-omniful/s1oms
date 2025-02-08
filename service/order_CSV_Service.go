package service

import (
	"fmt"
	"oms/models"

	scsv "encoding/csv"
	"os"
	"oms/repo"
	"github.com/omniful/go_commons/csv"
)

// ParseAndCreateOrdersFromCSV processes the CSV file and creates orders in MongoDB
func ParseAndCreateOrdersFromCSV(filePath string) error {
	fmt.Println("Starting CSV processing for file:", filePath)

	file, err := os.Open(filePath)
	if err != nil {
		return fmt.Errorf("error opening CSV file: %w", err)
	}
	defer file.Close() // Ensure the file is closed after function execution

	csvFileReader := scsv.NewReader(file) // Standard Go CSV reader from the standard library

	csvReader, err := csv.NewCommonCSV(
		csv.WithBatchSize(1), // Process one record at a time
		csv.WithSource(csv.Local),
		csv.WithLocalFileInfo(filePath),
		csv.WithCSVReader(csvFileReader), // Pass the initialized CSV reader
	)

	// Ensure the reader is properly initialized
	if csvReader == nil || csvReader.Reader == nil {
		return fmt.Errorf("failed to initialize CSV reader")
	}

	fmt.Println("CSV reader initialized successfully")

	// Process CSV records
	for !csvReader.IsEOF() {
		var orders []models.Order

		err := csvReader.ParseNextBatch(&orders)
		if err != nil {
			return fmt.Errorf("error parsing batch: %w", err)
		}
		// fmt.Println("Parsed batch:", orders)
		
		if err:=repo.CreateOrder(orders);err!=nil{
			fmt.Println("error in inserting data in mongod")
			return err
		}
	}

	fmt.Println("CSV processing completed successfully")
	return nil
}
