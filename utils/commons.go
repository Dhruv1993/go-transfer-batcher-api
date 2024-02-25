package utils

import (
	"github.com/gofrs/uuid"
)

// generates a uuid
func NewUUID() string {
	id, _ := uuid.NewV4()
	return id.String()
}