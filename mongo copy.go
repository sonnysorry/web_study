package main

import (
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
func handleAPIAuth(w http.ResponseWriter, r *http.Request){
	//ConnectDB 는 MongoDB 연결을 위한 Context, Client를 생성한 후 리턴한다.
	signkey := umsUser.SignKey
	appsUser := AppsUser{}
	appsUser, err = GetAppsUserByID(ctx, client, id)
	if err != nil {
		appsUser.ID = id
		appsUser.AccessLevel = "default"
		err = appsUser.AppsCreateToken(signkey)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
			}

		err = AddAppsUser(ctx, client, appsUser)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
			}
		}
	}

func ConnectDB() (context.Context, context.CancelFunc, *mongo.Client, error) {
	// Timeout 설정을 위한 Context생성
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)

	//mongoDB client 생성
	client, err := mongo.NewClient(options.Client().ApplyURI(*flagMongoDBURI))
	if err != nil {
		log.Print(err.Error())
		return ctx, cancel, nil, err
	}

	// db 연결 확인
	err = client.Connect(ctx)
	if err != nil {
		log.Print(err.Error())
		return ctx, cancel, client, err
	}
	err = client.Ping(ctx, readpref.Primary())
	if err != nil {
		log.Print(err.Error())
		return ctx, cancel, client, err
	}

	return ctx, cancel, client, nil
}

//---------------//
//----- Add -----//
//---------------//

// AddAppsUser 는 apps db에 user 정보를 저장하는 함수다.
func AddAppsUser(ctx context.Context, client *mongo.Client, user AppsUser) error {
	// findOne과 같지만 document는 반환하지 않는 옵션
	// mongoDB client 연결
	ctx, cancel, client, err := ConnectDB()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer cancel()
	defer client.Disconnect(ctx)
	
	options := options.Find()
	options.SetLimit(1)
	collection := client.Database("sample").Collection("users")
	cursor, err := collection.Find(ctx, bson.M{"id": user.ID}, options)
	if err != nil {
		return err
	}
	hasNext := cursor.TryNext(ctx)
	if hasNext {
		return errors.New("Apps에 존재하는 ID입니다")
	}

	_, err = collection.InsertOne(ctx, user)
	if err != nil {
		return err
	}
	return nil
}