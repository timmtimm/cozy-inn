package users

import (
	"cozy-inn/app/middleware"
)

type UserUseCase struct {
	userRepository Repository
	jwtAuth        *middleware.ConfigJWT
}

func NewUserUsecase(ur Repository, jwtAuth *middleware.ConfigJWT) UseCase {
	return &UserUseCase{
		userRepository: ur,
		jwtAuth:        jwtAuth,
	}
}

func (uu *UserUseCase) Register(userDomain *Domain) (string, error) {
	err := uu.userRepository.Register(userDomain)

	if err != nil {
		return "", err
	}

	token := uu.jwtAuth.GenerateToken(userDomain.Email, userDomain.Role)

	return token, nil
}

func (uu *UserUseCase) Login(userDomain *Domain) (string, error) {
	if err := uu.userRepository.Login(userDomain); err != nil {
		return "", err
	}

	userData := uu.userRepository.GetUserByEmail(userDomain.Email)
	token := uu.jwtAuth.GenerateToken(userData.Email, userData.Role)

	return token, nil
}
