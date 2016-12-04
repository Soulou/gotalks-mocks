import "testing"

func TestUser_RepositoriesCount(t *testing.T) {
	u := User{}
	n, err := u.RepositoriesCount()
	// ...
}
