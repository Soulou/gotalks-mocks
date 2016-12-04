import "testing"

func TestUser_RepositoriesCount(t *testing.T) {
	c := githubmock.Client{}
	u := User{GithubClient: c}
	n, err := u.RepositoriesCount()
	// ...
}
