package booklore

import (
	"context"
	"errors"
	"fmt"
	"net/http"
)

type Credentials struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func (c *Client) LoginWithCredentials(ctx context.Context, creds Credentials) error {
	resp, err := c.LoginUser(
		ctx,
		LoginUserJSONRequestBody{
			Username: StrToPtr(creds.Username),
			Password: StrToPtr(creds.Password),
		},
	)
	if err != nil {
		return err
	}
	response, err := ParseLoginUserResponse(resp)
	if err != nil {
		return err
	}
	if response.JSON200 == nil {
		return errors.New("unexpected response from server")
	}
	token, ok := (*response.JSON200)["refreshToken"]
	if !ok {
		return errors.New("no refresh token in response")
	}
	tokenFunc := func(ctx context.Context, req *http.Request) error {
		req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", token))
		return nil
	}
	c.RequestEditors = append(c.RequestEditors, tokenFunc)
	return nil
}

func (c *Client) ClientWithResponses() ClientWithResponses {
	return ClientWithResponses{c}
}

func StrToPtr(s string) *string {
	return &s
}
