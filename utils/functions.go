package utils

import(
	"strconv"
	"log"
)

// Helper function to parse integer from string
func ParseInt(value string) int {
	
	val, err := strconv.Atoi(value)
	if err != nil {
		log.Printf("Error parsing integer value: %v", err)
		return 0
	}
	return val
}

// Helper function to parse float from string
func ParseFloat(value string) float64 {
	
	val, err := strconv.ParseFloat(value, 64)
	if err != nil {
		log.Printf("Error parsing float value: %v", err)
		return 0.0
	}
	return val
}