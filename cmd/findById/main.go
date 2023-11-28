package main

import (
	"encoding/json"
	"net/http"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	_ "github.com/lib/pq"

	"github.com/chrishrb/bachelor-thesis/implementation/infrastructure/application/serverless/pkg/books"
	"github.com/chrishrb/bachelor-thesis/implementation/infrastructure/application/serverless/pkg/config"
	"github.com/chrishrb/bachelor-thesis/implementation/infrastructure/application/serverless/pkg/utils"
)

var store *books.PostgresStore

func init() {
	store = books.NewPostgresStore(config.DBConnectionString)
}

func handler(req events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	id := req.PathParameters["id"]

	// get book from db
	book, err := store.ListByID(id)
	if err != nil {
		return utils.GetResponseError(http.StatusNotFound, err), nil
	}

	// to json
	booksJSON, err := json.Marshal(book)
	if err != nil {
		return utils.GetResponseError(http.StatusInternalServerError, err), nil
	}

	return events.APIGatewayProxyResponse{
		Body:       string(booksJSON),
		StatusCode: http.StatusOK,
		Headers: map[string]string{
			"Content-Type":                 "application/json",
			"Access-Control-Allow-Headers": "Content-Type,X-Amz-Date,Authorization,X-Api-Key,x-requested-with",
			"Access-Control-Allow-Methods": "OPTIONS,GET",
			"Access-Control-Allow-Origin":  config.CorsAllowOrigin,
		},
	}, nil
}

func main() {
	lambda.Start(handler)
}
