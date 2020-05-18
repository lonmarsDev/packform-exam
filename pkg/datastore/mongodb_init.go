package datastore

import (
	"context"
	"sync"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var (
	DbCollections = Collections{}
	mutex         = &sync.Mutex{}
)

type Collections struct {
	CustomerCompanies *mongo.Collection
	Customers         *mongo.Collection
	Orders            *mongo.Collection
}

func DbInit(dbUrl, dbName string) {
	mutex.Lock()
	defer mutex.Unlock()
	database := CreateDbConnection(dbUrl, dbName)
	DbCollections = Collections{
		CustomerCompanies: database.Collection("customerCompanies"),
		Customers:         database.Collection("customers"),
		Orders:            database.Collection("orders"),
	}
}

func CreateDbConnection(url string, database string) *mongo.Database {
	client, err := mongo.NewClient(options.Client().ApplyURI(url))
	if err != nil {
		panic(err)
	}
	ctx, _ := context.WithTimeout(context.Background(), 5*time.Second)
	err = client.Connect(ctx)
	if err != nil {
		panic(err)
	}
	return client.Database(database)
}
