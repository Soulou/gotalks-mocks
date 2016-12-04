import "testing"

func TestGithubClient_GetRepositories(t *testing.T) {
	setupTestServer()
	defer teardownTestServer()

	c := &GithubClient{}
	if err != nil {
		t.Errorf("expected nil, got %v", err)
	}

	// etc.
}
