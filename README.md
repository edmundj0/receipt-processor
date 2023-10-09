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
  - Response: A JSON object containing the number of points awarded


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

## Point Rules

The following rules define how many points are awarded to a receipt:

* One point for every alphanumeric character in the retailer name
* 50 points if the total is a round dollar amount with no cents
* 25 points if the total is a multiple of 0.25
* 5 points for every two items on the receipt
* If the trimmed length of the item description is a multiple of 3, multiply the price by 0.2 and round up to the nearest integer. The result is the number of points earned
* 6 points if the day in the purchase date is odd
* 10 points if the time of purchase is after 2:00pm and before 4:00pm

