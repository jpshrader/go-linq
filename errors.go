package golinq

type NotFoundError struct{}

func (e NotFoundError) Error() string {
	return "no matching element(s) found"
}