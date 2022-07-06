package utils

type Validator interface {
	Validate(content interface{}) error
}
