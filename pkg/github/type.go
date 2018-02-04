package github

import "github.com/shurcooL/githubql"

// LoginQuery is GraphQL query to get username.
type LoginQuery struct {
	Viewer struct {
		Login githubql.String
	}
}
