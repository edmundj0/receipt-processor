package test

import (
    "testing"
    "receipt-processor/pkg"
)

func TestCalculatePoints(t *testing.T) {
    tests := []struct {
        receipt        *pkg.Receipt
        expectedPoints int
    }{
		{
			receipt: &pkg.Receipt{
				Retailer: "Target",
				PurchaseDate: "2022-01-01",
				PurchaseTime: "13:01",
				Total: "35.35",
				Items: []map[string]interface{}{
					{"shortDescription": "Mountain Dew 12PK", "price": "6.49"},
					{"shortDescription": "Emils Cheese Pizza", "price": "12.25"},
					{"shortDescription": "Knorr Creamy Chicken", "price": "1.26"},
					{"shortDescription": "Doritos Nacho Cheese", "price": "3.35"},
					{"shortDescription": "Klarbrunn 12-PK 12 FL OZ", "price": "12.00"},
				},
			},
			expectedPoints: 28,
		},
		{
			receipt: &pkg.Receipt{
				Retailer: "M&M Corner Market",
				PurchaseDate: "2022-03-20",
				PurchaseTime: "14:33",
				Total: "9.00",
				Items: []map[string]interface{}{
					{"shortDescription": "Gatorade", "price": "2.25"},
					{"shortDescription": "Gatorade", "price": "2.25"},
					{"shortDescription": "Gatorade", "price": "2.25"},
					{"shortDescription": "Gatorade", "price": "2.25"},
				},
			},
			expectedPoints: 109,
		},
        {
            receipt: &pkg.Receipt{
                Retailer:     "Test&&Retailer",
                PurchaseDate: "2022-05-05",
                PurchaseTime: "14:01",
                Total:        "100.00",
                Items: []map[string]interface{}{
                    {"shortDescription": "TestItem", "price": "50.00"},
                    {"shortDescription": "TestItem2", "price": "30.00"},
                },
            },
            expectedPoints: 114,
        },
        {
            receipt: &pkg.Receipt{
                Retailer:     "A n o t h e r R e t a i l e r",
                PurchaseDate: "2022-06-06",
                PurchaseTime: "15:59",
                Total:        "50.75",
                Items: []map[string]interface{}{
                    {"shortDescription": "abcdef", "price": "50.00"},
                    {"shortDescription": "Item", "price": "0.75"},
					{"shortDescription": "Item2", "price": "0.00"},
                },
            },
            expectedPoints: 65,
        },
    }

    for _, test := range tests {
        t.Run("", func(t *testing.T) {
            err := pkg.CalculatePoints(test.receipt)
            if err != nil{
                t.Errorf("CalculatePoints Error")
            }

            if test.receipt.Points != test.expectedPoints {
                t.Errorf("Expected %d points, but got %d", test.expectedPoints, test.receipt.Points)
            }
        })
    }
}

func TestIsValidRetailer(t *testing.T) {
    validRetailer := "Target"
    invalidRetailer := "   Invalid Retailer   "

    if !pkg.IsValidRetailer(validRetailer) {
        t.Errorf("Expected %s to be a valid retailer", validRetailer)
    }

    if pkg.IsValidRetailer(invalidRetailer) {
        t.Errorf("Expected %s to be an invalid retailer", invalidRetailer)
    }
}

func TestIsValidDate(t *testing.T) {
    validDate := "2022-01-01"
    invalidDate := "20221301"

    if !pkg.IsValidDate(validDate) {
        t.Errorf("Expected %s to be a valid date", validDate)
    }

    if pkg.IsValidDate(invalidDate) {
        t.Errorf("Expected %s to be an invalid date", invalidDate)
    }
}

func TestIsValidPurchaseTime(t *testing.T) {
    validTime := "13:01"
    invalidTime := "25:01" // Invalid hour

    if !pkg.IsValidPurchaseTime(validTime) {
        t.Errorf("Expected %s to be a valid purchase time", validTime)
    }

    if pkg.IsValidPurchaseTime(invalidTime) {
        t.Errorf("Expected %s to be an invalid purchase time", invalidTime)
    }
}

func TestIsValidTotal(t *testing.T) {
    validTotal := "35.35"
    invalidTotal := "abc" // Invalid format

    if !pkg.IsValidTotal(validTotal) {
        t.Errorf("Expected %s to be a valid total amount", validTotal)
    }

    if pkg.IsValidTotal(invalidTotal) {
        t.Errorf("Expected %s to be an invalid total amount", invalidTotal)
    }
}

func TestIsValidItems(t *testing.T) {
    validItems := []map[string]interface{}{
    {
        "shortDescription": "Mountain Dew 12PK",
        "price": "6.49",
    },
    {
        "shortDescription": "            ",
        "price": "234346593459589345834583458.34",
    },
    }

    invalidItems := []map[string]interface{}{
        {
            "shortDescription": "Invalid Item",
            // Missing "price" field
        },
        {
            // Missing "shortDescription" field
            "price": "232323232323232232323.20",
        },
        {
            "shortDescription": " ",
            "price": "$23.50", // Invalid price
        },
        {
            "shortDescription": "Invalid Item",
            "price": "33.333", // Invalid price
        },
    }

    for _, item := range validItems {
        if !pkg.IsValidItems(item) {
            t.Errorf("Expected item to be valid: %v", item)
        }
    }

    for _, item := range invalidItems {
        if pkg.IsValidItems(item) {
            t.Errorf("Expected item to be invalid: %v", item)
        }
    }
}
