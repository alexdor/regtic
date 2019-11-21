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

const unprocessableFindError = "{\"error\": \"Please provide a company name \"}"
const noCompanyFound = "{\"error\": \"Company name provided doesn't exist in the database \"}"

func FindCompanies(ctx context.Context, request events.APIGatewayProxyRequest, headers map[string]string) (interfaces.Response, error) {
	headers["X-Regtic-Func-Reply"] = "find_companies_handler"

	companyName := request.QueryStringParameters["name"]
	if len(companyName) == 0 {
		return interfaces.Response{StatusCode: 400, Body: unprocessableFindError}, nil
	}

	companies, err := db.FindCompany(ctx, companyName)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return interfaces.Response{StatusCode: 404, Body: noCompanyFound}, nil
		}
		// TODO: Replace with sentry
		fmt.Println(err)
		return interfaces.Response{StatusCode: 500}, err
	}

	body, err := json.Marshal(map[string]interface{}{
		"companies": companies,
	})
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
