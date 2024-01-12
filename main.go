package main

import (
	"fmt"
	"media-sync-tools/db"
	_ "media-sync-tools/docs"
	"media-sync-tools/routes"
	"net/http"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello, %q", r.URL.Path)
}

func main() {
	// http.HandleFunc("/", handler)
	// http.ListenAndServe(":8080", nil)

	router := gin.Default()

	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	db := db.SetupDatabase()
	routes.UserRoutes(router, db)
	router.Run(":8080")
}
