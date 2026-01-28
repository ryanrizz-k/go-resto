package rest

import (
	"github.com/rynr00/go-resto/internal/usecase/resto"
)

type handler struct {
	restoUsecase resto.Usecase
}

func NewHandler(restoUsecase resto.Usecase) *handler {
	return &handler{
		restoUsecase: restoUsecase,
	}
}
