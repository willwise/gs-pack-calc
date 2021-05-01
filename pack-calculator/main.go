package main

import (
	"encoding/json"
	"fmt"
	"math"
	"net/http"
	"sort"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

//create data structures
type RequestBody struct {
	Quantity int   `json:"quantity"`
	PacksArr []int `json:"packsarr"`
}

type Response struct {
	PacksNeeded []int
}

//function to handle the request check data and respond
func handler(request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {

	//check if json
	if request.Headers["content-type"] != "application/json" && request.Headers["Content-Type"] != "application/json" {
		return events.APIGatewayProxyResponse{
			StatusCode: http.StatusBadRequest,
			Body:       "request must be in json",
		}, nil
	}

	data := &RequestBody{}

	_ = json.Unmarshal([]byte(request.Body), data)

	//TODO
	//check input is correct format
	if data.Quantity <= 0 || data.PacksArr == nil || len(data.PacksArr) == 0 {
		return events.APIGatewayProxyResponse{
			StatusCode: http.StatusBadRequest,
			Body:       "input array is in incorrect format",
		}, nil
	}

	order := calcOrder(data.PacksArr, data.Quantity)

	response := Response{PacksNeeded: order}

	var jsonData []byte

	jsonData, _ = json.Marshal(response)

	return events.APIGatewayProxyResponse{
		StatusCode: http.StatusOK,
		Body:       string(jsonData),
		Headers: map[string]string{
			"Access-Control-Allow-Origin":  "*",
			"Access-Control-Allow-Headers": "*",
			"Access-Control-Allow-Methods": "POST, OPTIONS",
		},
	}, nil
}

//function to take the array input and order it from biggest to smallest
func sortArrayDesc(arr []int) []int {
	sort.Sort(sort.Reverse(sort.IntSlice(arr)))
	return arr
}

// Function to take the quantity input and loop through the array of packs allowed and return the array of packs required to fulfill the order
func calcOrder(arr []int, quantity int) []int {
	var returnArr []int
	// Order the array desc
	sortedArr := sortArrayDesc(arr)
	// Loop through array
	fmt.Println(len(sortedArr))

	// If the quantity is between the final 2 choose the larger number to prevent extra packs
	if quantity > sortedArr[len(sortedArr)-1] && quantity < sortedArr[len(sortedArr)-2] {
		returnArr = append(returnArr, sortedArr[len(sortedArr)-2])
		return returnArr
	}

	for i := 0; i < len(sortedArr); i++ {
		// Check if remaining items is 0
		if quantity == 0 {
			return returnArr
		}
		// Check if the pack size fits into the remaining items
		if quantity/sortedArr[i] >= 1 {
			// If it does work out the amount of times
			numOfPacks := quantity / sortedArr[i]
			for a := 0; a < numOfPacks; a++ {
				returnArr = append(returnArr, sortedArr[i])
			}
			// Work out how much (if any) the remainder is
			quantity = int(math.Mod(float64(quantity), float64(sortedArr[i])))
		}
		// If remainder is > 0 and it is the last one add another pack
		if i == len(sortedArr)-1 && quantity > 0 {
			returnArr = append(returnArr, sortedArr[i])
			return returnArr
		}

	}
	return returnArr
}

func main() {
	lambda.Start(handler)
}
