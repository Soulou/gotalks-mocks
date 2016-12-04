import "testing"

func TestUser_RepositoriesCount(t *testing.T) {
	c := github.GithubClient{Credentials: testCredentials()}
	u := User{GithubClient: c}
	n, err := u.RepositoriesCount()
	// ...
}
