type Client struct {
	GetRepositoriesData
	GetRepositoriesCalls []github.GetRepositoriesOpts
}

func (c *Client) GetRepositories(opts github.GetRepositoriesOpts) (github.Repositories, error) {
	c.GetRepositoriesCalls = append(c.GetRepositoriesCalls, opts)
	return c.GetRepositoriesData.Repositories, c.GetRepositoriesData.Err
}
