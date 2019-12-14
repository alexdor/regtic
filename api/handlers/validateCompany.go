package handlers

import (
	"context"
	"database/sql"
	"encoding/json"
	"errors"
	"fmt"

	"github.com/alexdor/regtic/api/db"
	"github.com/alexdor/regtic/api/interfaces"

	"github.com/aws/aws-lambda-go/events"
)

const unprocessableValidateError = "{\"error\": \"Please provide a company id \"}"
const companyIDNotFound = "{\"error\": \"Company id provided doesn't exist in the database \"}"

func ValidateCompanyHandler(ctx context.Context, request events.APIGatewayProxyRequest, headers map[string]string) (interfaces.Response, error) {
	headers["X-Regtic-Func-Reply"] = "validate_company_handler"

	companyID := request.QueryStringParameters["id"]
	if len(companyID) == 0 {
		return interfaces.Response{StatusCode: 400, Body: unprocessableValidateError, Headers: headers}, nil
	}

	response, err := db.ValidateCompany(ctx, companyID)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return interfaces.Response{StatusCode: 404, Body: companyIDNotFound}, nil
		}
		// TODO: Replace with sentry
		fmt.Println(err)
		return interfaces.Response{StatusCode: 500}, err
	}
	body, err := json.Marshal(response)
	if err != nil {
		return interfaces.Response{StatusCode: 500}, err
	}

	resp := interfaces.Response{
		StatusCode: 200,
		Body:       string(body),
		Headers:    headers,
	}

	return resp, nil
}
