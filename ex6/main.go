package main

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type TestRow struct {
	Title string  `bson:"title"`
	Price float64 `bson:"price"`
}

func main() {
	ctx := context.Background()

	mongoDB, err := mongo.NewClient(
		options.Client().ApplyURI("mongodb://192.168.64.165:27017"),
	)
	err = mongoDB.Connect(ctx)
	if err != nil {
		fmt.Println(err)
		return
	}

	findOption := options.Find()
	findOption.SetSort(bson.D{{"_id", 1}})
	cur, err := mongoDB.Database("marketplace").Collection("goods").Find(ctx, bson.D{}, nil)
	if err != nil {
		fmt.Println(err)
		return
	}

	defer cur.Close(ctx)
	for cur.Next(ctx) {
		var row TestRow
		err := cur.Decode(&row)
		if err != nil {
			fmt.Println(err)
		}

		fmt.Println("Информация о товаре")
		fmt.Println(fmt.Sprintf("Название %s", row.Title))
		fmt.Println(fmt.Sprintf("Цена %f", row.Price))
	}
}
