package users

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
