package main

import (
	"BCFrontend/models"
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	"io/ioutil"
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

func getUser(ctx *gin.Context) models.User {
	cookie := ctx.Request.Header.Get("Cookie")
	request, err := http.NewRequest("GET", viper.GetString("auth_service_addr")+"/sessions/me", nil)
	fmt.Println(cookie)
	if err != nil {
		fmt.Println(err.Error())
		return models.User{}
	}
	request.Header.Add("Cookie", cookie)
	client := http.Client{}
	response, err := client.Do(request)
	if err != nil {
		fmt.Println(err.Error())
		return models.User{}
	}
	body, err := ioutil.ReadAll(response.Body)
	fmt.Println(string(body))
	var output models.User
	err = json.Unmarshal(body, &output)
	fmt.Println(output)
	if err != nil {
		return models.User{}
	}
	return output
}

func run() {
	engine = gin.Default()
	engine.Static("/static", "../frontend")
	engine.LoadHTMLGlob("../frontend/*/*")

	engine.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.tmpl", gin.H{
			"auth_service": viper.GetString("auth_service"),
			"user":         getUser(c),
		})
	})
	engine.GET("/sign-up", func(c *gin.Context) {
		c.HTML(http.StatusOK, "signup.tmpl", gin.H{
			"auth_service": viper.GetString("auth_service"),
		})
	})
	engine.GET("/sign-in", func(c *gin.Context) {
		c.HTML(http.StatusOK, "signin.tmpl", gin.H{
			"auth_service": viper.GetString("auth_service"),
		})
	})

	err := engine.Run(viper.GetString("host") + ":" + viper.GetString("port"))
	if err != nil {
		panic(err.Error())
	}
}
