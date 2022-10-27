package users

import (
	"errors"
)

type UserUseCase struct {
	userRepository Repository
}

func NewUserUsecase(ur Repository) UseCase {
	return &UserUseCase{
		userRepository: ur,
	}
}

func (uu *UserUseCase) Register(userDomain *Domain) error {
	return uu.userRepository.Register(userDomain)
}

func (uu *UserUseCase) Login(userDomain *Domain) (string, error) {
	if uu.userRepository.Login(userDomain) != nil {
		return "", errors.New("wrong email or password")
	}

	return "token", nil
}
