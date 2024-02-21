package main

import (
	"math/rand"
	"testing"
)

// TestEstimatedMu checks if the estimated mu value is within 0.25 of 10.08.
func TestEstimatedMu(t *testing.T) {
	expectedMu := 10.08
	tolerance := 0.25

	// Fire up random generated and use seeding to fix the output.
	src := rand.NewSource(123)
	rnd := rand.New(src)

	// Generate observed data
	muTrue := 10.0
	sigma := 2.0
	n := 10000
	observedData := make([]float64, n)
	for i := range observedData {
		observedData[i] = rnd.NormFloat64()*sigma + muTrue
	}

	// ABC parameters
	start := 5.0
	stop := 15.0
	step := 0.1
	nSim := 100

	var acceptedMu []float64
	// Sequentially generate data and calculate mean
	for mu := start; mu <= stop; mu += step {
		simData := simulateData(nSim, mu, sigma, rnd)
		if abs(mean(simData)-mean(observedData)) <= tolerance {
			acceptedMu = append(acceptedMu, mu)
		}
	}

	estimatedMu := mean(acceptedMu)

	// Check if estimatedMu is within 0.25 of 10.08
	if abs(estimatedMu-expectedMu) > tolerance {
		t.Errorf("Estimated mu is not within %v of %v; got %v", tolerance, expectedMu, estimatedMu)
	}
}

// Benchmark the performance of the simulate data function.
func BenchmarkSimulateData(b *testing.B) {
	src := rand.NewSource(123)
	rnd := rand.New(src)
	mu := 10.0
	sigma := 2.0
	n := 10000 // n is configurable to test.

	// Run the function b.N times
	for i := 0; i < b.N; i++ {
		_ = simulateData(n, mu, sigma, rnd)
	}
}
