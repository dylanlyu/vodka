package uuid

import "github.com/google/uuid"

func PointerUUID(input uuid.UUID) *uuid.UUID { return &input }

func CreatedUUIDString() string {
	return uuid.New().String()
}

func ValidateUUID(input string) (uuid.UUID, error) {
	return uuid.Parse(input)
}

func StringToUUID(input string) uuid.UUID {
	return uuid.MustParse(input)
}
