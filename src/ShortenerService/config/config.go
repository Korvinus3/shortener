package config

import (
	"fmt"
	"github.com/koding/multiconfig"
)

//Config env conf
var Config *configuration


type configuration struct {
	HTTPPort int `default:"8083"`
	DataBase struct{
		Name string `default:"shortener.db"`
		URLsBucketName string `default:"shortener"`
	}
}

func init() {

	var mc *multiconfig.DefaultLoader

	mc = multiconfig.New()

	Config = new(configuration)

	mc.MustLoad(Config)

	fmt.Printf("Loaded config  %+v\n", Config)

}