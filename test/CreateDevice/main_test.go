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

// PutItem function for DynamoDB Mock
func (d *mockDynamoDB) PutItem(input *dynamodb.PutItemInput) (*dynamodb.PutItemOutput, error) {
	// mockDynamoDB PutItem function response
	return new(dynamodb.PutItemOutput), nil
}

// List of tests scenario
var tests = []Test{
	{
		// Bad json request test
		events.APIGatewayProxyRequest{
			Body: "{a2s1d}",
		},
		400,
	},
	{
		// lost payloads test (id, deviceModel)
		events.APIGatewayProxyRequest{
			Body: "{\"id\":\"1\", \"deviceModel\":\"deviceModel1\"}",
		},
		400,
	},
	{
		// lost payloads test
		events.APIGatewayProxyRequest{
			Body: "{\"id\":\"1\",\"note\":\"note1\", \"deviceModel\":\"deviceModel1\"}",
		},
		400,
	},
	{
		// lost payloads test
		events.APIGatewayProxyRequest{
			Body: "{\"name\":\"name1\", \"note\":\"note1\", \"serial\":\"serial1\"}",
		},
		400,
	},
	{
		//Complete request .Response 201
		events.APIGatewayProxyRequest{
			// Headers: map[string]string{"Content-Type": "application/json"},
			Body: "{\"id\":\"1\", \"deviceModel\":\"deviceModel1\", \"name\":\"name1\", \"note\":\"note1\", \"serial\":\"serial1\"}",
		},
		201,
	},
}

// TestHandler function
func TestHandler(t *testing.T) {
	for i, Item := range tests {
		// Send request of each test to Handler function
		response, _ := Handler(Item.request)
		if response.StatusCode != Item.response {

			t.Errorf("Test Fail in tests[%d]: Expected: %d, Actual: %d", i, Item.response, response.StatusCode)

		}

	}

}
