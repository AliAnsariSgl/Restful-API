package responses

import (
	"net/http"

	"github.com/aws/aws-lambda-go/events"
)

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
