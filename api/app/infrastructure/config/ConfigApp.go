package config

import (
	"github.com/tkanos/gonfig"
)
type Satelletes struct {
	Satellite1_x float32
	Satellite1_y float32
	Satellite2_x float32
	Satellite2_y float32
	Satellite3_x float32
	Satellite3_y float32
}
func GetConfig() Satelletes {
	configuration := Satelletes{}

	fileName := "./config.json"
	gonfig.GetConf(fileName, &configuration)
	return configuration
}
