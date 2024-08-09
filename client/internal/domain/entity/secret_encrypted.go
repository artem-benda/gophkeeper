// Package entity - доменные сущности
package entity

import "time"

// Secret - базовая информация о секретной информации
type SecretEncrypted struct {
	GUID       string
	Name       string
	CreatedAt  time.Time
	UpdatedAt  time.Time
	EncPayload []byte
}
