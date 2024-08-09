package entity

import "time"

// Secret - базовая информация о секретной информации
type Secret struct {
	GUID      string
	Name      string
	CreatedAt time.Time
	UpdatedAt time.Time
	Payload   *SecretPayload
}
