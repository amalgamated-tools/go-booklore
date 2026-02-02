package client

// AccessToken returns the access token from the HTTP response
func (r LoginUserResponse) AccessToken() string {
	if r.JSON200 != nil {
		// JSON200 is a *map[string]string
		if token, ok := (*r.JSON200)["accessToken"]; ok {
			return token
		}
	}
	return ""
}

// RefreshToken returns the refresh token from the HTTP response
func (r LoginUserResponse) RefreshToken() string {
	if r.JSON200 != nil {
		// JSON200 is a *map[string]string
		if token, ok := (*r.JSON200)["refreshToken"]; ok {
			return token
		}
	}
	return ""
}
