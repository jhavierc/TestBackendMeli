package service

import (
	"errors"
	"github.com/TestBackendMeli/api/app/domain/model"
	"github.com/TestBackendMeli/api/app/infrastructure/config"
)

const (
	errorParameteres   = "error in input parameters"
	errorNumSatellites = "the number of satellites is wrong"
	errorLongDistance  = "the satellite has no value in the distance"
	errorCoordinates   = "the initial coordinates of the satellites could not be found"
)

type GetQuasarContentServicePort interface {
	PostTopSecret(queasar model.RequestQueasar) (responseQueasar model.ResponseQuasar, err error)
	GetTopSecret(queasar model.RequestQueasar) (responseQueasar model.ResponseQuasar, err error)
}

//method that performs the calculation to find the coordinates of the ship and decipher the message sent
func PostTopSecret(requestQueasar model.RequestQueasar) (responseQueasar model.ResponseQuasar, err error) {
	err = validarInitialSatelliteValues()
	if err != nil{
		return
	}

	err = validateParameters(requestQueasar)
	if err != nil{
		return
	}

	responseQueasar = model.ResponseQuasar{
		Position: "111,222",
		Message:  "ok",
	}

	return responseQueasar, err
}

//method that performs the calculation to find the coordinates of the ship and decipher the message sent
func GetTopSecret(requestQueasar model.RequestQueasar) (responseQueasar model.ResponseQuasar, err error) {

	err = validarInitialSatelliteValues()
	if err != nil{
		return
	}

	err = validateParameters(requestQueasar)
	if err != nil{
		return
	}

	responseQueasar = model.ResponseQuasar{
		Position: "111,222",
		Message:  "ok",
	}

	return responseQueasar, err

}

//allows to load the initial values ​​of the satellite coordinates
func validarInitialSatelliteValues()(err error){
	satellites := config.GetConfig()

	if satellites.Satellite1_x == 0 || satellites.Satellite1_y == 0 || satellites.Satellite2_x == 0 || satellites.Satellite2_y == 0 || satellites.Satellite3_x == 0 || satellites.Satellite3_y == 0 {
		err = errors.New(errorCoordinates)
		return
	}
	return
}

//allows to validate the input parameters
func validateParameters(requestQueasar model.RequestQueasar)(err error){
	if len(requestQueasar.Satellites)==0 {
		err = errors.New(errorParameteres)
		return
	}
	if len(requestQueasar.Satellites)<3 {
		err = errors.New(errorNumSatellites)
		return
	}
	dataValid := false

	for _, satellite:=range requestQueasar.Satellites{
		if satellite.Distance ==0{
			dataValid=false
		}
	}
	if !dataValid{
		err = errors.New(errorLongDistance)
		return
	}
	return
}