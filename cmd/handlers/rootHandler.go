package handlers

import (
	"net/http"
	"fmt"
	"github.com/labstack/echo/v4"
	"github.com/google/uuid"
	"receipt-processor/pkg"
)

type Response struct {
	ID string `json:"id"`
}


var receiptStore = make(map[string]pkg.Receipt)
var pointsStore = make(map[string]int)


func Home(c echo.Context) error {
	receipts := make([]pkg.Receipt, 0, len(receiptStore))
	for _, receipt := range receiptStore {
		receipts = append(receipts, receipt)
	}

	return c.JSON(http.StatusOK, receipts)


	return c.String(http.StatusOK, "")
}

func GetReceiptPoints(c echo.Context) error {
	// receiptID from url parameter
	receiptID := c.Param("id")

	points, exists := pointsStore[receiptID]

	if !exists {
		return c.JSON(http.StatusNotFound, "Receipt Not Found")
	}

	response := map[string]interface{}{"points": points}
	return c.JSON(http.StatusOK, response)
}



func ProcessReceipt(c echo.Context) error {
	// Parse json into Receipt struct
	receipt := new(pkg.Receipt)
	if err := c.Bind(receipt); err != nil {
		return c.JSON(http.StatusBadRequest, "Invalid JSON")
	}

	receiptID := uuid.New().String()
	receipt.ID = receiptID

	receipt.Points = pkg.CalculatePoints(receipt)
	fmt.Println(receipt)

	receiptStore[receiptID] = *receipt
	pointsStore[receiptID] = receipt.Points

	response := Response{ID: receiptID}
	return c.JSON(http.StatusOK, response)
}
