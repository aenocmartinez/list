package usecase

import (
	"abix360/src/domain"
	"abix360/src/view/dto"
)

type CreateListUseCase struct{}

func (useCase *CreateListUseCase) Execute(listDto dto.ListDto) (code int, err error) {
	list := domain.NewList(listDto.Name)
	list.WithRepository(listRespository)
	err = list.Create()
	if err != nil {
		return 500, err
	}

	return 201, nil
}
