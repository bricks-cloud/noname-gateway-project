package common

type NotFoundError interface {
	Error() string
	NotFound()
}
