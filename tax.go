package main

import "math"

// TaxBracket represents a single tax bracket with its range and rate
type TaxBracket struct {
	Min  float64 // Inclusive minimum (e.g., 0, 20001)
	Max  float64 // Inclusive maximum (use math.Inf(1) for no upper limit)
	Rate float64 // Tax rate as decimal (e.g., 0.1 for 10%)
}

// TaxStrategy defines the interface for different tax calculation strategies
// This implements the Strategy Pattern, allowing different tax calculation
// algorithms to be used interchangeably
type TaxStrategy interface {
	CalculateAnnualTax(annualSalary float64) float64
	GetName() string
}

// ProgressiveTaxStrategy implements progressive tax calculation with brackets
type ProgressiveTaxStrategy struct {
	Brackets []TaxBracket
}

// NewProgressiveTaxStrategy creates a new progressive tax strategy with default brackets
func NewProgressiveTaxStrategy() *ProgressiveTaxStrategy {
	return &ProgressiveTaxStrategy{
		Brackets: []TaxBracket{
			{Min: 0, Max: 20000, Rate: 0.0},
			{Min: 20001, Max: 40000, Rate: 0.10},
			{Min: 40001, Max: 80000, Rate: 0.20},
			{Min: 80001, Max: 180000, Rate: 0.30},
			{Min: 180001, Max: math.Inf(1), Rate: 0.40},
		},
	}
}

// CalculateAnnualTax calculates tax using progressive brackets
func (s *ProgressiveTaxStrategy) CalculateAnnualTax(annualSalary float64) float64 {
	if annualSalary <= 0 {
		return 0
	}

	totalTax := 0.0

	for _, bracket := range s.Brackets {
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

// GetName returns the name of this tax strategy
func (s *ProgressiveTaxStrategy) GetName() string {
	return "Progressive Tax Strategy"
}

// DefaultTaxStrategy is the default tax calculation strategy used by the application
var DefaultTaxStrategy TaxStrategy = NewProgressiveTaxStrategy()

// ComputeAnnualTax calculates the total annual tax using the default strategy
// This function maintains backward compatibility with existing code
func ComputeAnnualTax(annualSalary float64) float64 {
	return DefaultTaxStrategy.CalculateAnnualTax(annualSalary)
}

// FlatTaxStrategy implements a simple flat tax rate
// This demonstrates the flexibility of the Strategy Pattern
type FlatTaxStrategy struct {
	Rate float64 // Flat tax rate (e.g., 0.15 for 15%)
}

// NewFlatTaxStrategy creates a new flat tax strategy
func NewFlatTaxStrategy(rate float64) *FlatTaxStrategy {
	return &FlatTaxStrategy{Rate: rate}
}

// CalculateAnnualTax calculates tax using a flat rate
func (s *FlatTaxStrategy) CalculateAnnualTax(annualSalary float64) float64 {
	if annualSalary <= 0 {
		return 0
	}
	return annualSalary * s.Rate
}

// GetName returns the name of this tax strategy
func (s *FlatTaxStrategy) GetName() string {
	return "Flat Tax Strategy"
}
