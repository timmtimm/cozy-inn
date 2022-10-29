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

	userData, err := uu.userRepository.GetUserByEmail(userDomain.Email)
	if err != nil {
		return "", err
	}

	token := uu.jwtAuth.GenerateToken(userData.Email, userData.Role)
	return token, nil
}

func (uu *UserUseCase) GetUserByEmail(email string) (Domain, error) {
	user, err := uu.userRepository.GetUserByEmail(email)
	if err != nil {
		return Domain{}, err
	}

	return user, nil
}

func (uu *UserUseCase) SudoUpdateUser(email string, userDomain *Domain) (Domain, error) {
	user, err := uu.userRepository.SudoUpdate(email, userDomain)
	if err != nil {
		return Domain{}, err
	}

	return user, nil
}

func (uu *UserUseCase) UpdateUser(email string, userDomain *Domain) (Domain, error) {
	user, err := uu.userRepository.Update(email, userDomain)
	if err != nil {
		return Domain{}, err
	}

	return user, nil
}

func (uu *UserUseCase) GetUserList() ([]Domain, error) {
	users, err := uu.userRepository.GetUserList()
	if err != nil {
		return nil, err
	}

	return users, nil
}
