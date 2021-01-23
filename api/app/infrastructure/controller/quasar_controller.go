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

type GetGlobalQuasarController struct {
	GetGlobalQuasarApplication handler.QuasarHandlerContentApplication
}

//ws
func (getGlobalQuasarApplication *GetGlobalQuasarController) PostTopSecret(context *gin.Context) {

	requestBody, _ := context.GetRawData()
	contentDataRequest := model.RequestQueasar{}
	if mapErr := config.MapRequestToStruct(requestBody, &contentDataRequest); mapErr != nil {
		context.Error(apierrors.NewBadRequestApiError(mapErr.Error()))
		return
	}
	logger.Infof(startProcessLogger, contentDataRequest.Satellites)
	response,err := getGlobalQuasarApplication.GetGlobalQuasarApplication.HandlerPostTopSecret(contentDataRequest)
	if err != nil {
		context.Error(err)
		return
	}

	logger.Infof(endingProcessLogger, contentDataRequest.Satellites)
	context.JSON(http.StatusOK,response)
}

//get pendiente cast json a model
func (getGlobalQuasarApplication *GetGlobalQuasarController) GetTopSecret(context *gin.Context) {

	requestBody, _ := context.GetRawData()

	//requestParam:= context.Query("value")

	contentDataRequest := model.RequestQueasar{}
	if mapErr := config.MapRequestToStruct(requestBody, &contentDataRequest); mapErr != nil {
		context.Error(apierrors.NewBadRequestApiError(mapErr.Error()))
		return
	}
	logger.Infof(startProcessLogger, contentDataRequest.Satellites)
	response,err := getGlobalQuasarApplication.GetGlobalQuasarApplication.HandlerGetTopSecret(contentDataRequest)
	if err != nil {
		context.Error(err)
		return
	}

	logger.Infof(endingProcessLogger, contentDataRequest.Satellites)
	context.JSON(http.StatusOK,response)
}
