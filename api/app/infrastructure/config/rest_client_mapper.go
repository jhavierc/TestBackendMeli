package config

import (
	"encoding/json"
	"github.com/TestBackendMeli/goutils/logger"
	"github.com/go-playground/validator"
	"github.com/pkg/errors"
	"reflect"
)

const (
	JSONParserError = "the body of the petition is incorrect, check and try again"
)

var (
	//config   = &validator {TagName: "validate"}
	validate = validator.New()
)

func MapRequestToStruct(payload []byte, genericStruct interface{}) error {
	jsonErr := json.Unmarshal(payload, genericStruct)
	if jsonErr != nil {
		logger.Error("MapRequestToStruct - jsonErr "+jsonErr.Error(), jsonErr)
		return errors.New(JSONParserError)
	}

	value := reflect.ValueOf(genericStruct)
	indirectValue := reflect.Indirect(value)

	if indirectValue.Kind() == reflect.Slice || indirectValue.Kind() == reflect.Map {
		return nil
	}

	return validateStruct(genericStruct)
}

func validateStruct(genericStruct interface{}) error {
	return validate.Struct(genericStruct)
}
