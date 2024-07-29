package service

type User struct{}

func (u User) Register(username, password string) error {}

func (u User) Login(username, password string) error {}
