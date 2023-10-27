package utils

import "github.com/google/uuid"

func GenerateUUID() string {
	// bikin uuid
	UUID := uuid.New()
	return UUID.String()
}
