#!/bin/bash

echo "Simple API client:"
echo "insert endpoint:"
read endpoint
echo "insert HTTP method:"
read method
echo "insert request body"
read request_body
echo "insert request body type "
read request_body_type

curl_command="curl -X $method"

if [ -n "$request_body" ]; then
  curl_command="$curl_command -d '$request_body'"
fi

if [ -n "$request_body_type" ]; then
  curl_command="$curl_command -H 'Content-Type: $request_body_type'"
fi

curl_command="$curl_command $endpoint"

eval $curl_command
