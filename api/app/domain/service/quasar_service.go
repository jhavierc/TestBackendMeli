package service

import (
	"errors"
	"fmt"
	"github.com/jhavierc/TestBackendMeli/api/app/domain/model"
	"github.com/jhavierc/TestBackendMeli/api/app/infrastructure/config"
	"github.com/jhavierc/TestBackendMeli/goutils/logger"
	"math"
)

const (
	errorParameteres   = "error in input parameters"
	errorNumSatellites = "the number of satellites is wrong"
	errorLongDistance  = "the satellite has no value in the distance"
	errorCoordinates   = "the initial coordinates of the satellites could not be found"
)

type GetQuasarServicePort interface {
	PostTopSecret(queasar model.RequestQueasar) (responseQueasar model.ResponseQuasar, err error)
	PostTopSecretSplit(satellite model.Satellite) (responseQueasar model.ResponseQuasar, err error)
}

type GetQuasarContentService struct{}

//method that performs the calculation to find the coordinates of the ship and decipher the message sent
func (getQuasarContentService *GetQuasarContentService) PostTopSecret(requestQueasar model.RequestQueasar) (responseQueasar model.ResponseQuasar, err error) {
	err = validarInitialSatelliteValues()
	if err != nil {
		return
	}

	err = validateParameters(requestQueasar)
	if err != nil {
		return
	}

	//llamada a la implementeacion al otro metodo
	//position := calculatePositionEmisorMessage(requestQueasar)

	x, y := GetLocation(requestQueasar.Satellites[0].Distance, requestQueasar.Satellites[1].Distance, requestQueasar.Satellites[2].Distance)
	position := model.Position{X: x, Y: y}

	message := GetMessage(requestQueasar.Satellites[0].Message, requestQueasar.Satellites[1].Message, requestQueasar.Satellites[2].Message)

	responseQueasar = model.ResponseQuasar{
		Position: position,
		Message:  message,
	}

	return responseQueasar, err
}

//method that performs the calculation to find the coordinates of the ship and decipher the message sent only satelly
func (getQuasarContentService *GetQuasarContentService) PostTopSecretSplit(satellite model.Satellite) (responseQueasar model.ResponseQuasar, err error) {
	err = validarInitialSatelliteValues()
	if err != nil {
		return
	}

	err = validateParametersSplit(satellite)
	if err != nil {
		return
	}

	//llamada a la implementeacion original
	//position := calculatePositionEmisorMessage(requestQueasar)

	x, y := GetLocation(satellite.Distance)
	position := model.Position{X: x, Y: y}

	message := GetMessage(satellite.Message)

	responseQueasar = model.ResponseQuasar{
		Position: position,
		Message:  message,
	}

	return responseQueasar, err
}


//allows to load the initial values ​​of the satellite coordinates
func validarInitialSatelliteValues() (err error) {
	satellites := config.GetConfig()
	if satellites.A1 == 0 || satellites.B1 == 0 || satellites.A2 == 0 || satellites.B2 == 0 || satellites.A3 == 0 || satellites.B3 == 0 {
		err = errors.New(errorCoordinates)
		return
	}
	return
}

//allows to validate the input parameters
func validateParameters(requestQueasar model.RequestQueasar) (err error) {
	if len(requestQueasar.Satellites) == 0 {
		err = errors.New(errorParameteres)
		return
	}
	if len(requestQueasar.Satellites) < 3 {
		err = errors.New(errorNumSatellites)
		return
	}
	dataValid := true

	for _, satellite := range requestQueasar.Satellites {
		logger.Info(satellite.Name + " - " + fmt.Sprint(satellite.Distance))
		if satellite.Distance == 0 {
			dataValid = false
		}
	}
	if !dataValid {
		err = errors.New(errorLongDistance)
		return
	}
	return
}

func validateParametersSplit(satellite model.Satellite) (err error) {
	dataValid := true
	logger.Info(satellite.Name + " - " + fmt.Sprint(satellite.Distance))
	if satellite.Distance == 0 {
		dataValid = false
	}
	if !dataValid {
		err = errors.New(errorLongDistance)
		return
	}
	return
}

func GetLocation(distances ...float64) (x, y float64) {
	distance1 := float64(0)
	distance2 := float64(0)
	distance3 := float64(0)

	for index, distance := range distances {
		if index == 0 {
			distance1 = distance
		} else if index == 1 {
			distance2 = distance
		} else {
			distance3 = distance
		}
	}

	s := config.GetConfig()

	op1 := (((s.A2 - s.A1) * (math.Pow(s.A3, 2) + math.Pow(s.B3, 2) - math.Pow(distance3, 2))) + ((s.A1 - s.A3) * (math.Pow(s.A2, 2) + math.Pow(s.B2, 2) - math.Pow(distance2, 2))) +
		((s.A3 - s.A2) * (math.Pow(s.A1, 2) + math.Pow(s.B1, 2) - math.Pow(distance1, 2))))

	op2 := 2 * ((s.B3 * (s.A2 - s.A1)) + (s.B2 * (s.A1 - s.A3)) + (s.B1 * (s.A3 - s.A2)))

	y = op1 / op2

	op3 := math.Pow(distance2, 2) + math.Pow(s.A1, 2) + math.Pow(s.B1, 2) - math.Pow(distance1, 2) - math.Pow(s.A2, 2) - math.Pow(s.B2, 2) - (2 * (s.B1 - s.B2) * y)

	op4 := 2 * (s.A1 - s.A2)

	x = op3 / op4

	return
}

//decryp message
func GetMessage(messages ...[]string) (msg string) {
	var m = map[int]string{}
	cont :=0
	for _, message := range messages {
		//fmt.Println("---------------------")
		if len(message)>cont{
			cont = len(message)
		}
		for i := 0; i < len(message); i++ {
			if len(m[i]) == 0 {
				m[i] = message[i]
			}

		}
	}
	for i :=0; i<cont; i++{
		if len(msg) == 0 {
			msg = msg + m[i]
		} else {
			msg = msg + " " + m[i]
		}
	}
	return
}

//Trilateration concept
func calculatePositionEmisorMessage(requestQueasar model.RequestQueasar) (position model.Position) {
	s := config.GetConfig()

	op1 := (((s.A2 - s.A1) * (math.Pow(s.A3, 2) + math.Pow(s.B3, 2) - math.Pow(requestQueasar.Satellites[2].Distance, 2))) + ((s.A1 - s.A3) * (math.Pow(s.A2, 2) + math.Pow(s.B2, 2) - math.Pow(requestQueasar.Satellites[1].Distance, 2))) +
		((s.A3 - s.A2) * (math.Pow(s.A1, 2) + math.Pow(s.B1, 2) - math.Pow(requestQueasar.Satellites[0].Distance, 2))))

	op2 := 2 * ((s.B3 * (s.A2 - s.A1)) + (s.B2 * (s.A1 - s.A3)) + (s.B1 * (s.A3 - s.A2)))

	y := op1 / op2

	op3 := math.Pow(requestQueasar.Satellites[1].Distance, 2) + math.Pow(s.A1, 2) + math.Pow(s.B1, 2) - math.Pow(requestQueasar.Satellites[0].Distance, 2) - math.Pow(s.A2, 2) - math.Pow(s.B2, 2) - (2 * (s.B1 - s.B2) * y)

	op4 := 2 * (s.A1 - s.A2)

	x := op3 / op4

	position = model.Position{X: x, Y: y}

	return
}

func isWordExists(words map[string]string, word string) bool {
	_, ok := words[word]
	return ok
}
