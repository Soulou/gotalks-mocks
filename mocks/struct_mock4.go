import "sync"

type Client struct {
	*sync.Mutex
	getRepositoriesData  GetRepositoriesData
	GetRepositoriesCalls []github.GetRepositoriesOpts
}

func NewClient() *Client {
	return &Client{Mutex: &sync.Mutex{}}
}

func (c *Client) SetGetRepositoriesData(data GetRepositoriesData) {
	c.Lock()
	defer c.Unlock()
	c.getRepositoriesData = data
}

func (c *Client) GetRepositories(opts github.GetRepositoriesOpts) (github.Repositories, error) {
	c.Lock()
	defer c.Unlock()
	c.GetRepositoriesCalls = append(c.GetRepositoriesCalls, opts)
	return c.GetRepositoriesData.Repositories, c.GetRepositoriesData.Err
}
