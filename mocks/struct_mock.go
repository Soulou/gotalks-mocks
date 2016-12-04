package githubmock

type Client struct{}

func (c *Client) GetRepositories() (github.Repositories, error) {
	return github.Repositories{}, nil
}
