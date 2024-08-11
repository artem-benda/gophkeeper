package entity

// User доменная модель пользователя
type User struct {
	ID           int64
	Login        string
	PasswordHash string
}
