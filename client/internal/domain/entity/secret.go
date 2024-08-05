// Package entity - доменные сущности
package entity

import "time"

// Secret - базовая информация о секретной информации
type Secret struct {
	GUID       string
	Name       string
	Metadata   string
	CreatedAt  time.Time
	UpdatedAt  time.Time
	EncPayload []byte
}
