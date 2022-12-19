package main

import (
	"context"
	"fmt"

	"github.com/labstack/echo/v4"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func NewMongoConnection() *mongo.Client {
	// connect to mongo
	client, err := mongo.NewClient(options.Client().ApplyURI("mongodb://localhost:27017"))
	if err != nil {
		fmt.Println("Error creating mongo client")
		panic(err)
	}
	err = client.Connect(context.Background())
	if err != nil {
		fmt.Println("Error connecting to mongo")
		panic(err)
	}
	return client
}

func main() {
	e := echo.New()
	// connect to mongo
	db := NewMongoConnection()
	defer db.Disconnect(context.Background())
	// ping mongo to check connection
	err := db.Ping(context.Background(), nil)
	if err != nil {
		fmt.Println("Error pinging mongo")
		panic(err)
	}
	todoCollection := db.Database("todo").Collection("todo")

	todoRepository := NewTodoItemMongoRepository(todoCollection)
	todoHandler := NewTodoItemHandler(todoRepository)

	// configure routing
	ConfigureRouting(e, todoHandler)
	// start server
	e.Logger.Fatal(e.Start(":8080"))

	// graceful shutdown
	// ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	// defer cancel()
	// if err := e.Shutdown(ctx); err != nil {
	// 	e.Logger.Fatal(err)
	// }
}
