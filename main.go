package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"mock-api-go/models"
	"regexp"
	"strings"

	"github.com/gin-gonic/gin"
)

// const (
// 	CONFIG_PATH = "config/config.json"
// 	API_PATH    = "api"
// )

var port string
var apiPath string

func init() {
	flag.StringVar(&port, "port", "8080", "启动服务的接口")
	flag.StringVar(&apiPath, "api-path", "api", "api文件所在的路径")

}
func main() {
	flag.Parse()
	r := gin.Default()

	initRoute(r)
	// config, err := initConfig()
	// if err != nil {
	// 	return
	// }
	r.Run(fmt.Sprintf(":%s", port))
}

/**
 * initRoute 初始化配置的类
 */
func initRoute(r *gin.Engine) error {
	// dirs, err := ioutil.ReadDir(apiPath)
	// if err != nil {
	// 	fmt.Println(err.Error())
	// 	return err
	// }
	allJsonPath, err := getAlljsonFilePath(apiPath)
	if err != nil {
		fmt.Println(err.Error())
		return err
	}

	for _, filePath := range allJsonPath {
		// // 如果文件名不对则忽略
		// if !rex.Match([]byte(dir.Name())) {
		// 	continue
		// }
		tmpRouteSlice := []models.Router{}
		// filePath := fmt.Sprintf("%s/%s", apiPath, dir.Name())
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

/**
 * 递归遍历整个json目录
 */
func getAlljsonFilePath(path string) ([]string, error) {
	dirs, err := ioutil.ReadDir(path)
	if err != nil {
		return nil, err
	}
	responseStrings := []string{}
	rex, _ := regexp.Compile(".json$")
	for _, oneDir := range dirs {
		// 如果是 json 文件，则拼接文件路径
		currentPath := fmt.Sprintf("%s/%s", path, oneDir.Name())
		if rex.Match([]byte(oneDir.Name())) {
			responseStrings = append(responseStrings, currentPath)
			continue
		}
		// 如果是文件夹，则递归寻找相关文件
		if oneDir.IsDir() {
			var tmpResponse []string
			tmpResponse, err = getAlljsonFilePath(currentPath)
			fmt.Println(tmpResponse)
			if err != nil {
				return nil, err
			}
			responseStrings = append(responseStrings, tmpResponse...)
			continue
		}
	}
	return responseStrings, nil
}
