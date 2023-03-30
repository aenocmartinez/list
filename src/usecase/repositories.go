package usecase

import (
	"abix360/src/dao"
	"abix360/src/domain"
)

var listRespository domain.ListRepository = dao.NewListDao()
