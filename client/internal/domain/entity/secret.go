package entity

import "time"

// Secret - информация о секрете с расшифрованными данными
type Secret struct {
	GUID      string
	Name      string
	CreatedAt time.Time
	UpdatedAt time.Time
	Payload   *SecretPayload
}
