package main

import (
	"abix360/src/usecase"
	"abix360/src/view/dto"
	"log"
)

func main() {

	list := dto.ListDto{
		Name: "Exaltate - Marcos Witt",
	}

	useCase := usecase.CreateListUseCase{}
	_, err := useCase.Execute(list)
	if err != nil {
		log.Println("Err: ", err)
		return
	}

	log.Println(list)
}
