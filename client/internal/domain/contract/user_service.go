package contract

type UserService interface {
	Register(login string, password string) error
	Login(login string, password string) (string, error)
}
