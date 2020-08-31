package main

import (
	"ads.cost.com/config"
	"ads.cost.com/db"
	"ads.cost.com/logger"
	"ads.cost.com/redis"
	"flag"
	"fmt"
	"os"
)

func main() {
	var configPath string
	flag.StringVar(&configPath, "config", "", "配置文件路径") //后续配置
	flag.Parse()

	if configPath == "" {
		fmt.Printf("必须指定配置路径");
		os.Exit(300)
	}

	var err error
	err = config.InitConfig(configPath)
	if err != nil {
		fmt.Printf("初始化配置文件失败,错误信息:%v", err)
		os.Exit(1)
	}

	logConfig := config.GetConfig().LogConfig
	err = logger.InitLogger(logConfig.LogPath, logConfig.LogLevel)
	if err != nil {
		fmt.Printf("初始化logger失败. 错误信息: %v", err)
		os.Exit(1)
	}

	logger.GetLogger().Info("初始化成功")

	err = db.InitEngine()
	if err != nil {
		fmt.Printf("数据库连接失败 %v", err)
		os.Exit(1)
	}

	err = redis.InitRedis()
	if err != nil {
		fmt.Printf("初始化reids失败 %v", err)
		os.Exit(1)
	}
}
