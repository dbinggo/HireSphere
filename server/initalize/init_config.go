package initalize

import (
	"flag"
	"github.com/Dbinggo/HireSphere/server/configs"
	"github.com/Dbinggo/HireSphere/server/global"
	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
	"time"
)

func InitConfig() {
	// 初始化时间为东八区的时间
	var cstZone = time.FixedZone("CST", 8*3600) // 东八
	time.Local = cstZone

	// 默认配置文件路径
	var configPath string
	flag.StringVar(&configPath, "c", global.Path+global.DEFAULT_CONFIG_FILE_PATH, "配置文件绝对路径或相对路径")
	flag.Parse()
	global.Logger.Infof("配置文件路径为 %s", configPath)
	// 初始化配置文件
	viper.SetConfigFile(configPath)
	viper.WatchConfig()
	// 观察配置文件变动
	viper.OnConfigChange(func(in fsnotify.Event) {
		global.Logger.Warnf("配置文件发生变化")
		if err := viper.Unmarshal(&configs.Conf); err != nil {
			global.Logger.Errorf("无法反序列化配置文件 %v", err)
		}
		global.Logger.Debugf("%+v", configs.Conf)
		global.Config = configs.Conf
	})
	// 将配置文件读入 viper
	if err := viper.ReadInConfig(); err != nil {
		global.Logger.Panicf("无法读取配置文件 err: %v", err)

	}
	// 解析到变量中
	if err := viper.Unmarshal(&configs.Conf); err != nil {
		global.Logger.Panicf("无法解析配置文件 err: %v", err)
	}
	global.Logger.Debugf("配置文件为 ： %+v", configs.Conf)
	global.Config = configs.Conf
}
