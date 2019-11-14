package service

type UserServiceImpl struct {
}

var UserService *UserServiceImpl

func init() {
	UserService = &UserServiceImpl{}
}
