package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"goblog/app/http/middlewares"
	"goblog/bootstrap"
	"goblog/pkg/database"
	"net/http"
)

var router *mux.Router

func main() {
	database.Initialize()

	bootstrap.SetupDB()
	router = bootstrap.SetupRoute()

	// 通过命名路由获取 URL 示例
	homeURL, _ := router.Get("home").URL()
	fmt.Println("homeURL: ", homeURL)
	articleURL, _ := router.Get("articles.show").URL("id", "23")
	fmt.Println("articleURL: ", articleURL)

	http.ListenAndServe(":3000", middlewares.RemoveTrailingSlash(router))
}
