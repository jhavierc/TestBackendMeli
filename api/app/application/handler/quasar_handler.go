package handler

import (
	"github.com/TestBackendMeli/api/app/domain/model"
	"github.com/TestBackendMeli/api/app/domain/service"
	"github.com/TestBackendMeli/goutils/logger"
)

const errorService = "error calling service"

type QuasarHandlerContentApplication interface {
	HandlerPostTopSecret(queasar model.RequestQueasar) (responseQueasar model.ResponseQuasar, err error)
	HandlerGetTopSecret(queasar model.RequestQueasar) (responseQueasar model.ResponseQuasar, err error)
}

type GetGlobalContentQuasarService struct {
	GetQuasarContentServicePort service.GetQuasarContentServicePort
}

func (getGlobalContentQuasarService *GetGlobalContentQuasarService) HandlerPostTopSecret(queasar model.RequestQueasar) (responseQueasar model.ResponseQuasar, err error) {
	if responseQueasar, err = getGlobalContentQuasarService.GetQuasarContentServicePort.PostTopSecret(queasar); err != nil {
		logger.Error(errorService, err)
		return
	}
	return
}

func (getGlobalContentQuasarService *GetGlobalContentQuasarService) HandlerGetTopSecret(queasar model.RequestQueasar) (responseQueasar model.ResponseQuasar, err error) {
	if responseQueasar, err = getGlobalContentQuasarService.GetQuasarContentServicePort.GetTopSecret(queasar); err != nil {
		logger.Error(errorService, err)
		return
	}
	return
}
