package pkg

import (
	"errors"
	"regexp"
	// "fmt"
)

func ValidateReceipt(receipt *Receipt) error {

	if len(receipt.Retailer) == 0 {
		return errors.New("Retailer field is required")
	}
	if !isValidRetailer(receipt.Retailer) {
        return errors.New("Invalid retailer name")
    }


	if len(receipt.PurchaseDate) == 0 {
        return errors.New("PurchaseDate field is required")
    }
    if !isValidDate(receipt.PurchaseDate) {
        return errors.New("Invalid purchase date")
    }


	if len(receipt.PurchaseTime) == 0 {
		return errors.New("PurchaseTime field is required")
	}
	if !isValidPurchaseTime(receipt.PurchaseTime) {
        return errors.New("Invalid purchase time")
    }

	if len(receipt.Total) == 0 {
		return errors.New("Total field is required")
	}
	if !isValidTotal(receipt.Total) {
        return errors.New("Invalid total amount")
    }


	if len(receipt.Items) == 0 {
		return errors.New("Items field is required")
	}
	for _, item := range receipt.Items {
		if !isValidItems(item) {
			return errors.New("Invalid item")
		}
	}

	return nil

}



func isValidRetailer(retailer string) bool {
	return regexp.MustCompile(`^\S+$`).MatchString(retailer)
}

func isValidDate(date string) bool {
	return regexp.MustCompile(`^\d{4}-\d{2}-\d{2}$`).MatchString(date)
}

func isValidPurchaseTime(time string) bool {
	return regexp.MustCompile(`^([01]\d|2[0-3]):([0-5]\d)$`).MatchString(time)
}

func isValidTotal(total string) bool {
    return regexp.MustCompile(`^\d+\.\d{2}$`).MatchString(total)
}

func isValidItems(item map[string]interface{}) bool {
	// check required fields
	requiredFields := []string{"shortDescription", "price"}
	for _, field := range requiredFields {
		if _, exists := item[field]; !exists {
			return false
		}
	}

	shortDescription, ok := item["shortDescription"].(string)
	// fmt.Println(shortDescription, ok)
	if !ok || !regexp.MustCompile(`^[\w\s\-]+$`).MatchString(shortDescription) {
		// fmt.Println("test")
        return false
    }

	price, ok := item["price"].(string)
    if !ok || !regexp.MustCompile(`^\d+\.\d{2}$`).MatchString(price) {
        return false
    }

	return true
}
