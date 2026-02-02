# go-booklore

A type-safe Go client SDK for the Booklore API, auto-generated from the OpenAPI 3.1.0 specification.

## Overview

`go-booklore` provides a modern, strongly-typed Go interface to the Booklore book management API. The SDK is automatically generated from the OpenAPI specification, ensuring it stays in sync with the API and provides full type safety for all requests and responses.

## Features

- **Type-safe**: Full compile-time type checking for all API requests and responses
- **Structured responses**: Response objects provide clean status code handling with typed fields for different HTTP status codes
- **Context support**: Built-in support for request cancellation and timeouts
- **Request customization**: Easy way to add custom headers, authentication, and middleware
- **Auto-generated**: Code generated directly from OpenAPI spec, ensuring API consistency
- **Authentication ready**: Built-in support for bearer token and custom authentication schemes
- **Error handling**: Comprehensive error handling with proper HTTP status code visibility

## Installation

```bash
go get github.com/amalgamated-tools/go-booklore
```

## Quick Start

### Basic Usage

```go
package main

import (
	"context"
	"fmt"
	"net/http"

	"github.com/amalgamated-tools/go-booklore/pkg/client"
)

func main() {
	// Create a client
	c, err := client.NewClientWithResponses("https://api.booklore.example.com")
	if err != nil {
		panic(err)
	}

	// Make a request
	resp, err := c.GetBooksWithResponse(context.Background(), nil)
	if err != nil {
		panic(err)
	}

	// Handle response
	if resp.StatusCode() == http.StatusOK {
		books := resp.JSON200
		fmt.Printf("Found %d books\n", len(*books))
	}
}
```

### Authentication

```go
import (
	"context"
	"net/http"
	"github.com/amalgamated-tools/go-booklore/pkg/client"
)

func main() {
	c, err := client.NewClientWithResponses(
		"https://api.booklore.example.com",
		client.WithRequestEditorFn(func(ctx context.Context, req *http.Request) error {
			req.Header.Set("Authorization", "Bearer "+token)
			return nil
		}),
	)
	if err != nil {
		panic(err)
	}

	// Use authenticated client
	resp, err := c.GetBooksWithResponse(context.Background(), nil)
	// ...
}
```

### Login Example

```go
loginResp, err := c.LoginUserWithResponse(ctx, client.LoginUserJSONRequestBody{
	Username: strToPtr("username"),
	Password: strToPtr("password"),
})

if loginResp.StatusCode() == http.StatusOK && loginResp.JSON200 != nil {
	token := loginResp.JSON200.AccessToken
	// Use token for subsequent authenticated requests
}
```

### Custom HTTP Client

```go
import (
	"net/http"
	"time"
)

httpClient := &http.Client{
	Timeout: 30 * time.Second,
}

c, err := client.NewClientWithResponses(
	"https://api.booklore.example.com",
	client.WithHTTPClient(httpClient),
)
```

## API Reference

### Client Initialization

```go
// Basic client
c, err := client.NewClientWithResponses(baseURL)

// With HTTP client
c, err := client.NewClientWithResponses(baseURL, client.WithHTTPClient(httpClient))

// With authentication
c, err := client.NewClientWithResponses(
	baseURL,
	client.WithRequestEditorFn(authFn),
)

// Combined options
c, err := client.NewClientWithResponses(
	baseURL,
	client.WithHTTPClient(httpClient),
	client.WithRequestEditorFn(authFn),
)
```

### Response Handling

All API methods ending with `WithResponse` return a response object with a `StatusCode()` method and status-specific fields:

```go
resp, err := c.GetBookWithResponse(ctx, bookID)
if err != nil {
	return err
}

switch resp.StatusCode() {
case http.StatusOK:
	book := resp.JSON200
	// Process book
	
case http.StatusNotFound:
	return fmt.Errorf("book not found")
	
case http.StatusForbidden:
	return fmt.Errorf("access denied")
	
default:
	return fmt.Errorf("unexpected status: %d", resp.StatusCode())
}
```

### Available Response Fields

Each response type includes typed fields for different HTTP status codes:

