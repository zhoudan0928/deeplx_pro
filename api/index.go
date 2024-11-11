package api

import (
	"net/http"
	"sync"

	"deeplx-pro/config"
	"deeplx-pro/initialize"
	"deeplx-pro/translator"
)

var initOnce sync.Once

func Handler(w http.ResponseWriter, r *http.Request) {
	// 初始化配置和翻译器，仅在首次请求时执行
	initOnce.Do(func() {
		// 初始化配置
		config.InitConfig()
		// 初始化翻译器
		translator.InitTranslator()
	})

	// 初始化Gin引擎和路由
	router := initialize.InitRouter()
	router.ServeHTTP(w, r)
}
