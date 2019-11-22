package handlers

import (
	"context"
	"net/http"
	"encoding/json"
	"errors"
	"fmt"

  "github.com/alexdor/regtic/api/interfaces"

	"github.com/aws/aws-lambda-go/events"
)

const noEmailAdded = "{\"error\": \"No email supplied.\"}"
const webhookUrl = "https://hooks.slack.com/services/T00000000/B00000000/XXXXXXXXXXXXXXXXXXXXXXXX"

func SendMailToSlack(ctx context.Context, request events.APIGatewayProxyRequest, headers map[string]string) (interfaces.Response, error) {
	headers["X-Regtic-Func-Reply"] = "send_mail_to_slack"

	email := request.QueryStringParameters["email"]
	if len(email) == 0 {
		return interfaces.Response{StatusCode: 400, Body: noEmailAdded}, nil
	}

	body, err := json.Marshal(map[string]interface{}{
		"text": email,
	})
	if err != nil {
		return interfaces.Response{StatusCode: 500}, err
	}

  http.Post(webhookUrl, "application/json", body)

	resp := interfaces.Response{
		StatusCode: 200,
		Body:       "",
		Headers:    headers,
	}

	return resp, nil
}
