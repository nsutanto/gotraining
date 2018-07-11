// All material is licensed under the Apache License Version 2.0, January 2004
// http://www.apache.org/licenses/LICENSE-2.0

// go build
// ./example5

// Sample program to validate a trained regression model on a holdout data set.
package main

import (
	"encoding/csv"
	"fmt"
	"log"
	"math"
	"os"
	"strconv"
)

func main() {

	// Open the diabetes dataset file.
	f, err := os.Open("../data/holdout.csv")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	// Create a new CSV reader reading from the opened file.
	reader := csv.NewReader(f)

	// Read in all of the CSV records
	reader.FieldsPerRecord = 11
	holdoutData, err := reader.ReadAll()
	if err != nil {
		log.Fatal(err)
	}

	// Loop over the holdout data predicting y and evaluating the prediction
	// with the mean absolute error.
	var mAE float64
	for i, record := range holdoutData {

		// Skip the header.
		if i == 0 {
			continue
		}

		// Parse the observed diabetes progression measure, or "y".
		yObserved, err := strconv.ParseFloat(record[10], 64)
		if err != nil {
			log.Fatal(err)
		}

		// Parse the bmi value.
		bmiVal, err := strconv.ParseFloat(record[2], 64)
		if err != nil {
			log.Fatal(err)
		}

		// Predict y with our trained model.
		yPredicted := predict(bmiVal)

		// Add the to the mean absolute error.
		mAE += math.Abs(yObserved-yPredicted) / float64(len(holdoutData))
	}

	// Output the MAE to standard out.
	fmt.Printf("\nMAE = %0.2f\n\n", mAE)
}

// predict uses our trained regression model to made a prediction based on a
// bmi value.
func predict(bmi float64) float64 {
	return 149.96 + bmi*916.19
}
