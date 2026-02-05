# Payslip Generator

A REST API service for generating monthly payslips with progressive tax computation.

## Features

- **Extensible Tax Bracket System**: Easily configurable tax brackets for different salary ranges
- **Monthly Payslip Generation**: Computes gross monthly income, monthly tax, and net income
- **REST API Endpoints**: POST endpoint for payslip generation and GET endpoint for records retrieval
- **SQLite Database**: Persists all salary computations with timestamps
- **Comprehensive Test Coverage**: Unit tests for all components

## Tax Brackets

| Salary Bracket | Rate |
|----------------|------|
| 0 - 20,000 | 0% |
| 20,001 - 40,000 | 10% |
| 40,001 - 80,000 | 20% |
| 80,001 - 180,000 | 30% |
| 180,001 and above | 40% |

## API Endpoints

### POST /payslip
Generate a monthly payslip for an employee.

**Request Body:**
```json
{
  "employee_name": "string",
  "annual_salary": "string"
}
```

**Response:**
```json
{
  "employee_name": "string",
  "gross_monthly_income": "string",
  "monthly_income_tax": "string",
  "net_monthly_income": "string"
}
```

**Example:**
```bash
curl -X POST http://localhost:1323/payslip \
  -H "Content-Type: application/json" \
  -d '{"employee_name":"Ren","annual_salary":"60000"}'
```

**Example Response:**
```json
{
  "employee_name": "Ren",
  "gross_monthly_income": "5000.00",
  "monthly_income_tax": "500.00",
  "net_monthly_income": "4500.00"
}
```

### GET /payslips
Retrieve all salary computation records.

**Response:**
```json
{
  "salary_computations": [
    {
      "time_stamp": "string",
      "employee_name": "string",
      "annual_salary": "string",
      "monthly_income_tax": "string"
    }
  ]
}
```

**Example:**
```bash
curl http://localhost:1323/payslips
```

## Installation & Running

### Prerequisites
- Go 1.25.7 or higher

### Install Dependencies
```bash
go mod download
```

### Run the Server
```bash
go run server.go tax.go payslip.go handlers.go database.go
```

The server will start on `http://localhost:1323`

### Run Tests
```bash
go test -v ./...
```

## Project Structure

```
.
├── server.go           # Main server and initialization
├── tax.go              # Tax bracket structure and computation logic
├── tax_test.go         # Tests for tax computation
├── payslip.go          # Payslip generation functions
├── payslip_test.go     # Tests for payslip generation
├── handlers.go         # HTTP request handlers
├── handlers_test.go    # Tests for HTTP handlers
├── database.go         # Database operations
├── database_test.go    # Tests for database operations
├── go.mod              # Go module dependencies
└── go.sum              # Dependency checksums
```

## Design Decisions

### Strategy Pattern Implementation
The project implements the **Strategy Pattern** for tax calculation (from [OODesign.com](https://www.oodesign.com/strategy-pattern.html)). This design pattern allows the tax calculation algorithm to be selected at runtime and makes it easy to add new tax calculation strategies without modifying existing code.

**Benefits:**
- **Extensibility**: Easy to add new tax strategies (e.g., regional taxes, flat taxes)
- **Testability**: Each strategy can be tested independently
- **Flexibility**: Switch between strategies at runtime
- **Open/Closed Principle**: Open for extension, closed for modification

**Implementation:**
- `TaxStrategy` interface defines the contract for all tax strategies
- `ProgressiveTaxStrategy` implements the current progressive tax system
- `FlatTaxStrategy` demonstrates an alternative simple tax calculation
- `GenerateMonthlyPayslipWithStrategy()` allows using any strategy

**Example:**
```go
// Use default progressive tax
payslip := GenerateMonthlyPayslip("Alice", 60000)

// Use a custom flat tax strategy
flatStrategy := NewFlatTaxStrategy(0.15) // 15% flat tax
payslip := GenerateMonthlyPayslipWithStrategy("Bob", 60000, flatStrategy)
```

### Extensible Tax Bracket System
The tax computation uses a configurable slice of `TaxBracket` structs, making it easy to modify tax rates without changing the core logic. Simply update the `TaxBrackets` variable in `tax.go` to change the tax structure.

### Progressive Tax Calculation
The `ComputeAnnualTax` function implements true progressive taxation, where income is taxed at different rates for different portions (brackets) of the total salary.

### Database Persistence
Uses SQLite for simplicity and portability. The database is automatically initialized on server startup and stores all computation records with RFC3339 formatted timestamps.

## Example Computations

### Annual Salary: $60,000
- First $20,000 at 0% = $0
- Next $20,000 at 10% = $2,000
- Next $20,000 at 20% = $4,000
- **Total Annual Tax: $6,000**
- **Monthly Tax: $500**
- **Gross Monthly: $5,000**
- **Net Monthly: $4,500**

### Annual Salary: $200,000
- First $20,000 at 0% = $0
- Next $20,000 at 10% = $2,000
- Next $40,000 at 20% = $8,000
- Next $100,000 at 30% = $30,000
- Next $20,000 at 40% = $8,000
- **Total Annual Tax: $48,000**
- **Monthly Tax: $4,000**
- **Gross Monthly: $16,666.67**
- **Net Monthly: $12,666.67**

## License

MIT
