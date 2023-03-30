package usecase

import "abix360/src/view/dto"

type FindListByIdUseCase struct{}

func (useCase *FindListByIdUseCase) Execute() (listDto dto.ListDto, err error) {
	return listDto, nil
}
