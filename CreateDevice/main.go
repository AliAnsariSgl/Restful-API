package main

import (
	"DevicesServ/DataModel"
	responses "DevicesServ/Responses"
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

var sess *session.Session

// Init function Initialize session
func init() {
	// Initialize a session
	region := os.Getenv("AWS_REGION")
	sess = session.Must(session.NewSession(&aws.Config{Region: &region}))

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
func makeNewDevice(PayloadJSON map[string]interface{}) model.Device {
	return model.Device{
		ID:          PayloadJSON["id"].(string),
		DeviceModel: PayloadJSON["deviceModel"].(string),
		Name:        PayloadJSON["name"].(string),
		Note:        PayloadJSON["note"].(string),
		Serial:      PayloadJSON["serial"].(string),
	}
}

//Insert Devices to database
func InsertDeviceToDatabase(device model.Device) error {

	// Create DynamoDB client
	svc := dynamodb.New(sess)
	//db := mockDynamoDB{svc}
	// Convert Device object to dynamodb AttributeValues
	av, err := dynamodbattribute.MarshalMap(device)
	if err != nil {
		return err
	}

	// Insert av into DevicesTable
	DevicesTable := aws.String(os.Getenv("DEVICES_TABLE"))
	_, err = svc.PutItem(&dynamodb.PutItemInput{
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
		return responses.IncompleteRequest(errors.New("")), nil
	}

	// finding missing payloads
	err := CheckPayloads(PostRequestJSON)
	if err != nil {
		return responses.IncompleteRequest(err), nil
	}

	// call method to insert request json payload into new device
	device := makeNewDevice(PostRequestJSON)

	// Insert device object to dynamodb
	if err := InsertDeviceToDatabase(device); err != nil {
		return responses.InternalServerError(), nil
	}

	// make response
	deviceJSON, err := json.Marshal(device)
	if err != nil {
		return responses.InternalServerError(), nil
	}
	//create new device
	return events.APIGatewayProxyResponse{
		StatusCode: http.StatusCreated,
		Headers:    map[string]string{"Content-Type": "application/json"},
		Body:       string(deviceJSON),
	}, nil
}

func main() {
	lambda.Start(Handler)
}
