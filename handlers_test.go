package main

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"os"
	"strings"
	"testing"

	"github.com/labstack/echo/v5"
)

func TestHandleGeneratePayslip(t *testing.T) {
	tests := []struct {
		name               string
		requestBody        string
		expectedStatus     int
		expectedEmployeeName string
		expectedGross      string
		expectedTax        string
		expectedNet        string
		expectError        bool
	}{
		{
			name: "Valid request - Ren with 60000",
			requestBody: `{
				"employee_name": "Ren",
				"annual_salary": "60000"
			}`,
			expectedStatus:       http.StatusOK,
			expectedEmployeeName: "Ren",
			expectedGross:        "5000.00",
			expectedTax:          "500.00",
			expectedNet:          "4500.00",
			expectError:          false,
		},
		{
			name: "Valid request - Alice with 200000",
			requestBody: `{
				"employee_name": "Alice",
				"annual_salary": "200000"
			}`,
			expectedStatus:       http.StatusOK,
			expectedEmployeeName: "Alice",
			expectedGross:        "16666.67",
			expectedTax:          "4000.00",
			expectedNet:          "12666.67",
			expectError:          false,
		},
		{
			name: "Missing employee name",
			requestBody: `{
				"annual_salary": "60000"
			}`,
			expectedStatus: http.StatusBadRequest,
			expectError:    true,
		},
		{
			name: "Invalid annual salary - not a number",
			requestBody: `{
				"employee_name": "John",
				"annual_salary": "invalid"
			}`,
			expectedStatus: http.StatusBadRequest,
			expectError:    true,
		},
		{
			name: "Negative annual salary",
			requestBody: `{
				"employee_name": "Jane",
				"annual_salary": "-5000"
			}`,
			expectedStatus: http.StatusBadRequest,
			expectError:    true,
		},
		{
			name:           "Invalid JSON",
			requestBody:    `{invalid json}`,
			expectedStatus: http.StatusBadRequest,
			expectError:    true,
		},
	}

		for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Setup
			e := echo.New()
			req := httptest.NewRequest(http.MethodPost, "/payslip", strings.NewReader(tt.requestBody))
			req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
			rec := httptest.NewRecorder()
			c := e.NewContext(req, rec)

			// Execute
			err := HandleGeneratePayslip(c)

			// Assert
			if err != nil {
				t.Fatalf("Handler returned unexpected error: %v", err)
			}

			if rec.Code != tt.expectedStatus {
				t.Errorf("Expected status %d, got %d", tt.expectedStatus, rec.Code)
			}

			if !tt.expectError {
				var response PayslipResponse
				if err := json.Unmarshal(rec.Body.Bytes(), &response); err != nil {
					t.Fatalf("Failed to unmarshal response: %v", err)
				}

				if response.EmployeeName != tt.expectedEmployeeName {
					t.Errorf("Expected employee_name %s, got %s", tt.expectedEmployeeName, response.EmployeeName)
				}

				if response.GrossMonthlyIncome != tt.expectedGross {
					t.Errorf("Expected gross_monthly_income %s, got %s", tt.expectedGross, response.GrossMonthlyIncome)
				}

				if response.MonthlyIncomeTax != tt.expectedTax {
					t.Errorf("Expected monthly_income_tax %s, got %s", tt.expectedTax, response.MonthlyIncomeTax)
				}

				if response.NetMonthlyIncome != tt.expectedNet {
					t.Errorf("Expected net_monthly_income %s, got %s", tt.expectedNet, response.NetMonthlyIncome)
				}
			} else {
				// Verify error response
				var errorResponse map[string]string
				if err := json.Unmarshal(rec.Body.Bytes(), &errorResponse); err != nil {
					t.Fatalf("Failed to unmarshal error response: %v", err)
				}

				if _, exists := errorResponse["error"]; !exists {
					t.Error("Expected error field in response")
				}
			}
		})
	}
}

func TestHandleGetSalaryComputations(t *testing.T) {
	// Setup test database
	testDBPath := "test_handlers.db"
	os.Remove(testDBPath)
	defer os.Remove(testDBPath)

	err := InitDB(testDBPath)
	if err != nil {
		t.Fatalf("Failed to initialize test database: %v", err)
	}
	defer CloseDB()

	// Add some test data
	SaveSalaryComputation("John Doe", "60000", "500.00")
	SaveSalaryComputation("Jane Smith", "80000", "837.08")

	// Test GET request
	e := echo.New()
	req := httptest.NewRequest(http.MethodGet, "/payslips", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	// Execute
	err = HandleGetSalaryComputations(c)
	if err != nil {
		t.Fatalf("Handler returned unexpected error: %v", err)
	}

	// Assert status code
	if rec.Code != http.StatusOK {
		t.Errorf("Expected status %d, got %d", http.StatusOK, rec.Code)
	}

	// Parse response
	var response SalaryComputationsResponse
	if err := json.Unmarshal(rec.Body.Bytes(), &response); err != nil {
		t.Fatalf("Failed to unmarshal response: %v", err)
	}

	// Verify we have 2 records
	if len(response.SalaryComputations) != 2 {
		t.Errorf("Expected 2 records, got %d", len(response.SalaryComputations))
	}

	// Verify records exist
	foundJohn := false
	foundJane := false
	for _, record := range response.SalaryComputations {
		if record.EmployeeName == "John Doe" && record.AnnualSalary == "60000" && record.MonthlyIncomeTax == "500.00" {
			foundJohn = true
		}
		if record.EmployeeName == "Jane Smith" && record.AnnualSalary == "80000" && record.MonthlyIncomeTax == "837.08" {
			foundJane = true
		}
		// Verify timestamp exists
		if record.Timestamp == "" {
			t.Error("Expected timestamp to be non-empty")
		}
	}

	if !foundJohn {
		t.Error("Expected to find John Doe record")
	}
	if !foundJane {
		t.Error("Expected to find Jane Smith record")
	}
}
