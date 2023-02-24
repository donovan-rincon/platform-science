package internal

import "sort"

type Pair struct {
	Shipment string
	Driver   string
	SS       float64
}

func ProcessAssignments() []Pair {
	// Sort the shipments and drivers by their lengths in increasing order.
	// This ensures that we consider the smallest pairs first, which are more likely to have common factors and therefore a higher SS
	// Sort shipments by the length of the address
	sort.Slice(input.Shipments, func(i, j int) bool {
		return len(input.Shipments[i]) < len(input.Shipments[j])
	})

	// Sort drivers by the length of their name
	sort.Slice(input.Drivers, func(i, j int) bool {
		return len(input.Drivers[i]) < len(input.Drivers[j])
	})

	// Compute the suitability scores for each driver-shipment pair and store
	// We initialize the map by looping over all the possible pairs and computing their score using the getSuitabilityScore function.
	// This takes O(n^2) time, but we only have to do it once
	scores := make(map[string]map[string]float64)
	for _, shipment := range input.Shipments {
		for _, driver := range input.Drivers {
			if _, ok := scores[shipment]; !ok {
				scores[shipment] = make(map[string]float64)
			}
			scores[shipment][driver] = baseSuitabilityScore(shipment, driver)
		}
	}

	assignments := make([]Pair, len(input.Shipments))
	usedDrivers := make(map[string]bool)

	// Iterate over the sorted shipments and try to find the best driver for each one.
	// We do this by looking at the suitability scores in the scores map and selecting the highest-scoring driver that hasn't already been used.
	for i, shipment := range input.Shipments {
		var driver string
		maxSS := -1.0

		for candidate, ss := range scores[shipment] {
			if usedDrivers[candidate] {
				continue
			}

			if ss > maxSS {
				driver = candidate
				maxSS = ss
			}
		}

		assignments[i] = Pair{Shipment: shipment, Driver: driver, SS: maxSS}
		usedDrivers[driver] = true
	}
	return assignments
}
