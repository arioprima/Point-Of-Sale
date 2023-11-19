package utils

import "github.com/google/uuid"

func GenerateUUID() string {
	generatedUUID := uuid.New()
	return generatedUUID.String()
}
