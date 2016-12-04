import "errors"

func Foo() error {
	return errors.New("error")
}
