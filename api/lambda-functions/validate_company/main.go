package main

import (
	"context"
	"encoding/json"
	"fmt"
	"os"

	"github.com/alexdor/regtic/api/db"

	"github.com/aws/aws-lambda-go/events"

	"github.com/aws/aws-lambda-go/lambda"
)

type Response events.APIGatewayProxyResponse

const unprocessableError = "{\"error\": \"Please provide a company id \"}"

var origin = os.Getenv("ORIGIN")

func Handler(ctx context.Context, request events.APIGatewayProxyRequest) (Response, error) {
	var headers = map[string]string{
		"Content-Type":        "application/json",
		"X-Regtic-Func-Reply": "validate_company_handler",
	}

	if request.Headers["Origin"] == origin {
		headers["Access-Control-Allow-Origin"] = origin
		headers["Access-Control-Allow-Credentials"] = "true"
	}
	companyID := request.QueryStringParameters["id"]
	if len(companyID) == 0 {
		return Response{StatusCode: 400, Body: unprocessableError, Headers: headers}, nil
	}

	response, err := db.ValidateCompany(ctx, companyID)
	if err != nil {
		// TODO: Replace with sentry
		fmt.Println(err)
		return Response{StatusCode: 500}, err
	}
	body, err := json.Marshal(response)
	if err != nil {
		return Response{StatusCode: 500}, err
	}

	resp := Response{
		StatusCode: 200,
		Body:       string(body),
		Headers:    headers,
	}

	return resp, nil
}

func main() {
	lambda.Start(Handler)
}
