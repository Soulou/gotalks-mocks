type GetRepositoriesData struct {
	Repositories github.Repositories
	Err          error
}

type Client struct {
	GetRepositoriesData
}

func (c *Client) GetRepositories() (github.Repositories, error) {
	return c.GetRepositoriesData.Repositories, c.GetRepositoriesData.Err
}
