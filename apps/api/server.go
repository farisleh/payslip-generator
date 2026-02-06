package main

import (
	"log"
	"net/http"
	"os"

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

	// CORS configuration
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"http://localhost:5173", "https://farisleh.github.io"},
		AllowMethods: []string{http.MethodGet, http.MethodPost, http.MethodPut, http.MethodDelete},
		AllowHeaders: []string{echo.HeaderOrigin, echo.HeaderContentType, echo.HeaderAccept},
	}))

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

	// Export salary computations as CSV
	e.GET("/payslips/export", func(c *echo.Context) error {
		return HandleExportCSV(c)
	})

	// Get port from environment or use default
	port := os.Getenv("PORT")
	if port == "" {
		port = "1323"
	}

	if err := e.Start(":" + port); err != nil {
		e.Logger.Error("failed to start server", "error", err)
	}
}
