package Database

import (
	"ToDo/Structures"
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
)

const dbName = "ToDo"
const collectionName = "rec"

var collection *mongo.Collection

//Functions with database
func init() {
	var clientOpts = options.Client().ApplyURI("mongodb://127.0.0.1:27017")
	var client, err = mongo.Connect(context.TODO(), clientOpts)
	if err != nil {
		log.Fatal(err)
	}
	if err = client.Ping(context.TODO(), nil); err != nil {
		log.Fatal(err)
	}
	collection = client.Database(dbName).Collection(collectionName)
}

func AddToDb(rec Structures.Records) {
	var _, err = collection.InsertOne(context.TODO(), rec)
	if err != nil {
		log.Fatal(err)
	}
}

func DelInDb(problem string) {
	var id, err = primitive.ObjectIDFromHex(problem)
	if err != nil {
		log.Fatal(err)
	}
	_, err = collection.DeleteOne(context.TODO(), bson.M{"id": id})
	if err != nil {
		log.Fatal(err)
	}
}
func DelAll() int64 {
	var c, err = collection.DeleteMany(context.TODO(), bson.M{}, nil)
	if err != nil {
		log.Fatal(err)
	}
	return c.DeletedCount
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
	if cur.Close(context.TODO()) != nil {
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
