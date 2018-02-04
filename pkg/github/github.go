package github

import (
	"context"

	"github.com/shurcooL/githubql"
	"golang.org/x/oauth2"
)

// NewClient creates githubql client.
func NewClient(token string) *Client {
	src := oauth2.StaticTokenSource(&oauth2.Token{AccessToken: token})
	c := oauth2.NewClient(context.Background(), src)
	return &Client{
		GQL: githubql.NewClient(c),
	}
}

// comment
func (c *Client) GetLogin() (*LoginQuery, error) {
	q := LoginQuery{}
	err := c.GQL.Query(context.Background(), &q, nil)
	return &q, err
}
