package common

import (
	"fmt"
	"github.com/spf13/viper"
	"os"
)

func ReadConfig(name, suffix string) map[string]interface{} {
	path, err := os.Getwd()
	if err != nil {
		panic(err)
	}

	config := viper.New()
	config.AddConfigPath(path + "\\config") //设置读取的文件路径
	config.SetConfigName(name)              //设置读取的文件名
	config.SetConfigType(suffix)            //设置文件的类型
	//尝试进行配置读取
	if err := config.ReadInConfig(); err != nil {
		panic(fmt.Sprintf("read %#v fail err:%#v \n", name+"."+suffix, err.Error()))
	}
	return config.AllSettings()
}

type ConfigInfo struct {
	Name   string
	Suffix string
}
