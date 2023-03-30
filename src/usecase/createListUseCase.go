package usecase

import "abix360/src/view/dto"

type CreateListUseCase struct{}

func (useCase *CreateListUseCase) Execute(listDto dto.ListDto) (code int, err error) {
	return 201, nil
}
