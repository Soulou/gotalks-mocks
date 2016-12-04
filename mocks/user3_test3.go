import "testing"

func TestUser_RepositoriesCount(t *testing.T) {
	c := githubmock.Client{}
	c.GetRepositoriesData = githubmock.GetRepositoriesData{
		Repositories: github.Repositories{{Name: "repo1"}, {Name: "repo2"}},
	}
	u := User{GithubClient: c}
	n, err := u.RepositoriesCount()
	if err != nil {
		t.Errorf("expecting nil err, got %v", err)
	}
	if n != 2 {
		t.Errorf("expecting 2 repositories, got %v", n)
	}
	if len(c.GetRepositoriesCalls) != 1 {
		t.Errorf("expecting github client to be called once, it wasn't")
	}
	if !c.GetRepositoriesCalls[0].StrictOwner {
		t.Errorf("expecting StrictOwner to be false on github client, it was true")
	}
}
