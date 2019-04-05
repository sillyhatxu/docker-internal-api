package main

import (
	"github.com/gin-gonic/gin"
	"io/ioutil"
	"log"
	"net/http"
)

func main() {
	router := gin.Default()
	stockRouterGroup := router.Group("/inventories")
	{
		stockRouterGroup.GET("", queryInventories)
	}
	router.Run(":8080")
}

type Result struct {
	API string `json:"api"`

	InternalBody string `json:"internal_body"`
}

func queryInventories(context *gin.Context) {
	log.Println("api -> query inventories")
	resp, err := http.Get("http://internal-module-name:8080/internal-inventories")
	if err != nil {
		context.JSON(http.StatusOK, err.Error())
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		context.JSON(http.StatusOK, err.Error())
	}
	context.JSON(http.StatusOK, Result{API: "This is api result", InternalBody: string(body)})
}
