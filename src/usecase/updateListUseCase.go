package usecase

import "abix360/src/view/dto"

type UpdateListUseCase struct{}

func (useCase *UpdateListUseCase) Execute(listDto dto.ListDto) (code int, err error) {
	return 200, nil
}
