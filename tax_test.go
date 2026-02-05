package main

import (
	"testing"
)

func TestComputeAnnualTax(t *testing.T) {
	tests := []struct {
		name          string
		annualSalary  float64
		expectedTax   float64
		description   string
	}{
		{
			name:         "Sample 1: Annual Salary 60000",
			annualSalary: 60000,
			expectedTax:  6000,
			description:  "First 20k at 0%, next 20k at 10%, next 20k at 20%",
		},
		{
			name:         "Sample 2: Annual Salary 200000",
			annualSalary: 200000,
			expectedTax:  48000,
			description:  "Progressive tax through all brackets",
		},
		{
			name:         "Sample 3: Annual Salary 80150",
			annualSalary: 80150,
			expectedTax:  10045,
			description:  "Tax includes partial bracket at 80001-80150",
		},
		{
			name:         "Salary at bracket boundary (20000)",
			annualSalary: 20000,
			expectedTax:  0,
			description:  "Entire salary in 0% bracket",
		},
		{
			name:         "Salary just above first bracket (20001)",
			annualSalary: 20001,
			expectedTax:  0.10,
			description:  "First 20k at 0%, 1 dollar at 10%",
		},
		{
			name:         "Salary at second bracket boundary (40000)",
			annualSalary: 40000,
			expectedTax:  2000,
			description:  "First 20k at 0%, next 20k at 10%",
		},
		{
			name:         "Zero salary",
			annualSalary: 0,
			expectedTax:  0,
			description:  "No tax on zero income",
		},
		{
			name:         "Negative salary",
			annualSalary: -1000,
			expectedTax:  0,
			description:  "No tax on negative income",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := ComputeAnnualTax(tt.annualSalary)
			if !floatEquals(result, tt.expectedTax, 0.01) {
				t.Errorf("ComputeAnnualTax(%f) = %f; want %f. %s",
					tt.annualSalary, result, tt.expectedTax, tt.description)
			}
		})
	}
}

func floatEquals(a, b, tolerance float64) bool {
	diff := a - b
	if diff < 0 {
		diff = -diff
	}
	return diff < tolerance
}
