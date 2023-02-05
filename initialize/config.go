package initialize

import (
	. "code.sohuno.com/sky/robin/global"
	"flag"
	"fmt"
	"github.com/spf13/viper"
)

func Init() {
	env := flag.String("e", "dev", "env info")
	InitLogger()
	//初始化配置
	InitConfig(*env)
	//初始化数据库
	InitDB()
}

func InitConfig(env string) {
	fileName := fmt.Sprintf("config/%s.yaml", env)
	v := viper.New()
	v.SetConfigFile(fileName)
	v.SetConfigType("yaml") //设置文件的类型
	if err := v.ReadInConfig(); err != nil {
		panic(err)
	}
	if err := v.Unmarshal(&ServerConfig); err != nil {
		panic(err)
	}

	Logger.Info("load config success")
}
