package config

import (
	"github.com/tkanos/gonfig"
)
type Satelletes struct {
	A1 float64
	B1 float64
	A2 float64
	B2 float64
	A3 float64
	B3 float64
}
func GetConfig() Satelletes {
	configuration := Satelletes{}

	fileName := "./config.json"
	gonfig.GetConf(fileName, &configuration)
	return configuration
}
