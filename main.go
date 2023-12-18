package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
)

var router *gin.Engine

func main() { 
	e := connect()
	if e != nil {
		fmt.Println(e.Error())
		return
	}
	router = gin.Default()

	router.Static("/assets", cfg.Assets)

	router.LoadHTMLFiles(cfg.HTML + "index.html")

	router.GET("/", index)

	router.Run(cfg.ServerHost + ":" + cfg.ServerPort)
}

func index (c *gin.Context) { 
	var cat category
	e := cat.Select()
	if e != nil {
		fmt.Println(e.Error())
	}
	c.HTML(200, "index.html", gin.H{"Category": cat.Rows})
}