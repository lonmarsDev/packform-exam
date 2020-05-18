package main

import (
	"net/http"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/lonmarsDev/packform-exam/internal/services"
	"github.com/lonmarsDev/packform-exam/pkg/config"
	"github.com/lonmarsDev/packform-exam/pkg/datastore"
)

func main() {

	config.Init()
	datastore.DbInit(config.AppConfig.GetString("database.mongo_db.database_url"), config.AppConfig.GetString("database.mongo_db.database_name"))
	err := datastore.DbPgInit(config.AppConfig.GetString("database.postgres.username"), config.AppConfig.GetString("database.postgres.password"), config.AppConfig.GetString("database.postgres.host"), config.AppConfig.GetString("database.postgres.db_name"), config.AppConfig.GetInt("database.postgres.port"))
	if err != nil {
		panic(err)
	}

	router := gin.Default()

	// Query string parameters are parsed using the existing underlying request object.
	router.Use(cors.Default())
	router.GET("/orders", func(c *gin.Context) {
		search := c.Query("query")
		page := c.Query("page")
		createdDateFrom := c.Query("date_from")
		createdDateTo := c.Query("date_to")
		allOrder := services.AllOrder(page, search, createdDateFrom, createdDateTo)
		c.JSON(http.StatusOK, allOrder)

	})
	router.Run(":8080")
}
