package github

import (
	"context"

	"github.com/shurcooL/githubql"
	"golang.org/x/oauth2"
)

// NewClient creates githubql client.
func NewClient(token string) *githubql.Client {
	src := oauth2.StaticTokenSource(&oauth2.Token{AccessToken: token})
	c := oauth2.NewClient(context.Background(), src)
	return githubql.NewClient(c)
}
