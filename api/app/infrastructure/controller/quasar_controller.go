package controller

import (
	"github.com/TestBackendMeli/api/app/application/handler"
	"github.com/TestBackendMeli/api/app/domain/model"
	"github.com/TestBackendMeli/api/app/infrastructure/config"
	"github.com/TestBackendMeli/goutils/apierrors"
	"github.com/TestBackendMeli/goutils/logger"
	"github.com/gin-gonic/gin"
	"net/http"
)

const (
	typeID              = "type_id"
	startProcessLogger  = "processing insurance with id: %s"
	endingProcessLogger = "insurance processing with id %s was completed"
)

type GetQuasarController struct {
	GetQuasarApplication handler.QuasarHandlerApplication
}

//ws nivel 2
func (getQuasarController *GetQuasarController) PostTopSecret(context *gin.Context) {

	requestBody, _ := context.GetRawData()
	contentDataRequest := model.RequestQueasar{}
	if mapErr := config.MapRequestToStruct(requestBody, &contentDataRequest); mapErr != nil {
		context.Error(apierrors.NewBadRequestApiError(mapErr.Error()))
		return
	}
	logger.Infof(startProcessLogger, contentDataRequest.Satellites)
	response,err := getQuasarController.GetQuasarApplication.HandlerPostTopSecret(contentDataRequest)
	if err != nil {
		context.JSON(http.StatusBadRequest,nil)
		return
	}

	logger.Infof(endingProcessLogger, contentDataRequest.Satellites)
	context.JSON(http.StatusOK,response)
}

//ws nivel 3
func (getQuasarController *GetQuasarController) PostTopSecretSplit(context *gin.Context) {

	requestBody, _ := context.GetRawData()
	requestParam:= context.Param("satellite_name")

	contentDataRequest := model.Satellite{}
	contentDataRequest.Name =requestParam
	if mapErr := config.MapRequestToStruct(requestBody, &contentDataRequest); mapErr != nil {
		context.Error(apierrors.NewBadRequestApiError(mapErr.Error()))
		return
	}
	logger.Infof(startProcessLogger, contentDataRequest)
	response,err := getQuasarController.GetQuasarApplication.HandlerPostTopSecretSplit(contentDataRequest)
	if err != nil {
		context.JSON(http.StatusBadRequest,nil)
		return
	}

	logger.Infof(endingProcessLogger, contentDataRequest)
	context.JSON(http.StatusOK,response)
}
