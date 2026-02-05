package main

import "math"

// TaxBracket represents a single tax bracket with its range and rate
type TaxBracket struct {
	Min  float64 // Inclusive minimum (e.g., 0, 20001)
	Max  float64 // Inclusive maximum (use math.Inf(1) for no upper limit)
	Rate float64 // Tax rate as decimal (e.g., 0.1 for 10%)
}

// TaxBrackets holds all tax brackets in order
var TaxBrackets = []TaxBracket{
	{Min: 0, Max: 20000, Rate: 0.0},
	{Min: 20001, Max: 40000, Rate: 0.10},
	{Min: 40001, Max: 80000, Rate: 0.20},
	{Min: 80001, Max: 180000, Rate: 0.30},
	{Min: 180001, Max: math.Inf(1), Rate: 0.40},
}

// ComputeAnnualTax calculates the total annual tax based on progressive tax brackets
func ComputeAnnualTax(annualSalary float64) float64 {
	if annualSalary <= 0 {
		return 0
	}

	totalTax := 0.0
	
	for _, bracket := range TaxBrackets {
		// Skip if salary doesn't reach this bracket
		if annualSalary < bracket.Min {
			break
		}

		// Calculate the upper bound for this bracket
		upperBound := bracket.Max
		if annualSalary < upperBound {
			upperBound = annualSalary
		}

		// Calculate taxable amount in this bracket
		taxableInBracket := upperBound - bracket.Min + 1
		
		// Calculate tax for this bracket
		tax := taxableInBracket * bracket.Rate
		totalTax += tax
	}

	return totalTax
}
