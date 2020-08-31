package config

import (
	"encoding/json"
	"github.com/pkg/errors"
	"io/ioutil"
)
// 配置文件  获取
type LogConfig struct {
	LogPath string `json:"log_path"`
	LogLevel string `json:"log_level"`
}

type Config struct {
	LogConfig LogConfig `json:"log_config"`
	DBConfig DBConfig `json:"db_config"`
	RedisConfig RedisConfig `json:"redis_config"`
}



// 数据库 配置
type DBConfig struct {
	DbHost     string `json:"db_host"`
	DbPort     string `json:"db_port"`
	DbUser     string `json:"db_user"`
	DbPassword string `json:"db_password"`
	DbName     string `json:"db_name"`
}

//redis配置
type RedisConfig struct {
	Addr		string	`json:"addr"`
	Password	string	`json:"password"`
	DB 			int	`json:"db"`
}

//解析json 处理
var conf Config
// 初始化配置文件
func InitConfig(configPath string) error  {
	configFile, err := ioutil.ReadFile(configPath)
	if err != nil {
		err = errors.Wrap(err, "读取配置文件失败")
		return err
	}
	err = json.Unmarshal(configFile, &conf)
	if err != nil {
		err = errors.Wrap(err,"删除配置文件失败。")
	}
	return nil
}

func GetConfig() Config{
	return conf
}
