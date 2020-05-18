package main

import (
	"github.com/lonmarsDev/packform-exam/internal/services"
	"github.com/lonmarsDev/packform-exam/pkg/config"
	"github.com/lonmarsDev/packform-exam/pkg/datastore"
)

func main() {
	config.Init()
	datastore.DbInit(config.AppConfig.GetString("database.mongo_db.database_url"), config.AppConfig.GetString("database.mongo_db.database_name"))
	services.ImportCustomer()
	services.ImportCustomerCompany()
	//services.ImportOrder()
	err := datastore.DbPgInit(config.AppConfig.GetString("database.postgres.username"), config.AppConfig.GetString("database.postgres.password"), config.AppConfig.GetString("database.postgres.host"), config.AppConfig.GetString("database.postgres.db_name"), config.AppConfig.GetInt("database.postgres.port"))
	if err != nil {
		panic(err)
	}
	//Postgres schema migration
	datastore.Migrate()
	//Import Order
	services.ImportOrder()
	//Import Order Item
	services.ImportOrderItem()
	//Import Delivery
	services.ImportDelivery()

}
