package main

import (
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"strconv"
)

func main() {
	router := gin.Default()
	stockRouterGroup := router.Group("/internal-inventories")
	{
		stockRouterGroup.GET("", queryInventories)
	}
	test := router.Group("/info")
	{
		test.GET("", info)
	}
	router.Run(":8080")
}

type Inventory struct {
	Id string `json:"id"`

	ProductId string `json:"product_id"`

	Quantity int `json:"quantity"`
}

func info(context *gin.Context) {
	context.JSON(http.StatusOK, "OK")
}

func queryInventories(context *gin.Context) {
	log.Println("internal api -> query inventories")
	var inventories []Inventory
	for i := 1; i <= 20; i++ {
		inventories = append(inventories, *&Inventory{Id: strconv.Itoa(i), ProductId: "P_" + strconv.Itoa(i), Quantity: i})
	}
	context.JSON(http.StatusOK, inventories)
}
