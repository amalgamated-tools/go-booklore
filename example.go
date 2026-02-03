package main

import (
	"bytes"
	"fmt"
	"os"
	"strings"
	"text/template"

	"github.com/buger/jsonparser"
)

type TestEndpoint struct {
	Path        string
	Method      string
	OperationID string
	Summary     string
}

func main() {
	specBytes, err := os.ReadFile("api-docs.json")
	if err != nil {
		panic(err)
	}

	var endpoints []TestEndpoint

	err = jsonparser.ObjectEach(specBytes, func(pathKey []byte, pathValue []byte, _ jsonparser.ValueType, _ int) error {
		path := string(pathKey)

		err = jsonparser.ObjectEach(pathValue, func(methodKey []byte, methodValue []byte, _ jsonparser.ValueType, _ int) error {
			method := strings.ToUpper(string(methodKey))

			// Skip non-HTTP methods
			if !isHTTPMethod(method) {
				return nil
			}

			operationID, _ := jsonparser.GetString(methodValue, "operationId")
			summary, _ := jsonparser.GetString(methodValue, "summary")

			endpoints = append(endpoints, TestEndpoint{
				Path:        path,
				Method:      method,
				OperationID: operationID,
				Summary:     summary,
			})

			return nil
		})
		if err != nil {
			return err
		}
		return nil
	}, "paths")

	if err != nil {
		panic(err)
	}

	// Generate test file
	testCode := generateTestFile(endpoints)

	if err := os.WriteFile("pkg/booklore/client_test.go", []byte(testCode), 0644); err != nil {
		panic(err)
	}

	fmt.Printf("Generated %d test functions in pkg/booklore/client_test.go\n", len(endpoints))
}

func isHTTPMethod(method string) bool {
	switch method {
	case "GET", "POST", "PUT", "PATCH", "DELETE", "HEAD", "OPTIONS":
		return true
	default:
		return false
	}
}

func generateTestFile(endpoints []TestEndpoint) string {
	tmpl := `package booklore

import (
	"context"
	"fmt"
	"os"
	"testing"
)

// Generated test functions for API endpoints
// These tests hit the real API endpoint configured in environment variables

{{range .Endpoints}}
// Test{{.MethodName}} tests the {{.Method}} {{.Path}} endpoint
func Test{{.MethodName}}(t *testing.T) {
	ctx := context.Background()
	client, err := setupTestClient()
	if err != nil {
		t.Skipf("skipping test: %v", err)
	}

	// TODO: Set up parameters for {{.OperationID}}
	// TODO: Add assertions to validate response

	_ = client
	_ = ctx
}

{{end}}

// setupTestClient creates a test client using environment variables
func setupTestClient() (*ClientWithResponses, error) {
	server := os.Getenv("BOOKLORE_SERVER")
	if server == "" {
		return nil, fmt.Errorf("BOOKLORE_SERVER environment variable not set")
	}

	client, err := NewClientWithResponses(server)
	if err != nil {
		return nil, err
	}

	return client, nil
}
`

	data := struct {
		Endpoints []struct {
			MethodName  string
			Method      string
			Path        string
			OperationID string
		}
	}{
		Endpoints: make([]struct {
			MethodName  string
			Method      string
			Path        string
			OperationID string
		}, len(endpoints)),
	}

	for i, ep := range endpoints {
		data.Endpoints[i] = struct {
			MethodName  string
			Method      string
			Path        string
			OperationID string
		}{
			MethodName:  generateTestFunctionName(ep),
			Method:      ep.Method,
			Path:        ep.Path,
			OperationID: ep.OperationID,
		}
	}

	t, err := template.New("test").Parse(tmpl)
	if err != nil {
		panic(err)
	}

	var buf bytes.Buffer
	if err := t.Execute(&buf, data); err != nil {
		panic(err)
	}

	return buf.String()
}

func generateTestFunctionName(ep TestEndpoint) string {
	// Use OperationID if available, otherwise derive from path and method
	if ep.OperationID != "" {
		return toPascalCase(ep.OperationID) + "_" + ep.Method
	}

	// Derive from path: /api/books/{id}/details -> BooksDetails_GET
	parts := strings.Split(strings.Trim(ep.Path, "/"), "/")
	var name string
	for _, part := range parts {
		if strings.HasPrefix(part, "{") && strings.HasSuffix(part, "}") {
			continue // skip path parameters
		}
		name += toPascalCase(part)
	}

	return name + "_" + ep.Method
}

func toPascalCase(s string) string {
	parts := strings.FieldsFunc(s, func(r rune) bool {
		return r == '_' || r == '-'
	})

	var result string
	for _, part := range parts {
		if len(part) > 0 {
			result += strings.ToUpper(part[:1]) + strings.ToLower(part[1:])
		}
	}

	return result
}

// func main() {
// 	var server, username, password string
// 	server = os.Getenv("BOOKLORE_SERVER")
// 	username = os.Getenv("BOOKLORE_USERNAME")
// 	password = os.Getenv("BOOKLORE_PASSWORD")
// 	if server == "" || username == "" || password == "" {
// 		panic("Please set BOOKLORE_SERVER, BOOKLORE_USERNAME, and BOOKLORE_PASSWORD environment variables")
// 	}

// 	client, err := booklore.NewClient(
// 		server,
// 	)
// 	if err != nil {
// 		panic(err)
// 	}
// 	err = client.LoginWithCredentials(context.Background(), booklore.Credentials{
// 		Username: username,
// 		Password: password,
// 	})
// 	if err != nil {
// 		panic(err)
// 	}
// 	clientWithResponses := client.ClientWithResponses()
// 	r, err := clientWithResponses.ListBooksWithResponse(
// 		context.Background(),
// 		nil,
// 	)
// 	if err != nil {
// 		panic(err)
// 	}
// 	if r.JSON200 == nil {
// 		panic("received nil JSON200 response from ListBooksWithResponse")
// 	}
// 	for _, book := range *r.JSON200 {
// 		println(book.Id)
// 	}
// }
