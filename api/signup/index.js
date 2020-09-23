/* eslint-disable @typescript-eslint/no-unsafe-assignment */
/* eslint-disable @typescript-eslint/no-unsafe-member-access */
const { IncomingWebhook } = require("@slack/webhook");
const validator = require("email-validator");
const webhook = new IncomingWebhook(process.env.SLACK_WEBHOOK_URL);
const headers = {
  "Content-Type": "application/json",
  "Access-Control-Allow-Credentials": true,
  "X-Regtic-Func-Reply": "signup_handler",
  "Access-Control-Allow-Origin": process.env.ORIGIN,
};
module.exports.signup = async (event) => {
  const isEmailValid =
    event.queryStringParameters &&
    validator.validate(event.queryStringParameters.email);
  if (!isEmailValid) {
    return {
      statusCode: 403,
      body: JSON.stringify({
        error: "Please provide a valid email",
      }),
      headers,
    };
  }

  try {
    const email = event.queryStringParameters.email;
    await webhook.send({
      text: email,
    });
  } catch (error) {
    return {
      statusCode: 500,
      headers,
      body: JSON.stringify({
        error: "Internal server error",
      }),
    };
  }

  return {
    statusCode: 200,
    headers,
    body: JSON.stringify({
      message: "OK",
    }),
  };
};
