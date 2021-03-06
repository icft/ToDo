package Database

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
)

const HOST = "mongodb://127.0.0.1:27017"
const dbName = "todo"
const collectionName = "notes"

var collection *mongo.Collection

func init() {
	var clientOpts = options.Client().ApplyURI(HOST)
	var client, err = mongo.Connect(context.TODO(), clientOpts)
	if err != nil {
		log.Fatal(err)
	}
	if err = client.Ping(context.TODO(), nil); err != nil {
		log.Fatal(err)
	}
	fmt.Println("******************")
	fmt.Println("Connection complete")
	collection = client.Database(dbName).Collection(collectionName)
	fmt.Println("Get collection complete")
	fmt.Printf("******************\n\n")
}

func AddToDb(rec Note) {
	if _, err := collection.InsertOne(context.TODO(), rec); err != nil {
		log.Fatal(err)
	}
}

func DelInDb(problem string) {
	var id, err = primitive.ObjectIDFromHex(problem)
	if err != nil {
		log.Fatal(err)
	}
	if _, err = collection.DeleteOne(context.TODO(), bson.M{"id": id}); err != nil {
		log.Fatal(err)
	}
}
func DelAll() {
	if _, err := collection.DeleteMany(context.TODO(), bson.M{}, nil); err != nil {
		log.Fatal(err)
	}
}

func GetAll() []primitive.M {
	var cur, err = collection.Find(context.TODO(), bson.D{})
	if err != nil {
		log.Fatal(err)
	}
	var results []primitive.M
	for cur.Next(context.TODO()) {
		var result bson.M
		if err := cur.Decode(&result); err != nil {
			log.Fatal(err)
		}
		results = append(results, result)
	}
	if err = cur.Err(); err != nil {
		log.Fatal(err)
	}
	if err = cur.Close(context.TODO()); err != nil {
		log.Fatal(err)
	}
	return results
}

func UpdateRec(problem string, stat int) {
	var id, err = primitive.ObjectIDFromHex(problem)
	if err != nil {
		log.Fatal(err)
	}
	if stat == 1 {
		_, err = collection.UpdateOne(context.TODO(), bson.M{"id": id}, bson.M{"$set": bson.M{"status": true}})
	} else {
		_, err = collection.UpdateOne(context.TODO(), bson.M{"id": id}, bson.M{"$set": bson.M{"status": false}})
	}
	if err != nil {
		log.Fatal(err)
	}
}
