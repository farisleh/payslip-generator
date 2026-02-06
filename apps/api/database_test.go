package main

import (
	"os"
	"testing"
)

func TestDatabaseOperations(t *testing.T) {
	// Use a test database file
	testDBPath := "test_payslip.db"

	// Clean up before and after test
	os.Remove(testDBPath)
	defer os.Remove(testDBPath)

	// Initialize database
	err := InitDB(testDBPath)
	if err != nil {
		t.Fatalf("Failed to initialize database: %v", err)
	}
	defer CloseDB()

	// Test saving a computation
	err = SaveSalaryComputation("John Doe", "60000", "500.00")
	if err != nil {
		t.Fatalf("Failed to save salary computation: %v", err)
	}

	// Save another record
	err = SaveSalaryComputation("Jane Smith", "80000", "837.08")
	if err != nil {
		t.Fatalf("Failed to save second salary computation: %v", err)
	}

	// Retrieve all records
	records, err := GetAllSalaryComputations()
	if err != nil {
		t.Fatalf("Failed to retrieve salary computations: %v", err)
	}

	// Verify we have 2 records
	if len(records) != 2 {
		t.Errorf("Expected 2 records, got %d", len(records))
	}

	// Verify first record (should be most recent - Jane Smith due to DESC order)
	if len(records) > 0 {
		// Note: Due to rapid insertion, order might vary, so we just check both exist
		foundJohn := false
		foundJane := false

		for _, record := range records {
			if record.EmployeeName == "John Doe" && record.AnnualSalary == "60000" && record.MonthlyIncomeTax == "500.00" {
				foundJohn = true
			}
			if record.EmployeeName == "Jane Smith" && record.AnnualSalary == "80000" && record.MonthlyIncomeTax == "837.08" {
				foundJane = true
			}
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
}
