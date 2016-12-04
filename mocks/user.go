type User struct {
	GithubUsername string
}

func (u *User) RepositoriesCount() (int, error) {
	c := &github.GithubClient{
		Credentials: u.GithubCredentials(),
	}
	rs, err := c.GetRepositories()
	if err != nil {
		return 0, err
	}
	return len(rs), nil
}
