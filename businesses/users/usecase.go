package users

import (
	"cozy-inn/app/middleware"
	"errors"
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

func (uu *UserUseCase) Register(userDomain *Domain) error {
	return uu.userRepository.Register(userDomain)
}

func (uu *UserUseCase) Login(userDomain *Domain) (string, error) {
	if uu.userRepository.Login(userDomain) != nil {
		return "", errors.New("wrong email or password")
	}

	userData := uu.userRepository.GetuserByEmail(userDomain.Email)
	token := uu.jwtAuth.GenerateToken(userData.Email, userData.Role)

	return token, nil
}
