package usecase

type DeleteListUseCase struct{}

func (useCase *DeleteListUseCase) Execute() (code int, err error) {
	return 200, nil
}
