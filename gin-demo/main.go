package main

import (
	"errors"
	"net/http"
	"time"

	"github.com/Bowser1704/go-study/gin-demo/config"
	"github.com/Bowser1704/go-study/gin-demo/router"

	"github.com/gin-gonic/gin"
	"github.com/lexkong/log"
	"github.com/spf13/pflag"
	"github.com/spf13/viper"
)

var (
	//利用命令行参数 指定config的位置如果没有就是default
	cfg = pflag.StringP("config", "c", "", "gin-demo config file path")
)

func main() {
	pflag.Parse()

	//init config
	if err := config.Init(*cfg); err != nil {
		panic(err)
	}

	// Set gin mode
	gin.SetMode(viper.GetString("runmode"))
	g := gin.New()

	middlewares := []gin.HandlerFunc{}

	//routes
	router.Load(
		// *gin.Engine
		g,

		//middlewars
		middlewares..., )

	//ping the server to make sure router is working
	go func() {
		if err := pingServer(); err != nil {
			log.Fatal("the router has no response", err)
		}
		log.Info("the router has been deployed(部署) successfully")
	}()

	log.Infof("start to listening %s",viper.GetString("addr"))
	log.Info(http.ListenAndServe(viper.GetString("addr"), g).Error())
}

// ping
func pingServer() error {
	for i := 0; i < viper.GetInt("max_ping_count"); i++ {
		resp, err := http.Get(viper.GetString("url") + "/sd/health")
		if err == nil && resp.StatusCode == 200 {
			return nil
		}

		log.Info("waiting a second")
		time.Sleep(time.Second)
	}
	return errors.New("cannot connect")
}
