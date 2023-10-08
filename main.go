package main

import (
	"github.com/labstack/echo/v4"
	"receipt-processor/cmd/handlers"
)

func main() {
	e := echo.New()
	e.GET("/", handlers.Home)
	e.POST("/receipts/process", handlers.ProcessReceipt)
	e.Logger.Fatal(e.Start(":8080"))
}
