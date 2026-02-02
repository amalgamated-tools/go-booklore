package main

import (
	"context"

	"github.com/amalgamated-tools/go-booklore/pkg/booklore"
)

func main() {
	client, err := booklore.NewClient(
		"https://demo.booklore.org",
	)
	if err != nil {
		panic(err)
	}
	err = client.LoginWithCredentials(context.Background(), booklore.Credentials{
		Username: "booklore",
		Password: "<https://github.com/booklore-app/booklore?tab=readme-ov-file#-live-demo-explore-booklore-in-action>",
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
		println(*book.Title)
	}
}
