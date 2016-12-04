type User struct {
	GithubUsername string
	GithubClient   github.API
}

func (u *User) RepositoriesCount() (int, error) {
	c := u.GetGithubClient()
	rs, err := c.GetRepositories()
	if err != nil {
		return 0, err
	}
	return len(rs), nil
}

func (u *User) GetGithubClient() github.API {
	if u.GithubClient != nil {
		return u.GithubClient
	} else {
		return &github.GithubClient{
			Credentials: u.GithubCredentials(),
		}
	}
}
