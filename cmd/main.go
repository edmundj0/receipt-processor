package main

import (
	"github.com/labstack/echo/v4"
	"receipt-processor/handlers"
)

func main() {
	e := echo.New()
	e.GET("/", handlers.Home)
	e.GET("/receipts/:id/points", handlers.GetReceiptPoints)
	e.POST("/receipts/process", handlers.ProcessReceipt)
	e.Logger.Fatal(e.Start(":8080"))
}
