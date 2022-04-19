package main

import (
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	"net/http"
)

var engine *gin.Engine

func main() {
	err := parseConfig()
	if err != nil {
		panic(err)
	}

	run()
}

func parseConfig() error {
	configPath := "../configuration"

	viper.AddConfigPath(configPath)
	viper.SetConfigName("config")
	return viper.ReadInConfig()
}

func run() {
	engine = gin.Default()
	engine.Static("/static", "../frontend")
	engine.LoadHTMLGlob("../frontend/*/*")

	engine.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.tmpl", gin.H{
			"auth_service": viper.GetString("auth_service"),
		})
	})
	engine.GET("/sign-up", func(c *gin.Context) {
		c.HTML(http.StatusOK, "signup.tmpl", gin.H{
			"auth_service": viper.GetString("auth_service"),
		})
	})

	err := engine.Run(viper.GetString("host") + ":" + viper.GetString("port"))
	if err != nil {
		panic(err.Error())
	}
}
