// Package entity - доменные сущности
package entity

import "time"

// Secret - информация о секрете с зашифрованными данными
type SecretEncrypted struct {
	GUID       string
	Name       string
	CreatedAt  time.Time
	UpdatedAt  time.Time
	EncPayload []byte
}
