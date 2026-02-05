package main

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v5"
)

// PayslipRequest represents the request body for POST /payslip
type PayslipRequest struct {
	EmployeeName  string `json:"employee_name"`
	AnnualSalary  string `json:"annual_salary"`
}

// PayslipResponse represents the response body for POST /payslip
type PayslipResponse struct {
	EmployeeName       string `json:"employee_name"`
	GrossMonthlyIncome string `json:"gross_monthly_income"`
	MonthlyIncomeTax   string `json:"monthly_income_tax"`
	NetMonthlyIncome   string `json:"net_monthly_income"`
}

// HandleGeneratePayslip handles POST /payslip endpoint
func HandleGeneratePayslip(c *echo.Context) error {
	var req PayslipRequest
	
	if err := (*c).Bind(&req); err != nil {
		return (*c).JSON(http.StatusBadRequest, map[string]string{
			"error": "Invalid request body",
		})
	}
	
	// Validate employee name
	if req.EmployeeName == "" {
		return (*c).JSON(http.StatusBadRequest, map[string]string{
			"error": "employee_name is required",
		})
	}
	
	// Parse annual salary
	annualSalary, err := strconv.ParseFloat(req.AnnualSalary, 64)
	if err != nil {
		return (*c).JSON(http.StatusBadRequest, map[string]string{
			"error": "annual_salary must be a valid number",
		})
	}
	
	if annualSalary < 0 {
		return (*c).JSON(http.StatusBadRequest, map[string]string{
			"error": "annual_salary must be non-negative",
		})
	}
	
	// Generate payslip
	payslip := GenerateMonthlyPayslip(req.EmployeeName, annualSalary)
	
	// Format response with 2 decimal places
	response := PayslipResponse{
		EmployeeName:       payslip.EmployeeName,
		GrossMonthlyIncome: fmt.Sprintf("%.2f", payslip.GrossMonthlyIncome),
		MonthlyIncomeTax:   fmt.Sprintf("%.2f", payslip.MonthlyIncomeTax),
		NetMonthlyIncome:   fmt.Sprintf("%.2f", payslip.NetMonthlyIncome),
	}
	
	// Save to database
	err = SaveSalaryComputation(
		req.EmployeeName,
		req.AnnualSalary,
		response.MonthlyIncomeTax,
	)
	if err != nil {
		// Log error but don't fail the request
		(*c).Logger().Error("Failed to save to database", "error", err)
	}
	
	return (*c).JSON(http.StatusOK, response)
}
