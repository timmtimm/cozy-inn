package users

import (
	"cozy-inn/app/middleware"
	"errors"
	"time"
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

func (uu *UserUseCase) GetUserByEmail(email string) (Domain, error) {
	user, err := uu.userRepository.GetUserByEmail(email)
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

func (uu *UserUseCase) UserRegister(userInput Domain) (string, error) {
	err := uu.userRepository.Register(userInput)

	if err != nil {
		return "", err
	}

	token := uu.jwtAuth.GenerateToken(userInput.Email, userInput.Role)
	return token, nil
}

func (uu *UserUseCase) AdminRegister(userInput Domain) error {
	avaliableRoles := []string{"user", "receptionist"}
	roleFound := false
	for _, avaliableRole := range avaliableRoles {
		if userInput.Role == avaliableRole {
			roleFound = true
			break
		}
	}

	if !roleFound {
		return errors.New("invalid role")
	}

	err := uu.userRepository.Register(userInput)
	if err != nil {
		return err
	}

	return nil
}

func (uu *UserUseCase) Login(userInput Domain) (string, error) {
	if err := uu.userRepository.Login(userInput); err != nil {
		return "", err
	}

	userData, err := uu.userRepository.GetUserByEmail(userInput.Email)
	if err != nil {
		return "", err
	}

	token := uu.jwtAuth.GenerateToken(userData.Email, userData.Role)
	return token, nil
}

func (uu *UserUseCase) UserUpdate(email string, userInput Domain) (Domain, error) {
	user, err := uu.userRepository.GetUserByEmail(email)
	if err != nil {
		return Domain{}, err
	}

	user.Name = userInput.Name
	user.ImageID_URL = userInput.ImageID_URL
	user.UpdatedAt = time.Now()

	err = uu.userRepository.Update(email, user)
	if err != nil {
		return Domain{}, err
	}

	return user, nil
}

func (uu *UserUseCase) AdminUpdate(email string, userInput Domain) (Domain, error) {
	avaliableRoles := []string{"user", "receptionist"}
	roleFound := false
	for _, avaliableRole := range avaliableRoles {
		if userInput.Role == avaliableRole {
			roleFound = true
			break
		}
	}

	if !roleFound {
		return Domain{}, errors.New("invalid role")
	}

	if userInput.Role == "admin" {
		return Domain{}, errors.New("can't change role to admin")
	}

	user, err := uu.userRepository.GetUserByEmail(email)
	if err != nil {
		return Domain{}, err
	}

	user.Role = userInput.Role
	user.Status = userInput.Status
	user.UpdatedAt = time.Now()

	err = uu.userRepository.Update(email, user)
	if err != nil {
		return Domain{}, err
	}

	return user, nil
}

func (uu *UserUseCase) AdminDelete(email string) error {
	err := uu.userRepository.Delete(email)
	if err != nil {
		return err
	}

	return nil
}
