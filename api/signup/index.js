const { IncomingWebhook } = require("@slack/webhook");
const validator = require("email-validator");
const webhook = new IncomingWebhook(process.env.SLACK_WEBHOOK_URL);

module.exports.signup = async event => {
  const isEmailValid = validator.validate(event.queryStringParameters.email);
  if (!isEmailValid) {
    return {
      statusCode: 403,
      body: JSON.stringify({
        error: "Please provide a valid email"
      })
    };
  }

  try {
    const email = event.queryStringParameters.email;
    await webhook.send({
      text: email
    });
  } catch (error) {
    return {
      statusCode: 500,
      body: JSON.stringify({
        error: "Internal server error"
      })
    };
  }

  return {
    statusCode: 200,
    body: JSON.stringify({
      message: "OK"
    })
  };
};
