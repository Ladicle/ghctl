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

// GetLogin login to GitHub and returns username.
func (c *Client) GetLogin() (string, error) {
	var q struct {
		Viewer struct {
			Login githubql.String
		}
	}
	err := c.GQL.Query(context.Background(), &q, nil)
	return string(q.Viewer.Login), err
}
}
