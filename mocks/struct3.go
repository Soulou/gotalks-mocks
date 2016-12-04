type API interface {
	GetRepositories(GetRepositoriesOpts) (Repositories, error)
}

type GetRepositoriesOpts struct {
	StrictOwner bool
}

func (c *GithubClient) GetRepositories(opts GetRepositoriesOpts) (Repositories, error) {
	// ...
	return repositories, nil
}
