# Receipt Processor Service

## Overview

The Receipt Processor Service is a Go-based application that processes receipts and calculates the amount of points awarded based on particular rules. It provides two API endpoints for processing receipts and retrieving points for a receipt.

- Endpoint: Process Receipts
  - Path: `/receipts/process`
  - Method: POST
  - Payload: Receipt JSON
  - Response: JSON containing an ID for the receipt

- Endpoint: Get Points
  - Path: `/receipts/{id}/points`
  - Method: GET
  - Response: A JSON object containing the number of points awarded for a particular receipt

- Endpoint: All Receipts
  - Path `/`
  - Method: GET
  - Response: An Array of all receipts in the database


 ## Prerequisites

- Containerized Deployment: Docker [Docker Installation Documentation](https://docs.docker.com/engine/install/)
- For local development: Go (Golang) installed on your system [Go Installation Documentation](https://go.dev/doc/install)

## Getting Started

To run the Receipt Processor Service locally, follow these steps:

1. Clone the repository
```
git clone https://github.com/edmundj0/receipt-processor.git && cd receipt-processor
```
2. Start the service with Docker
```
docker compose up
```
The service will be accessible at http://localhost:8080


## Usage

1. Send a receipt for processessing to POST `/receipts/process`. An ID is returned.
  ```
  curl -X POST -H "Content-Type: application/json" -d @example http://localhost:8080/receipts/process # Use the example payload below
  ```

2. Look up the receipt by ID with GET request to `/receipts/{id}/points`.

```
curl -X GET http://localhost:8080/receipts/{id}/points # Replace the id with the id returned from the previous step
```

3. View all receipts entered with GET request to `/`.

```
curl -X GET http://localhost:8080/
```

Example Payload:
```
{
  "retailer": "Target",
  "purchaseDate": "2022-01-01",
  "purchaseTime": "13:01",
  "items": [
    {
      "shortDescription": "Mountain Dew 12PK",
      "price": "6.49"
    },{
      "shortDescription": "Emils Cheese Pizza",
      "price": "12.25"
    },{
      "shortDescription": "Knorr Creamy Chicken",
      "price": "1.26"
    },{
      "shortDescription": "Doritos Nacho Cheese",
      "price": "3.35"
    },{
      "shortDescription": "   Klarbrunn 12-PK 12 FL OZ  ",
      "price": "12.00"
    }
  ],
  "total": "35.35"
}

```


## Point Rules

The following rules define how many points are awarded to a receipt:

* One point for every alphanumeric character in the retailer name
* 50 points if the total is a round dollar amount with no cents
* 25 points if the total is a multiple of 0.25
* 5 points for every two items on the receipt
* If the trimmed length of the item description is a multiple of 3, multiply the price by 0.2 and round up to the nearest integer. The result is the number of points earned
* 6 points if the day in the purchase date is odd
* 10 points if the time of purchase is after 2:00pm and before 4:00pm
