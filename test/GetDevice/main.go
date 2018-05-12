package main

import (
	model "DevicesServ/DataModel"
	responses "DevicesServ/Responses"
	"encoding/json"
	"net/http"
	"os"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
)

var sess *session.Session

// Init function Initialize session
func init() {
	// Initialize a session
	region := os.Getenv("AWS_REGION")
	sess = session.Must(session.NewSession(&aws.Config{Region: &region}))

}

// return a specific device if exists
func Handler(req events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	// Get request id, and check if id is valid
	ID := req.PathParameters["id"]
	if ID == "" {
		return events.APIGatewayProxyResponse{
			StatusCode: http.StatusNotFound,
			Body:       "Not Found",
		}, nil
	}

	// Create DynamoDB client
	svc := dynamodb.New(sess)
	db := mockDynamoDB{svc}
	// Get Device
	DevicesTable := aws.String(os.Getenv("DEVICES_TABLE"))
	result, err := db.GetItem(&dynamodb.GetItemInput{
		TableName: DevicesTable,
		Key: map[string]*dynamodb.AttributeValue{
			"id": {
				S: aws.String(ID),
			},
		},
	})

	if err != nil {
		return responses.InternalServerError(), nil
	}

	// If length of result is zero response 404
	if len(result.Item) == 0 {
		return events.APIGatewayProxyResponse{
			StatusCode: http.StatusNotFound,
			Body:       "Not Found",
		}, nil
	}

	// Create empty Device object for return to response
	device := model.Device{}

	// Convert dynamodb result to Device object
	if err = dynamodbattribute.UnmarshalMap(result.Item, &device); err != nil {
		return responses.InternalServerError(), nil
	}

	// Return found device to response
	deviceJSON, err := json.Marshal(device)
	if err != nil {
		return responses.InternalServerError(), nil
	}
	//return response for found device
	return events.APIGatewayProxyResponse{
		StatusCode: http.StatusOK,
		Body:       string(deviceJSON),
	}, nil
}

func main() {
	lambda.Start(Handler)
}
