package validator_port

type Validator interface {
	Validate(field interface{}) []string
}
