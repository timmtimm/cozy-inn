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
