package main

import (
	"context"
	"fmt"
	"log"
	"transaction/config"
	"transaction/constants"
	"transaction/controllers"
	"transaction/routes"
	"transaction/services"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
)

var (
	mongoclient *mongo.Client
	ctx         context.Context
	server      *gin.Engine
)

func initRoutes() {
	routes.Default(server)
}

func initApp(mongoClient *mongo.Client) {
	ctx = context.TODO()
	//create collection
	profileCollection := mongoClient.Database(constants.DatabaseName).Collection("profiles")
	//pass the collection and context to initantiate the service
	profileService := services.InitTranasctionService(profileCollection, ctx)
	//pass the service to the controller
	profileController := controllers.InitTranasctionController(profileService)
	//pass the controller to the routes
	routes.TransactionRoute(server, profileController)
}

func main() {
	server = gin.Default()
	mongoclient, err := config.ConnectDataBase()
	defer mongoclient.Disconnect(ctx)
	if err != nil {
		panic(err)
	}
	initRoutes()
	initApp(mongoclient)
	fmt.Println("server running on port", constants.Port)
	log.Fatal(server.Run(constants.Port))
}