- `JSON200` - for 200 OK responses
- `JSON201` - for 201 Created responses
- `JSON204` - for 204 No Content responses
- `JSON400` - for 400 Bad Request responses
- `JSON403` - for 403 Forbidden responses
- `JSON404` - for 404 Not Found responses
- `JSON500` - for 500 Internal Server Error responses

## Data Models

The SDK includes strongly-typed models for all API entities:

- `Book` - Book information
- `Library` - Library information
- `Shelf` - Book shelf/collection
- `User` - User profile
- `BookLoreUser` - Extended user information
- `BookMetadata` - Book metadata
- And many more...

See `pkg/client/client.gen.go` for the complete list of available types.

## Development

### Regenerating the SDK

The SDK is auto-generated from the OpenAPI specification. To regenerate after API changes:

```bash
# Install dependencies
go mod tidy

# Generate SDK from OpenAPI spec
go generate ./...
```

The generation process uses `oapi-codegen` to create type-safe client code from the OpenAPI specification.

### Project Structure

```
.
├── main.go                 # Example usage
├── go.mod                  # Go module definition
├── pkg/
│   └── client/
│       ├── client.gen.go   # Auto-generated client code
│       └── models.go       # Auto-generated data models
├── generate.go             # Code generation directives
└── cfg.yaml               # Configuration
```

## Requirements

- Go 1.25.5 or later

## Dependencies

- `github.com/getkin/kin-openapi` - OpenAPI parsing and validation
- `github.com/oapi-codegen/runtime` - Runtime support for generated code

## Error Handling

The SDK handles errors at two levels:

1. **Network/Client errors**: Returned directly as error values
2. **API errors**: Returned in response status codes and JSON bodies

```go
resp, err := c.GetBookWithResponse(ctx, bookID)

// Handle network errors
if err != nil {
	return err
}

// Handle API errors via status code
if resp.StatusCode() != http.StatusOK {
	if resp.JSON400 != nil {
		return fmt.Errorf("validation error: %v", resp.JSON400)
	}
	if resp.JSON404 != nil {
		return fmt.Errorf("not found: %v", resp.JSON404)
	}
}

// Success
book := resp.JSON200
```

## Context and Timeouts

All API methods accept a `context.Context` for managing request timeouts and cancellation:

```go
ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
defer cancel()

resp, err := c.GetBooksWithResponse(ctx, nil)
```

## Examples

### Get All Books

```go
resp, err := c.GetBooksWithResponse(ctx, nil)
if err != nil {
	return err
}

if resp.StatusCode() == http.StatusOK {
	books := resp.JSON200
	for _, book := range *books {
		fmt.Printf("%s by %s\n", book.Title, book.Author)
	}
}
```

### Create a Book

```go
createResp, err := c.CreateBookWithResponse(ctx, client.CreateBookJSONRequestBody{
	Title:  strToPtr("The Go Programming Language"),
	Author: strToPtr("Donovan & Kernighan"),
})

if createResp.StatusCode() == http.StatusCreated && createResp.JSON201 != nil {
	newBook := createResp.JSON201
	fmt.Printf("Created book with ID: %d\n", newBook.ID)
}
```

### Update a Book

```go
updateResp, err := c.UpdateBookWithResponse(ctx, bookID, client.UpdateBookJSONRequestBody{
	Title: strToPtr("Updated Title"),
})

if updateResp.StatusCode() == http.StatusOK {
	fmt.Println("Book updated successfully")
}
```

### Delete a Book

```go
deleteResp, err := c.DeleteBookWithResponse(ctx, bookID)

if deleteResp.StatusCode() == http.StatusNoContent {
	fmt.Println("Book deleted successfully")
}
```

## Contributing

Contributions are welcome! Since this SDK is auto-generated, changes should be made to the OpenAPI specification rather than the generated code.

## License

[Add your license information here]

## Support

For issues or questions:
- Check the [SDK Usage Guide](./SDK_USAGE.md)
- Review the OpenAPI specification in `api-docs.json`
- Open an issue in the repository

## Changelog

### v0.1.0
- Initial release
- Auto-generated SDK from OpenAPI 3.1.0 specification
- Support for all Booklore API endpoints
- Type-safe request and response handling
