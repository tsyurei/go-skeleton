package service

type AuthServiceImpl struct {
}

var AuthService *AuthServiceImpl

func init() {
	AuthService = &AuthServiceImpl{}
}
