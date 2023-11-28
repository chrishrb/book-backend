package utils

import (
	"encoding/json"

	"github.com/aws/aws-lambda-go/events"

	"github.com/chrishrb/bachelor-thesis/implementation/infrastructure/application/serverless/pkg/config"
	"github.com/chrishrb/bachelor-thesis/implementation/infrastructure/application/serverless/pkg/response"
)

func GetResponseError(statusCode int, err error) events.APIGatewayProxyResponse {
	body, _ := json.Marshal(response.Error{StatusCode: statusCode, Details: err.Error()})

	return events.APIGatewayProxyResponse{
		Headers: map[string]string{
			"Content-Type":                 "application/json",
			"Access-Control-Allow-Headers": "Content-Type,X-Amz-Date,Authorization,X-Api-Key,x-requested-with",
			"Access-Control-Allow-Methods": "POST,GET,OPTIONS,DELETE,PATCH",
			"Access-Control-Allow-Origin":  config.CorsAllowOrigin,
		},
		StatusCode: statusCode,
		Body:       string(body),
	}
}
