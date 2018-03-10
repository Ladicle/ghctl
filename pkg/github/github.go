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
func (c *Client) GetRepository(org, repo string) (*Repository, error) {
	var q struct {
		Repository struct {
			Name             githubql.String
			ID               githubql.ID
			RepositoryTopics struct {
				Nodes []struct {
					Topic struct {
						Name githubql.String
					}
				}
			} `graphql:"repositoryTopics(first: $nodeN)"`
		} `graphql:"repository(owner: $owner, name: $name)"`
	}

	param := map[string]interface{}{
		"owner": githubql.String(org),
		"name":  githubql.String(repo),
		"nodeN": githubql.Int(maxNode),
	}

	err := c.GQL.Query(context.Background(), &q, param)
	if err != nil {
		return nil, err
	}

	var topics []string
	for _, n := range q.Repository.RepositoryTopics.Nodes {
		topics = append(topics, string(n.Topic.Name))
	}
	return &Repository{
		Name:   string(q.Repository.Name),
		ID:     fmt.Sprintf("%v", q.Repository.ID),
		Topics: topics,
	}, nil
}

// GetOrganization returns organization data.
func (c *Client) GetOrganization(org string) (*Organization, error) {
	var q struct {
		RepositoryOwner struct {
			Repositories struct {
				Nodes []struct {
					Name             githubql.String
					ID               githubql.ID
					RepositoryTopics struct {
						Nodes []struct {
							Topic struct {
								Name githubql.String
							}
						}
					} `graphql:"repositoryTopics(first: $nodeN)"`
				}
				PageInfo struct {
					EndCursor   githubql.String
					HasNextPage githubql.Boolean
				}
			} `graphql:"repositories(first: $nodeN, after: $after)"`
		} `graphql:"repositoryOwner(login: $login)"`
	}

	var repos []Repository
	lastID := (*githubql.String)(nil)
	for {
		param := map[string]interface{}{
			"nodeN": githubql.Int(maxNode),
			"login": githubql.String(org),
			"after": lastID,
		}

		err := c.GQL.Query(context.Background(), &q, param)
		if err != nil {
			return nil, err
		}

		for _, node := range q.RepositoryOwner.Repositories.Nodes {
			var topics []string
			for _, n := range node.RepositoryTopics.Nodes {
				topics = append(topics, string(n.Topic.Name))
			}
			repos = append(repos, Repository{
				Name:   string(node.Name),
				ID:     fmt.Sprintf("%v", node.ID),
				Topics: topics,
			})
		}

		if !q.RepositoryOwner.Repositories.PageInfo.HasNextPage {
			break
		}
		lastID = githubql.NewString(q.RepositoryOwner.Repositories.PageInfo.EndCursor)
	}

	return &Organization{
		Name:         org,
		Repositories: repos,
	}, nil
}
