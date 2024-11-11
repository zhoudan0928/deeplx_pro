package api

import (
	"deeplx-pro/config"
	"deeplx-pro/initialize"
	"deeplx-pro/translator"
	"net/http"

	"github.com/joho/godotenv"
)

func init() {
	// 加载环境变量
	godotenv.Load()

	// 初始化配置
	config.InitConfig()

	// 初始化翻译器
	translator.InitTranslator()
}

func Handler(w http.ResponseWriter, r *http.Request) {
	// 初始化Gin引擎和路由
	router := initialize.InitRouter()
	router.ServeHTTP(w, r)
}
