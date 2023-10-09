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
            points, err := pkg.CalculatePoints(test.receipt)
            if err != nil{
                t.Errorf("CalculatePoints Error")
            }

            if points != test.expectedPoints {
                t.Errorf("Expected %d points, but got %d", test.expectedPoints, points)
            }
        })
    }
}
