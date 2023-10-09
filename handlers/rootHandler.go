package handlers

import (
	"net/http"
	"github.com/labstack/echo/v4"
	"github.com/google/uuid"
	"receipt-processor/pkg"
)



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

	receipt := new(pkg.Receipt)
	if err := c.Bind(receipt); err != nil {
		return c.JSON(http.StatusBadRequest, "Invalid JSON")
	}

	receiptID := uuid.New().String()
	receipt.ID = receiptID

	points, err := pkg.CalculatePoints(receipt)
    if err != nil {
        return c.JSON(http.StatusInternalServerError, map[string]string{
            "error": err.Error(),
        })
    }

	receipt.Points = points

	receiptStore[receiptID] = *receipt
	pointsStore[receiptID] = receipt.Points

	response := pkg.Response{ID: receiptID}
	return c.JSON(http.StatusOK, response)
}
