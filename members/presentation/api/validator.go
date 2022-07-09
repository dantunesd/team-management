package api

type Validator interface {
	Validate(content interface{}) error
}
