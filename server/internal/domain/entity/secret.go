package entity

import (
	"time"
)

// Secret - доменная модель секрета
type Secret struct {
	ID         int64
	GUID       string
	Name       string
	EncPayload []byte
	CreatedAt  time.Time
	UpdatedAt  time.Time
}
