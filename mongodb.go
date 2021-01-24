package main

import (
	"context"
	"fmt"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

type student struct {
	Name string
	Roll int
}

func main() {
	client, err := mongo.NewClient(options.Client().ApplyURI("mongodb+srv://tarungorka09:Tarun%408886@cluster0.lhufx.mongodb.net/Example1?retryWrites=true&w=majority"))
	if err != nil {
		fmt.Print("Error in connecting")
	}
	ctx, _ := context.WithTimeout(context.Background(), 30*time.Second)
	err = client.Connect(ctx)
	if err != nil {
		log.Fatal(err)
	}
	defer client.Disconnect(ctx)
	err = client.Ping(ctx, readpref.Primary())
	if err != nil {
		log.Fatal(err)
		fmt.Print("error in pinging")
	}
	var t student
	databases, err := client.ListDatabaseNames(ctx, bson.M{})
	fmt.Println(databases)
	Example1Database := client.Database("Example1")
	InfoCollection := Example1Database.Collection("Info")
	err = InfoCollection.FindOne(ctx, bson.D{{"Name", "Tarun"}}).Decode(&t)
	fmt.Println(t)
}
