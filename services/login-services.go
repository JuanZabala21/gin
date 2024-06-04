package services

type LoginService interface {
	Login(username string, password string) bool
}

type loginService struct {
	authorizeUserName string
	authorizePassword string
}

func NewLoginService() LoginService {
	return &loginService{
		authorizeUserName: "Juan",
		authorizePassword: "1234",
	}
}

func (service *loginService) Login(username string, password string) bool {
	return service.authorizeUserName == username &&
		service.authorizePassword == password
}
