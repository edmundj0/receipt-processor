package pkg

import (
	"strings"
	"math"
	"strconv"
	"unicode"
	"fmt"
)

type Receipt struct {
	ID string `json:"id"`
	Retailer string `json:"retailer"`
	PurchaseDate string `json:"purchaseDate"`
	PurchaseTime string `json:"purchaseTime"`
	Total string `json:"total"`
	Items []map[string]interface{} `json:"items"`
	Points int `json:"points"`
}

func CalculatePoints(receipt *Receipt) int {
	points := 0

	// One point for every alphanumeric character in the retailer name
	for _, char := range receipt.Retailer {
		if unicode.IsLetter(char) || unicode.IsDigit(char) {
			points++
		}
	}

	fmt.Println(points,"alphanumeric")

	// 50 points if the total is a round dollar amount with no cents
	if strings.HasSuffix(receipt.Total, ".00") {
		points += 50
		fmt.Println("50 for no cents")
	}

	// 25 points if the total is a multiple of 0.25
	totalFloat, _ := strconv.ParseFloat(receipt.Total, 64)
	if math.Mod(totalFloat, 0.25) == 0 {
		points += 25
		fmt.Println("25 for multiple of 0.25")
	}

	// 5 points for every two items on the receipt
	count := len(receipt.Items)
	fmt.Println("two items", (count/2) * 5)
	points += (count / 2) * 5

	// If the trimmed length of the item description is a multiple of 3, multiply the price by 0.2 and round up to the nearest integer. The result is the number of points earned
	for _, item := range receipt.Items {
		description := strings.TrimSpace(item["shortDescription"].(string))
		price, _ := strconv.ParseFloat(item["price"].(string), 64)
		if (len(description)%3 == 0){
			points += int(math.Ceil(price * 0.2))
		}
	}

	// 6 points if the day in the purchase date is odd
	lastTwoDigits := receipt.PurchaseDate[len(receipt.PurchaseDate)-2:]
	dayInt, err := strconv.Atoi(lastTwoDigits)
	if err == nil && dayInt%2 != 0 {
		points += 6
	}

	// 10 points if the time of purchase is after 2:00pm and before 4:00pm
	timeStr := receipt.PurchaseTime
	timeParts := strings.Split(timeStr, ":")
	hours, err := strconv.Atoi(timeParts[0])
	minutes, err := strconv.Atoi(timeParts[1])

	timeInt := hours * 100 + minutes

	if timeInt > 1400 && timeInt < 1600 {
		points += 10
		fmt.Println("purchase time")
	}

	return points

}
