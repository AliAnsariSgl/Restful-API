package main

import (
	"encoding/json"
	"errors"
	"net/http"
	"os"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
)

// Device data model
type Device struct {
	ID          string `json:"id,omitempty"`
	DeviceModel string `json:"deviceModel,omitempty"`
	Name        string `json:"name,omitempty"`
	Note        string `json:"note,omitempty"`
	Serial      string `json:"serial,omitempty"`
}

//return bad request payload error
func IncompleteRequest(err error) events.APIGatewayProxyResponse {
	return events.APIGatewayProxyResponse{
		StatusCode: http.StatusBadRequest,
		Body:       "Bad Request\n" + err.Error(),
	}
}

// return internal server error
func InternalServerError() events.APIGatewayProxyResponse {
	return events.APIGatewayProxyResponse{
		StatusCode: http.StatusInternalServerError,
		Body:       "Internal Server Error",
	}
}

// Check Missing Payloads
func CheckPayloads(PostRequestJSON map[string]interface{}) error {
	Flag := false
	ErrorText := " Incomplete Request:needed payloads:"

	if PostRequestJSON["id"] == nil {
		Flag = true
		ErrorText += " 'id' "
	}
	if PostRequestJSON["deviceModel"] == nil {
		Flag = true
		ErrorText += " 'deviceModel' "
	}
	if PostRequestJSON["name"] == nil {
		Flag = true
		ErrorText += " 'name' "
	}
	if PostRequestJSON["note"] == nil {
		Flag = true
		ErrorText += " 'note' "
	}
	if PostRequestJSON["serial"] == nil {
		Flag = true
		ErrorText += " 'serial' "
	}

	//  return error
	if Flag == true {
		return errors.New(ErrorText)
	}

	return nil
}

// insert request json payload into new device
func makeNewDevice(PayloadJSON map[string]interface{}) Device {
	return Device{
		ID:          PayloadJSON["id"].(string),
		DeviceModel: PayloadJSON["deviceModel"].(string),
		Name:        PayloadJSON["name"].(string),
		Note:        PayloadJSON["note"].(string),
		Serial:      PayloadJSON["serial"].(string),
	}
}

//Insert Devices to database
func InsertDeviceToDatabase(device Device) error {
	// bad option !
	// sess, err := session.NewSession(&aws.Config{
	// 	Region:      aws.String("us-east-2"),
	// 	Credentials: credentials.NewStaticCredentials("AKIAJF6WYW3AP4UQU2XA", "uVc+iL2714LMM16C7cqCGh7XjOFRljBDkdtSKjyO", ""),
	// })

	// Initialize a session in AWS_REGION Using SDK
	region := os.Getenv("AWS_REGION")
	sess := session.Must(session.NewSession(&aws.Config{Region: &region}))

	// Create DynamoDB client
	svc := dynamodb.New(sess)
	db := mockDynamoDB{svc}
	// Convert Device object to dynamodb AttributeValues
	av, err := dynamodbattribute.MarshalMap(device)
	if err != nil {
		return err
	}

	// Insert av into DevicesTable
	DevicesTable := aws.String(os.Getenv("DEVICES_TABLE"))
	_, err = db.PutItem(&dynamodb.PutItemInput{
		Item:      av,
		TableName: DevicesTable,
	})

	if err != nil {
		return err
	}

	return nil
}

// Handler function
func Handler(r events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	// Unmarshal request payload
	var PostRequestJSON map[string]interface{}
	if err := json.Unmarshal([]byte(r.Body), &PostRequestJSON); err != nil {
		return IncompleteRequest(errors.New("")), nil
	}

	// finding missing payloads
	err := CheckPayloads(PostRequestJSON)
	if err != nil {
		return IncompleteRequest(err), nil
	}

	// call method to insert request json payload into new device
	device := makeNewDevice(PostRequestJSON)

	// Insert device object to dynamodb
	if err := InsertDeviceToDatabase(device); err != nil {
		return InternalServerError(), nil
	}

	// make response
	deviceJSON, err := json.Marshal(device)
	if err != nil {
		return InternalServerError(), nil
	}
	//create new device
	return events.APIGatewayProxyResponse{
		StatusCode: http.StatusCreated,
		Body:       string(deviceJSON),
	}, nil
}

func main() {
	lambda.Start(Handler)
}
