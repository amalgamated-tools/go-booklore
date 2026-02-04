package booklore

import "context"

func (c *Client) ListAllBooks(ctx context.Context) ([]Book, error) {
	respClient := c.ClientWithResponses()
	remotebooks, err := respClient.ListBooksWithResponse(ctx, nil)
	if err != nil {
		return nil, err
	}
	books := append([]Book{}, *remotebooks.JSON200...)
	return books, nil
}
