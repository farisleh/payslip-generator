package main

import (
	"fmt"
)

// PayslipData holds the computed payslip information
type PayslipData struct {
	EmployeeName       string
	GrossMonthlyIncome float64
	MonthlyIncomeTax   float64
	NetMonthlyIncome   float64
}

// GenerateMonthlyPayslip computes monthly payslip data based on annual salary
func GenerateMonthlyPayslip(employeeName string, annualSalary float64) PayslipData {
	return GenerateMonthlyPayslipWithStrategy(employeeName, annualSalary, DefaultTaxStrategy)
}

// GenerateMonthlyPayslipWithStrategy computes monthly payslip using a specific tax strategy
// This demonstrates the Strategy Pattern in action, allowing different tax calculations
func GenerateMonthlyPayslipWithStrategy(employeeName string, annualSalary float64, strategy TaxStrategy) PayslipData {
	// Calculate annual tax using the provided strategy
	annualTax := strategy.CalculateAnnualTax(annualSalary)

	// Convert to monthly values
	grossMonthlyIncome := annualSalary / 12.0
	monthlyIncomeTax := annualTax / 12.0
	netMonthlyIncome := grossMonthlyIncome - monthlyIncomeTax

	return PayslipData{
		EmployeeName:       employeeName,
		GrossMonthlyIncome: grossMonthlyIncome,
		MonthlyIncomeTax:   monthlyIncomeTax,
		NetMonthlyIncome:   netMonthlyIncome,
	}
}

// PrintPayslip prints the payslip in the required format
func PrintPayslip(data PayslipData) {
	fmt.Printf("Monthly Payslip for: %q\n", data.EmployeeName)
	fmt.Printf("Gross Monthly Income: $%.2f\n", data.GrossMonthlyIncome)
	fmt.Printf("Monthly Income Tax: $%.2f\n", data.MonthlyIncomeTax)
	fmt.Printf("Net Monthly Income: $%.2f\n", data.NetMonthlyIncome)
}
