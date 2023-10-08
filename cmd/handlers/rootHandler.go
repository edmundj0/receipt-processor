package handlers

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/google/uuid"
)

type Receipt struct {
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
	return c.String(http.StatusOK, "test")
}

func ProcessReceipt(c echo.Context) error {
	// Parse json into Receipt struct
	receipt := new(Receipt)
	if err := c.Bind(receipt); err != nil {
		return c.JSON(http.StatusBadRequest, "Invalid JSON")
	}

	receiptID := uuid.New().String()
	receiptStore[receiptID] = *receipt

	response := Response{ID: receiptID}
	return c.JSON(http.StatusOK, response)
}
