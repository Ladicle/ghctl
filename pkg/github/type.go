package github

import (
	"github.com/shurcooL/githubql"
)

// Client manages GitHub API requests
type Client struct {
	GQL *githubql.Client
}

// LoginQuery is GraphQL query to get username.
type LoginQuery struct {
	Viewer struct {
		Login githubql.String
	}
}
