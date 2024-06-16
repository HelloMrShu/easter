package global

import (
	"fmt"
	"github.com/spf13/viper"
)

func Initialize(env string) {
	// init config
	InitConfig(env)
	// init log
	InitLogger()
	// init db
	//InitDB()
}

func InitConfig(env string) {
	fileName := fmt.Sprintf("config/%s.yaml", env)

	v := viper.New()
	v.SetConfigFile(fileName)
	v.SetConfigType("yaml")

	if err := v.ReadInConfig(); err != nil {
		panic(err)
	}
	if err := v.Unmarshal(&ServerConfig); err != nil {
		panic(err)
	}
}
