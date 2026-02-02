#!/bin/bash
set -e
echo "Downloading Spec"
wget -q -O api-docs.json https://demo.booklore.org/api/v1/api-docs

# Use jq to fix the OpenAPI spec
# 1. Fix exclusiveMinimum: 0 -> true
# 2. Convert all "*/*" content types to "application/json" (required for oapi-codegen to generate JSON200 fields)
# 3. Remove /** suffix from paths
# 4. Add the missing {token} and {bookId} parameters to any path that includes them but doesn't define them
jq '
  (.. | select(.exclusiveMinimum? == 0)) .exclusiveMinimum = true |
  (.. | select(.content? != null) | .content) |= 
    with_entries(
      if .key == "*/*" then
        .key = "application/json"
      else
        .
      end
    ) |
  .paths |= with_entries(
    .key |= gsub("/\\*\\*"; "") |
    
    # Helper to add missing path parameters
    def add_param(name):
      if (.value.parameters? // [] | any(.name == name)) then
        .
      else
        .value.parameters += [{
          "name": name,
          "in": "path",
          "required": true,
          "schema": { "type": "integer", "format": "int64" }
        }]
      end;

    if .key | contains("{token}") then
      # token is usually string
      if (.value.parameters? // [] | any(.name == "token")) then
        .
      else
        .value.parameters += [{
          "name": "token",
          "in": "path",
          "required": true,
          "schema": { "type": "string" }
        }]
      end
    else . end |
    
    if .key | contains("{bookId}") then
      add_param("bookId")
    else . end
  )
' api-docs.json > api-docs.tmp.json && mv api-docs.tmp.json api-docs.json

echo "api-docs.json fixed."

echo "Generating client code"
go mod tidy
go generate ./...
