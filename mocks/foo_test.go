import "testing"

func TestFoo(t *testing.T) {
	err := Foo()
	if err != nil {
		t.Errorf("expecting nil error, got %v", err)
	}
}
