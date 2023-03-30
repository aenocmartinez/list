package usecase

type AddVideoUseCase struct{}

func (useCase *AddVideoUseCase) Execute() (code int, err error) {
	return 201, nil
}
