package validator_port

import "laughing-succostash/internal/core/domain"

type UserValidator interface {
	Validate(user *domain.User) []string
	RegisterTranslation() error
	RegisterTagRequired() error
}
