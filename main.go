package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	startTime := time.Now() // Fire up the timer

	// Start a random generator and fix the seeding to keep stable results.
	src := rand.NewSource(123)
	rnd := rand.New(src)

	// Generate observed data.  Use preallocated slice to improve performance.
	muTrue := 10.0
	sigma := 2.0
	n := 10000
	observedData := make([]float64, n) // Preallocation is already done here, which is good practice
	for i := range observedData {
		observedData[i] = rnd.NormFloat64()*sigma + muTrue
	}

	// ABC parameters
	start := 5.0
	stop := 15.0
	step := 0.1
	tolerance := 0.5
	nSim := 100

	estimatedSize := int((stop-start)/step) + 1 // Calculate an upper limit based on the range and step of mu
	acceptedMu := make([]float64, 0, estimatedSize)

	// ABC algorithm
	for mu := start; mu <= stop; mu += step {
		simData := simulateData(nSim, mu, sigma, rnd) // Use rnd
		if abs(mean(simData)-mean(observedData)) <= tolerance {
			acceptedMu = append(acceptedMu, mu)
		}
	}

	estimatedMu := mean(acceptedMu)
	fmt.Printf("Estimated mu: %v\n", estimatedMu)

	// Stop timing and print
	executionTime := time.Since(startTime)
	fmt.Printf("Execution time: %s\n", executionTime)
}

// Modified simulateData to accept *rand.Rand for deterministic random number generation
// No change needed here for preallocation, as slices are created within each call and used immediately.
func simulateData(n int, mu, sigma float64, rnd *rand.Rand) []float64 {
	data := make([]float64, n)
	for i := range data {
		data[i] = rnd.NormFloat64()*sigma + mu
	}
	return data
}

func mean(numbers []float64) float64 {
	total := 0.0
	for _, number := range numbers {
		total += number
	}
	return total / float64(len(numbers))
}

func abs(x float64) float64 {
	if x < 0 {
		return -x
	}
	return x
}
