package handlers

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/google/uuid"
)

type Receipt struct {
	ID string `json:"id"`
	Retailer string `json:"retailer"`
	PurchaseDate string `json:"purchaseDate"`
	PurchaseTime string `json:"purchaseTime"`
	Total string `json:"total"`
	Items []map[string]interface{} `json:"items"`
}

type Response struct {
	ID string `json:"id"`
}


var receiptStore = make(map[string]Receipt)


func Home(c echo.Context) error {
	receipts := make([]Receipt, 0, len(receiptStore))
	for _, receipt := range receiptStore {
		receipts = append(receipts, receipt)
	}

	return c.JSON(http.StatusOK, receipts)


	return c.String(http.StatusOK, "")
}

func ProcessReceipt(c echo.Context) error {
	// Parse json into Receipt struct
	receipt := new(Receipt)
	if err := c.Bind(receipt); err != nil {
		return c.JSON(http.StatusBadRequest, "Invalid JSON")
	}

	receiptID := uuid.New().String()
	receipt.ID = receiptID
	receiptStore[receiptID] = *receipt

	response := Response{ID: receiptID}
	return c.JSON(http.StatusOK, response)
}
