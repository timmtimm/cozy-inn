package rooms

type RoomUseCase struct {
	roomRepository Repository
}

func NewRoomUsecase(rr Repository) UseCase {
	return &RoomUseCase{
		roomRepository: rr,
	}
}

func (ru *RoomUseCase) GetAllRoom() ([]Domain, error) {
	return ru.roomRepository.GetAllRoom()
}

func (ru *RoomUseCase) CreateRoom(roomDomain *Domain) error {
	err := ru.roomRepository.CreateRoom(roomDomain)
	if err != nil {
		return err
	}

	return nil
}

func (ru *RoomUseCase) UpdateRoom(roomDomain *Domain) (Domain, error) {
	roomData, err := ru.roomRepository.UpdateRoom(roomDomain)
	if err != nil {
		return Domain{}, err
	}

	return roomData, nil
}

func (ru *RoomUseCase) DeleteRoom(roomType string) error {
	err := ru.roomRepository.DeleteRoom(roomType)
	if err != nil {
		return err
	}

	return nil
}
