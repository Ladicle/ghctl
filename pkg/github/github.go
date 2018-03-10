package github

import (
	"context"
	"fmt"

	"github.com/shurcooL/githubql"
	"golang.org/x/oauth2"
)

const maxNode = 100

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

// GetRepository returns repository data.
func (c *Client) GetRepository(orgName, repoName string) (*Repository, error) {
	var q struct {
		Repository struct {
			Name   githubql.String
			ID     githubql.ID
			Labels struct {
				Nodes []struct {
					Name githubql.String
				}
			} `graphql:"labels(first: $maxNode)"`
		} `graphql:"repository(owner: \"$orgName\", name: \"$repoName\")"`
	}

	err := c.GQL.Query(context.Background(), &q, nil)
	if err != nil {
		return nil, err
	}

	var labels []string
	for _, ln := range q.Repository.Labels.Nodes {
		labels = append(labels, string(ln.Name))
	}
	return &Repository{
		Name:   string(q.Repository.Name),
		ID:     fmt.Sprintf("%v", q.Repository.ID),
		Labels: labels,
	}, nil
}

// GetOrganization returns organization data.
func (c *Client) GetOrganization(orgName string) (*Organization, error) {
	cursor := ""
	total := 0
	org := Organization{Name: orgName}

	// cursor is used in query
	_ = cursor

	for {
		var q struct {
			RepositoryOwner struct {
				Repositories struct {
					TotalCount githubql.Int
					Nodes      []struct {
						Name   githubql.String
						ID     githubql.ID
						Labels struct {
							Nodes []struct {
								Name githubql.String
							}
						} `graphql:"labels(first: $maxNodes)"`
					}
				} `graphql:"repositories(first: $maxNode$cursor)"`
			} `graphql:"repositoryOwner(login: \"$orgName\")"`
		}

		err := c.GQL.Query(context.Background(), &q, nil)
		if err != nil {
			return nil, err
		}

		for _, node := range q.RepositoryOwner.Repositories.Nodes {
			var labels []string
			for _, ln := range node.Labels.Nodes {
				labels = append(labels, string(ln.Name))
			}
			repo := Repository{
				Name:   string(node.Name),
				ID:     fmt.Sprintf("%v", node.ID),
				Labels: labels,
			}
			org.Repositories = append(org.Repositories, repo)
		}

		total += maxNode
		if total >= int(q.RepositoryOwner.Repositories.TotalCount) {
			break
		}
		lastID := org.Repositories[len(org.Repositories)-1].ID
		cursor = fmt.Sprintf(", after: \"%s\"", lastID)
	}

	return &org, nil
}
