package handler

import (
	"github.com/TestBackendMeli/api/app/domain/model"
	"github.com/TestBackendMeli/api/app/domain/service"
	"github.com/TestBackendMeli/goutils/logger"
)

const errorService = "error calling service"

type QuasarHandlerApplication interface {
	HandlerPostTopSecret(queasar model.RequestQueasar) (responseQueasar model.ResponseQuasar, err error)
	HandlerPostTopSecretSplit(satellite model.Satellite) (responseQueasar model.ResponseQuasar, err error)
}

type QuasarHandler struct {
	GetQuasarServicePort service.GetQuasarServicePort
}

func (quasarHandler *QuasarHandler) HandlerPostTopSecret(queasar model.RequestQueasar) (responseQuasar model.ResponseQuasar, err error) {
	if responseQuasar, err = quasarHandler.GetQuasarServicePort.PostTopSecret(queasar); err != nil {
		logger.Error(errorService, err)
		return
	}
	return
}

func (quasarHandler *QuasarHandler) HandlerPostTopSecretSplit(satellite model.Satellite) (responseQuasar model.ResponseQuasar, err error) {
	if responseQuasar, err = quasarHandler.GetQuasarServicePort.PostTopSecretSplit(satellite); err != nil {
		logger.Error(errorService, err)
		return
	}
	return
}


