package main

import (
	"database/sql"
	"log"
	"time"

	_ "modernc.org/sqlite"
)

var db *sql.DB

// InitDB initializes the database connection and creates tables
func InitDB(dbPath string) error {
	var err error
	db, err = sql.Open("sqlite", dbPath)
	if err != nil {
		return err
	}

	// Create salary_computations table
	createTableSQL := `
	CREATE TABLE IF NOT EXISTS salary_computations (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		timestamp TEXT NOT NULL,
		employee_name TEXT NOT NULL,
		annual_salary TEXT NOT NULL,
		monthly_income_tax TEXT NOT NULL
	);
	`

	_, err = db.Exec(createTableSQL)
	if err != nil {
		return err
	}

	log.Println("Database initialized successfully")
	return nil
}

// SaveSalaryComputation saves a salary computation record to the database
func SaveSalaryComputation(employeeName, annualSalary, monthlyIncomeTax string) error {
	timestamp := time.Now().Format(time.RFC3339)

	insertSQL := `
	INSERT INTO salary_computations (timestamp, employee_name, annual_salary, monthly_income_tax)
	VALUES (?, ?, ?, ?)
	`

	_, err := db.Exec(insertSQL, timestamp, employeeName, annualSalary, monthlyIncomeTax)
	return err
}

// SalaryComputationRecord represents a record from the database
type SalaryComputationRecord struct {
	Timestamp        string `json:"time_stamp"`
	EmployeeName     string `json:"employee_name"`
	AnnualSalary     string `json:"annual_salary"`
	MonthlyIncomeTax string `json:"monthly_income_tax"`
}

// GetAllSalaryComputations retrieves all salary computation records
func GetAllSalaryComputations() ([]SalaryComputationRecord, error) {
	rows, err := db.Query(`
		SELECT timestamp, employee_name, annual_salary, monthly_income_tax
		FROM salary_computations
		ORDER BY timestamp DESC
	`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var records []SalaryComputationRecord
	for rows.Next() {
		var record SalaryComputationRecord
		err := rows.Scan(
			&record.Timestamp,
			&record.EmployeeName,
			&record.AnnualSalary,
			&record.MonthlyIncomeTax,
		)
		if err != nil {
			return nil, err
		}
		records = append(records, record)
	}

	return records, nil
}

// CloseDB closes the database connection
func CloseDB() error {
	if db != nil {
		return db.Close()
	}
	return nil
}
