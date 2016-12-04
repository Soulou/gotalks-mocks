type API interface {
	GetRepositories() (Repositories, error)
}

type GithubClient struct {
	Credentials GithubCredentials
}

func (c *GithubClient) GetRepositories() (Repositories, error) {
	// HTTP request using Credentials
	// ...
	return repositories, nil
}
