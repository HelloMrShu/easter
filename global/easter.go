package global

import (
	"flag"
	"fmt"
	"github.com/spf13/viper"
)

func Initialize() {
	env := flag.String("e", "dev", "env info")
	// init config
	InitConfig(*env)
	// init log
	InitLogger()
	// init db
	//InitDB()
	Logger.Info("start server successfully!")
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
