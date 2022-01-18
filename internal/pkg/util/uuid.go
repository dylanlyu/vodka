package util

import (
	"github.com/google/uuid"
)

func GenerateUUID() string {
	return uuid.New().String()
}

func ValidateUUID(input string) (uuid.UUID, error) {
	return uuid.Parse(input)
}
