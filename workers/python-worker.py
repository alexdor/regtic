import http.client
from datetime import datetime


def hello(event, context):
    return {
        "statusCode": 200,
        "body": "<html><body><p>Hello! It is now FUCK YOU TIME.</p></body></html>",
        "headers": {"Content-Type": "text/html"},
    }
