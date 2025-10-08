package service

import (
	"github.com/1nterdigital/aka-im-discover/internal/api/util"
	"github.com/1nterdigital/aka-im-discover/internal/usecase"
	"github.com/1nterdigital/aka-im-discover/pkg/common/imapi"
)

func New(apiName string, imApiCaller imapi.CallerInterface, api *util.Api, uc usecase.UseCase) *Api {
	return &Api{
		AppName:     apiName,
		Api:         api,
		imApiCaller: imApiCaller,
		uc:          uc,
	}
}

type Api struct {
	*util.Api
	AppName     string
	imApiCaller imapi.CallerInterface
	uc          usecase.UseCase
}

func (a *Api) HealthUseCase() usecase.UseCase {
	return a.uc
}

func (a *Api) DiscoverUseCase() usecase.UseCase {
	return a.uc
}

func (a *Api) EventUseCase() usecase.UseCase {
	return a.uc
}
