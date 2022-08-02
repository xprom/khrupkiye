package main

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Tovar struct {
	Title string `bson:"title"`
	Text  string `bson:"text"`
	Price int    `bson:"price"`
}

/**
Фоново нужно запустить базу данных
> docker run -ti --network=host -p 27017:27017 -d mongo:latest
> mongo --host=127.0.0.1
*/

func main() {
	ctx := context.Background()

	mongoDB, err := mongo.NewClient(
		options.Client().ApplyURI("mongodb://192.168.64.174:27017"),
	)
	err = mongoDB.Connect(ctx)
	if err != nil {
		fmt.Println(err)
		return
	}

	cur, err := mongoDB.Database("marketplace").Collection("catalog").Find(ctx, bson.D{}, nil)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer cur.Close(ctx)

	for cur.Next(ctx) {
		var tovar Tovar
		err := cur.Decode(&tovar)
		if err != nil {
			fmt.Println(err)
		}

		fmt.Println("Информация о товаре")
		fmt.Println(fmt.Sprintf("Название %s", tovar.Title))
		fmt.Println(fmt.Sprintf("Цена %d", tovar.Price))
	}
}
