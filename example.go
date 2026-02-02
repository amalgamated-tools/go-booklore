package main

import (
	"context"
	"os"

	"github.com/amalgamated-tools/go-booklore/pkg/booklore"
)

func main() {
	var server, username, password string
	server = os.Getenv("BOOKLORE_SERVER")
	username = os.Getenv("BOOKLORE_USERNAME")
	password = os.Getenv("BOOKLORE_PASSWORD")
	if server == "" || username == "" || password == "" {
		panic("Please set BOOKLORE_SERVER, BOOKLORE_USERNAME, and BOOKLORE_PASSWORD environment variables")
	}

	client, err := booklore.NewClient(
		server,
	)
	if err != nil {
		panic(err)
	}
	err = client.LoginWithCredentials(context.Background(), booklore.Credentials{
		Username: username,
		Password: password,
	})
	if err != nil {
		panic(err)
	}
	clientWithResponses := client.ClientWithResponses()
	r, err := clientWithResponses.ListBooksWithResponse(
		context.Background(),
		nil,
	)
	if err != nil {
		panic(err)
	}
	if r.JSON200 == nil {
		panic("received nil JSON200 response from ListBooksWithResponse")
	}
	for _, book := range *r.JSON200 {
		println(book.Id)
	}
}
