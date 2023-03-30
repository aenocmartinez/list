package usecase

type RemoveVideoUseCase struct{}

func (useCase *RemoveVideoUseCase) Execute() (code int, err error) {
	return 200, nil
}
