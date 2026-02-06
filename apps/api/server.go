package main

import (
	"log"
	"net/http"

	"github.com/labstack/echo/v5"
	"github.com/labstack/echo/v5/middleware"
)

func main() {
	// Initialize database with absolute path in the api directory
	if err := InitDB("./payslip.db"); err != nil {
		log.Fatalf("Failed to initialize database: %v", err)
	}
	defer CloseDB()

	e := echo.New()
	e.Use(middleware.RequestLogger())

	e.GET("/", func(c *echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})

	// Payslip generation endpoint
	e.POST("/payslip", func(c *echo.Context) error {
		return HandleGeneratePayslip(c)
	})

	// Get all salary computations
	e.GET("/payslips", func(c *echo.Context) error {
		return HandleGetSalaryComputations(c)
	})

	if err := e.Start(":1323"); err != nil {
		e.Logger.Error("failed to start server", "error", err)
	}
}
