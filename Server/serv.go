package Server

import (
	"ToDo/Structures"
	"context"
	"encoding/json"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"net/http"
	"os"
)

const HOST = "mongodb://localhost:27017"
const dbName = "ToDo"
const collectionName = "rec"

var collection *mongo.Collection

func init_db() {
	var clientOpts = options.Client().ApplyURI(HOST)
	var client, err = mongo.Connect(context.TODO(), clientOpts)
	if err != nil {
		fmt.Println("Connect error: ", err)
		os.Exit(1)
	}
	err = client.Ping(context.TODO(), nil)
	if err != nil {
		fmt.Println("Ping error: ", err)
		os.Exit(2)
	}
	collection = client.Database(dbName).Collection(collectionName)
}

func AddToDb(rec Structures.Records) {
	var _, err = collection.InsertOne(context.TODO(), rec)
	if err != nil {
		fmt.Println("Add record error: ", err)
		os.Exit(3)
	}
}

func DelInDb(problem string) {
	var id, err = primitive.ObjectIDFromHex(problem)
	if err != nil {
		fmt.Println("Create ObjectID error: ", err)
		os.Exit(4)
	}
	_, err = collection.DeleteOne(context.TODO(), bson.M{"id": id})
	if err != nil {
		fmt.Println("Delete record error: ", err)
		os.Exit(5)
	}
}

func updateRec(problem string, stat int) {
	var id, err = primitive.ObjectIDFromHex(problem)
	if err != nil {
		fmt.Println("Create ObjectID error: ", err)
		os.Exit(6)
	}
	if stat == 1 {
		_, err = collection.UpdateOne(context.TODO(), bson.M{"id": id}, bson.M{"$set": bson.M{"status": true}})
	} else {
		_, err = collection.UpdateOne(context.TODO(), bson.M{"id": id}, bson.M{"$set": bson.M{"status": false}})
	}
	if err != nil {
		fmt.Println("Update record error: ", err)
		os.Exit(7)
	}
}

func CreateProblem(w http.ResponseWriter, r *http.Request) {
	var jsonDecoder = json.NewDecoder(r.Body)
	var rec Structures.Records
	err := jsonDecoder.Decode(&rec)
	if err != nil {
		fmt.Println("Decode error", err)
	}
}
