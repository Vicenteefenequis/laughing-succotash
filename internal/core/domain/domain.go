package domain

import "github.com/google/uuid"

func generateUid() string {
	return uuid.New().String()
}
