package services

type Auth struct {
}

func GetUserService() *Auth {
	return new(Auth)
}

func (a Auth) Login(user string, password string) bool {

	if user == "tamas" && password == "tamas" {
		return true
	}

	return false
}
