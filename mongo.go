package main

import (
	"fmt"
	"net/http"
	"context"
	"errors"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

func connection() {
	clientOptions := options.Client().ApplyURI("10.20.31.178:27017")
	client, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("연결 성공")
	usersCollection	:= client.Database("sample").Collection("users")
	fmt.Println(usersCollection)

	cursor, err := usersCollection.Find(context.TODO(), bson.D{{}})
	if err != nil {
		fmt.Println("에러")
		fmt.Println(err)
	}
	for cursor.Next(context.TODO()){
		var elem bson.M
		err := cursor.Decode(&elem)
		if err != nill {
			fmt.Println(elem)
		}
	}

	insertResult, _ := usersCollection.InsertOne(context.TODO(), bson.D{
		
	})
	err = client.Disconnect(context.TODO())
	if err != nill {
		log.Fatal(err)
	}

	fmt.Println("연결 종료")
}