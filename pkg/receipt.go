package pkg

import (
	"strings"
	"math"
	"strconv"
	"unicode"
)

type Response struct {
	ID string `json:"id"`
}

type Receipt struct {
	ID string `json:"id"`
	Retailer string `json:"retailer"`
	PurchaseDate string `json:"purchaseDate"`
	PurchaseTime string `json:"purchaseTime"`
	Total string `json:"total"`
	Items []map[string]interface{} `json:"items"`
	Points int `json:"points"`
}


func CalculateAlphaNumericInRetailerName(receipt *Receipt) {

    for _, char := range receipt.Retailer {
        if unicode.IsLetter(char) || unicode.IsDigit(char) {
            receipt.Points++
        }
    }

}

func CalculateIfTotalIsRoundDollar(receipt *Receipt) {

    if strings.HasSuffix(receipt.Total, ".00") {
        receipt.Points += 50
    }

}

func CalculateIfTotalMultipleOf25(receipt *Receipt) error {
    totalFloat, err := strconv.ParseFloat(receipt.Total, 64)
    if err != nil {
        return err
    }

    if math.Mod(totalFloat, 0.25) == 0 {
        receipt.Points += 25
    }

    return nil
}


func CalcEveryTwoItemsOnReceipt(receipt *Receipt) {

    count := len(receipt.Items)
    receipt.Points += (count / 2) * 5

}


func CalcTrimmedLengthOfItemDesc(receipt *Receipt) error {

    for _, item := range receipt.Items {
        description := strings.TrimSpace(item["shortDescription"].(string))
        price, err := strconv.ParseFloat(item["price"].(string), 64)
        if err != nil {
            return err
        }
        if (len(description)%3 == 0) {
            receipt.Points += int(math.Ceil(price * 0.2))
        }
    }

    return nil
}

func CalcIsPurchaseDateOdd(receipt *Receipt) error {

    lastTwoDigits := receipt.PurchaseDate[len(receipt.PurchaseDate)-2:]
    dayInt, err := strconv.Atoi(lastTwoDigits)
    if err != nil {
        return err
    }
    if dayInt%2 != 0 {
        receipt.Points += 6
    }

    return nil
}


func CalculateTimeOfPurchase(receipt *Receipt) error {

    timeStr := receipt.PurchaseTime
    timeParts := strings.Split(timeStr, ":")
    hours, err := strconv.Atoi(timeParts[0])
    if err != nil {
        return err
    }
    minutes, err := strconv.Atoi(timeParts[1])
    if err != nil {
        return err
    }

    timeInt := hours * 100 + minutes

    if timeInt > 1400 && timeInt < 1600 {
        receipt.Points += 10
    }

    return nil
}

func CalculatePoints(receipt *Receipt) error {
    CalculateAlphaNumericInRetailerName(receipt)
    CalculateIfTotalIsRoundDollar(receipt)
    if err := CalculateIfTotalMultipleOf25(receipt); err != nil {
        return err
    }
    CalcEveryTwoItemsOnReceipt(receipt)
    if err := CalcTrimmedLengthOfItemDesc(receipt); err != nil {
        return err
    }
    if err := CalcIsPurchaseDateOdd(receipt); err != nil {
        return err
    }
    if err := CalculateTimeOfPurchase(receipt); err != nil {
        return err
    }

    return nil
}

