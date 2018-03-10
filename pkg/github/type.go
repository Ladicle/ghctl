package github

import (
	"github.com/shurcooL/githubql"
)

// Client manages GitHub API requests
type Client struct {
	GQL *githubql.Client
}

// Organization represents GitHub organization data.
type Organization struct {
	Name         string
	Repositories []Repository
}

// Repository represents GitHub repository data.
type Repository struct {
	Name   string
	ID     string
	Labels []string
}
