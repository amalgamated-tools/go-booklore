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

func (c *Client) GetBook(ctx context.Context, id int64) (*Book, error) {
	respClient := c.ClientWithResponses()
	remotebook, err := respClient.GetBookByIdWithResponse(ctx, id, nil)
	if err != nil {
		return nil, err
	}
	return remotebook.JSON200, nil
}
