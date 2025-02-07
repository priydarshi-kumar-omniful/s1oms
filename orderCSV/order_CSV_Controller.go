package orderCSV

import (
	"context"
	"fmt"
	"log"
	"oms/models"
	"oms/utils"
	"os"

	"github.com/omniful/go_commons/csv"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

// MongoClient is the MongoDB client reference
var MongoClient *mongo.Client

// ParseAndCreateOrdersFromCSV processes the CSV file and creates orders in MongoDB
func ParseAndCreateOrdersFromCSV(filePath string) error {
	// Step 1: Open the CSV file
	file, err := os.Open(filePath)
	if err != nil {
		return fmt.Errorf("error opening CSV file: %w", err)
	}
	defer file.Close()

	// Step 2: Initialize CSV reader using the CommonCSV package
	csvReader, err := csv.NewCommonCSV(
		csv.WithBatchSize(1), // Process one record at a time (you can adjust batch size)
		csv.WithSource(csv.Local),
		csv.WithLocalFileInfo(filePath),
	)
	if err != nil {
		return fmt.Errorf("error initializing CSV reader: %w", err)
	}

	// Step 3: Initialize the CSV reader
	err = csvReader.InitializeReader(context.TODO())
	if err != nil {
		return fmt.Errorf("error initializing reader: %w", err)
	}

	// Step 4: Read and process records from CSV
	for !csvReader.IsEOF() {
		records, err := csvReader.ReadNextBatch()
		if err != nil {
			return fmt.Errorf("error reading next batch: %w", err)
		}

		// Step 5: Create orders from CSV records
		for _, record := range records {
			tenantID := record[0]
			customerID := record[1]
			productID := record[2]
			name := record[3]

			// Parse quantity (ensure it's an integer)
			quantity := utils.ParseInt(record[4])
			if err != nil {
				log.Printf("Error parsing quantity for product %s: %v. Skipping record.", productID, err)
				continue // Skip this record
			}

			// Parse price (ensure it's a float)
			price := utils.ParseFloat(record[5])
			if err != nil {
				log.Printf("Error parsing price for product %s: %v. Skipping record.", productID, err)
				continue // Skip this record
			}

			// Parse warehouse_id and sku_id
			warehouseID := record[6]
			skuID := record[7]

			// Create an order item
			orderItem := models.OrderItem{
				ProductID:   productID,
				SKUID:       skuID,
				Name:        name,
				Quantity:    quantity,
				WarehouseID: warehouseID,
				Price:       price,
			}

			// Create the order
			order := models.Order{
				ID:          primitive.NewObjectID(),
				TenantID:    tenantID,
				CustomerID:  customerID,
				Items:       []models.OrderItem{orderItem},
				TotalAmount: float64(quantity) * price, // Calculate total for this item
				Status:      models.OnHold,
			}

			// Insert the order into MongoDB
			collection := MongoClient.Database("oms_Service").Collection("orders")
			_, insertErr := collection.InsertOne(context.Background(), order)
			if insertErr != nil {
				log.Printf("Error inserting order for CustomerID %s: %v", customerID, insertErr)
			} else {
				log.Printf("Order created successfully for CustomerID %s", customerID)
			}
		}
	}

	return nil
}
