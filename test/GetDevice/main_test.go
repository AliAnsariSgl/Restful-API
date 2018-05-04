package main

import (
	"testing"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbiface"
)

//  DynamoDB Mock
type mockDynamoDB struct {
	dynamodbiface.DynamoDBAPI
}

// Test type to request Handler and get response
type Test struct {
	request  events.APIGatewayProxyRequest
	response int
}

// List of test scenario
var tests = []Test{
	{
		// Test: id is Missing
		events.APIGatewayProxyRequest{
			PathParameters: map[string]string{
				"id": "",
			},
		},
		404,
	},
	{
		// Test: device does not exist
		events.APIGatewayProxyRequest{
			PathParameters: map[string]string{
				"id": "Not Found",
			},
		},
		404,
	},
	{
		// Test: complete request
		events.APIGatewayProxyRequest{
			PathParameters: map[string]string{
				"id": "id1",
			},
		},
		200,
	},
}

// GetItem method for mock DynamoDB
func (d *mockDynamoDB) GetItem(input *dynamodb.GetItemInput) (*dynamodb.GetItemOutput, error) {

	out := dynamodb.GetItemOutput{}
	id := input.Key["id"].S
	if *id == "id1" {
		out.SetItem(
			map[string]*dynamodb.AttributeValue{
				"id": &dynamodb.AttributeValue{S: id},
			},
		)
	}
	// Otherwise, just return the empty outout
	return &out, nil
}

// Actual test function
func TestHandler(t *testing.T) {
	// test and check the response for every test
	for i, test := range tests {
		response, _ := Handler(test.request)
		//  does not match --> error
		if response.StatusCode != test.response {
			t.Errorf("#%d: Expected: %d, Actual: %d, %s", i, test.response, response.StatusCode, response.Body)
		}
	}
}
