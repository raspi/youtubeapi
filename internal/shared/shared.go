package shared

var _ error = NotFound{}

type NotFound struct {
}

func (n NotFound) Error() string {
	return `not found`
}
