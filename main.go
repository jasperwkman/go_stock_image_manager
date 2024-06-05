package main

import (
	"fmt"
	"html/template"
	"os"
	"strings"

	"go_stock_image_manager/database"
	"go_stock_image_manager/handlers"

	"github.com/gin-gonic/gin"
	"github.com/go-ini/ini"
)

func main() {
	iniPath := "D:/projects/juliebox/db_config.ini"
	if len(os.Args) > 1 {
		iniPath = os.Args[1]
	}

	cfg, err := ini.Load(iniPath)
	if err != nil {
		panic(err)
	}
	server := cfg.Section("database").Key("server").String()
	port := cfg.Section("database").Key("port").String()
	user := cfg.Section("database").Key("username").String()
	password := cfg.Section("database").Key("password").String()
	dbname := cfg.Section("database").Key("database").String()

	db, err := database.Connect(server, port, user, password, dbname)
	if err != nil {
		panic(err)
	}
	defer db.Close()

	r := gin.Default()

	// define the custom split function
	//must defind before loading the templates by (r.LoadHTMLGlob("templates/*"))
	r.SetFuncMap(template.FuncMap{
		"split": strings.Split,
	})

	r.LoadHTMLGlob("templates/*")

	// Serve static files
	r.Static("/static", "./static")
	r.StaticFile("/favicon.ico", "./favicon.ico")

	r.GET("/", func(c *gin.Context) {
		handlers.IndexHandler(c, db)
	})

	r.GET("/imageDetail", func(c *gin.Context) {
		handlers.ImageDetailHandler(c, db)
	})

	r.POST("/api/setImageStatus", func(c *gin.Context) {
		handlers.SetImageStatusHandler(c, db)
	})

	r.POST("/api/setImageDescription", func(c *gin.Context) {
		handlers.SetImageDesctiptionHandler(c, db)
	})

	r.POST("/api/setImageTags", func(c *gin.Context) {
		handlers.SetImageTagsHandler(c, db)
	})

	r.GET("/api/setImageReady", func(c *gin.Context) {
		fmt.Println("/api/setImageReady")
		handlers.SetImageReadyHandler(c, db)
	})

	r.Run(":8080")
}
