//go:generate sqlboiler psql --no-hooks --struct-tag-casing camel
package main

import (
	"context"
	"os"
	"strings"

	"github.com/alexdor/regtic/api/handlers"
	"github.com/alexdor/regtic/api/interfaces"
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

var origin = os.Getenv("ORIGIN")

func Router(ctx context.Context, request events.APIGatewayProxyRequest) (interfaces.Response, error) {
	var headers = map[string]string{
		"Content-Type":                "application/json",
		"Access-Control-Allow-Origin": origin,
	}

	if request.Headers["Origin"] == origin {
		headers["Access-Control-Allow-Credentials"] = "true"
	}
	// TODO: Improve this
	switch {
	case strings.Contains(request.Path, "find_companies"):
		return handlers.FindCompanies(ctx, request.QueryStringParameters["name"], headers)
	case strings.Contains(request.Path, "validate_company"):
		return handlers.ValidateCompanyHandler(ctx, request.QueryStringParameters["id"], headers)
	}

	return interfaces.Response{
		StatusCode: 404,
		Headers:    headers,
	}, nil
}

func main() {
	lambda.Start(Router)
}
