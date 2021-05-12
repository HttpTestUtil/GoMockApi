package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"mock-api-go/models"
	"os"
	"regexp"
	"strings"

	"github.com/gin-gonic/gin"
)

const (
	CONFIG_PATH = "config/config.json"
	API_PATH    = "api"
)

func main() {
	r := gin.Default()

	initRoute(r)
	config, err := initConfig()
	if err != nil {
		return
	}
	r.Run(fmt.Sprintf(":%d", config.ListenPort))
	fmt.Printf("server listen on %d", config.ListenPort)
}

/**
 * initConfig 初始化配置项
 */
func initConfig() (*models.Config, error) {
	// 读取配置项
	config := &models.Config{}
	json_str, err := ioutil.ReadFile(CONFIG_PATH)
	if err != nil {
		fmt.Println(err.Error())
		return config, err
	}
	json.Unmarshal(json_str, config)
	return config, nil
}

/**
 * initRoute 初始化配置的类
 */
func initRoute(r *gin.Engine) error {
	dirs, err := os.ReadDir(API_PATH)
	if err != nil {
		fmt.Println(err.Error())
		return err
	}
	rex, _ := regexp.Compile(".json$")
	for _, dir := range dirs {
		// 如果文件名不对则忽略
		if !rex.Match([]byte(dir.Name())) {
			continue
		}
		tmpRouteSlice := []models.Router{}
		filePath := fmt.Sprintf("%s/%s", API_PATH, dir.Name())
		json_str, err := ioutil.ReadFile(filePath)
		if err != nil {
			continue
		}
		json.Unmarshal(json_str, &tmpRouteSlice)
		fmt.Print(tmpRouteSlice)
		for _, routerData := range tmpRouteSlice {

			switch strings.ToLower(routerData.Method) {
			case models.METHOD_GET:
				r.GET(routerData.Route, func(c *gin.Context) {
					c.JSON(200, routerData.Response)
				})
			case models.METHOD_POST:
				r.POST(routerData.Route, func(c *gin.Context) {
					c.JSON(200, routerData.Response)
				})
			}

		}
	}
	return nil
}
