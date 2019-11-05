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

const unprocessableError = "{\"error\": \"Please provide a company name \"}"

var origin = os.Getenv("ORIGIN")

func Handler(ctx context.Context, request events.APIGatewayProxyRequest) (Response, error) {
	var headers = map[string]string{
		"Content-Type":        "application/json",
		"X-Regtic-Func-Reply": "find_companies_handler",
	}

	if request.Headers["Origin"] == origin {
		headers["Access-Control-Allow-Origin"] = origin
		headers["Access-Control-Allow-Credentials"] = "true"
	}
	companyName := request.QueryStringParameters["name"]
	if len(companyName) == 0 {
		return Response{StatusCode: 400, Body: unprocessableError}, nil
	}

	companies, err := db.FindCompany(ctx, companyName)
	if err != nil {
		// TODO: Replace with sentry
		fmt.Println(err)
		return Response{StatusCode: 500}, err
	}

	body, err := json.Marshal(map[string]interface{}{
		"companies": companies,
	})
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
