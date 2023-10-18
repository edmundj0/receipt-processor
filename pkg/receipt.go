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






// func CalculatePoints(receipt *Receipt) (int, error) {
//     points := 0

//     // One point for every alphanumeric character in the retailer name
//     for _, char := range receipt.Retailer {
//         if unicode.IsLetter(char) || unicode.IsDigit(char) {
//             points++
//         }
//     }

//     // 50 points if the total is a round dollar amount with no cents
//     if strings.HasSuffix(receipt.Total, ".00") {
//         points += 50
//     }

//     // 25 points if the total is a multiple of 0.25
//     totalFloat, err := strconv.ParseFloat(receipt.Total, 64)
//     if err != nil {
//         return 0, err
//     }
//     if math.Mod(totalFloat, 0.25) == 0 {
//         points += 25
//     }

//     // 5 points for every two items on the receipt
//     count := len(receipt.Items)
//     points += (count / 2) * 5

//     // If the trimmed length of the item description is a multiple of 3, multiply the price by 0.2 and round up to the nearest integer. The result is the number of points earned
//     for _, item := range receipt.Items {
//         description := strings.TrimSpace(item["shortDescription"].(string))
//         price, err := strconv.ParseFloat(item["price"].(string), 64)
//         if err != nil {
//             return 0, err
//         }
//         if (len(description)%3 == 0) {
//             points += int(math.Ceil(price * 0.2))
//         }
//     }

//     // 6 points if the day in the purchase date is odd
//     lastTwoDigits := receipt.PurchaseDate[len(receipt.PurchaseDate)-2:]
//     dayInt, err := strconv.Atoi(lastTwoDigits)
//     if err != nil {
//         return 0, err
//     }
//     if dayInt%2 != 0 {
//         points += 6
//     }

//     // 10 points if the time of purchase is after 2:00pm and before 4:00pm
//     timeStr := receipt.PurchaseTime
//     timeParts := strings.Split(timeStr, ":")
//     hours, err := strconv.Atoi(timeParts[0])
//     if err != nil {
//         return 0, err
//     }
//     minutes, err := strconv.Atoi(timeParts[1])
//     if err != nil {
//         return 0, err // Return an error if parsing fails
//     }

//     timeInt := hours * 100 + minutes

//     if timeInt > 1400 && timeInt < 1600 {
//         points += 10
//     }

//     return points, nil
// }
