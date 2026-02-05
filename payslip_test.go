package main

import (
	"testing"
)

func TestGenerateMonthlyPayslip(t *testing.T) {
	tests := []struct {
		name                      string
		employeeName              string
		annualSalary              float64
		expectedGrossMonthly      float64
		expectedMonthlyTax        float64
		expectedNetMonthly        float64
	}{
		{
			name:                 "Sample case: Ren with 60000 salary",
			employeeName:         "Ren",
			annualSalary:         60000,
			expectedGrossMonthly: 5000.00,
			expectedMonthlyTax:   500.00,
			expectedNetMonthly:   4500.00,
		},
		{
			name:                 "High salary case: 200000",
			employeeName:         "Alice",
			annualSalary:         200000,
			expectedGrossMonthly: 16666.67,
			expectedMonthlyTax:   4000.00,
			expectedNetMonthly:   12666.67,
		},
		{
			name:                 "Edge case: 80150",
			employeeName:         "Bob",
			annualSalary:         80150,
			expectedGrossMonthly: 6679.17,
			expectedMonthlyTax:   837.08,
			expectedNetMonthly:   5842.09,
		},
		{
			name:                 "Low salary: 20000",
			employeeName:         "Charlie",
			annualSalary:         20000,
			expectedGrossMonthly: 1666.67,
			expectedMonthlyTax:   0.00,
			expectedNetMonthly:   1666.67,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := GenerateMonthlyPayslip(tt.employeeName, tt.annualSalary)
			
			if result.EmployeeName != tt.employeeName {
				t.Errorf("EmployeeName = %s; want %s", result.EmployeeName, tt.employeeName)
			}
			
			if !floatEquals(result.GrossMonthlyIncome, tt.expectedGrossMonthly, 0.01) {
				t.Errorf("GrossMonthlyIncome = %.2f; want %.2f",
					result.GrossMonthlyIncome, tt.expectedGrossMonthly)
			}
			
			if !floatEquals(result.MonthlyIncomeTax, tt.expectedMonthlyTax, 0.01) {
				t.Errorf("MonthlyIncomeTax = %.2f; want %.2f",
					result.MonthlyIncomeTax, tt.expectedMonthlyTax)
			}
			
			if !floatEquals(result.NetMonthlyIncome, tt.expectedNetMonthly, 0.01) {
				t.Errorf("NetMonthlyIncome = %.2f; want %.2f",
					result.NetMonthlyIncome, tt.expectedNetMonthly)
			}
		})
	}
}

func TestPrintPayslip(t *testing.T) {
	// This is a basic test to ensure PrintPayslip doesn't panic
	// In a real scenario, we might capture stdout to verify output
	data := PayslipData{
		EmployeeName:       "TestUser",
		GrossMonthlyIncome: 5000.00,
		MonthlyIncomeTax:   500.00,
		NetMonthlyIncome:   4500.00,
	}
	
	// Should not panic
	PrintPayslip(data)
}
