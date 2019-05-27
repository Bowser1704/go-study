package config

import (
	_ "github.com/golang/net/dns/dnsmessage"
	"strings"

	"github.com/fsnotify/fsnotify"
	"github.com/lexkong/log"
	"github.com/spf13/viper"
)

type Config struct {
	Name	string	//配置文件名字
}

func Init(cfg string) error {
	c := Config{
		Name:cfg,
	}

	//初始化配置文件
	if err := c.initConfig(); err != nil {
		return err
	}

	//初始化日志包
	c.initLog()

	//监控配置文件变化并热加载程序
	c.watchConfig()

	return nil
}

//定义上面用到的方法
func (c *Config) initConfig() error {
	if c.Name != "" {
		viper.SetConfigFile(c.Name)	//如果指定文件则使用指定的配置文件
	}else {
		//viper.AddConfigPath("~/go-work/src/github.com/Bowser1704/go-study/gin-demo/conf")	//添加默认的配置文件环境
		viper.SetConfigFile("conf/config.yaml")	//默认配置文件名字
	}

	//viper.SetConfigType("yaml")	//默认文件配置格式
	viper.AutomaticEnv()	//读取匹配的环境变量
	viper.SetEnvPrefix("APISERVER")	//读取环境变量的前缀为APISERVER
	replacer := strings.NewReplacer(".","_")	//替换.为_
	viper.SetEnvKeyReplacer(replacer)
	if err := viper.ReadInConfig(); err != nil {	//viper解析配置文件
		return err
	}

	return nil
}

func (c *Config) initLog() {
	passLagerCfg := log.PassLagerCfg{

		Writers:        viper.GetString("log.writers"),
		LoggerLevel:    viper.GetString("log.logger_level"),
		LoggerFile:     viper.GetString("log.logger_file"),
		LogFormatText:  viper.GetBool("log.log_format_text"),
		RollingPolicy:  viper.GetString("log.rollingPolicy"),
		LogRotateDate:  viper.GetInt("log.log_rotate_date"),
		LogRotateSize:  viper.GetInt("log.log_rotate_size"),
		LogBackupCount: viper.GetInt("log.log_backup_count"),
	}

	log.InitWithConfig(&passLagerCfg)
}

//监控配置变化并且热加载程序
func (c *Config) watchConfig() {
	viper.WatchConfig()
	viper.OnConfigChange(func(e fsnotify.Event) {
		log.Infof("Config file changed: %s", e.Name)
	})
}
