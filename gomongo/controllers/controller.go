package controllers

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/geekysaurabh001/gomongo/models"
	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const connectionString string = "mongodb+srv://saurabhsri:zcninZnIEG7irqw5@cluster0.z3munre.mongodb.net/netflix?retryWrites=true&w=majority"
const dbName string = "netflix"
const colName string = "watchlist"

var collection *mongo.Collection // take the reference of the mongodb connection

// connect with mongodb with special golang method init which runs only once at the beginning of the app start
func init() {
	// client options
	clientOption := options.Client().ApplyURI(connectionString)

	// connect with mongodb
	client, err := mongo.Connect(context.TODO(), clientOption)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("MongoDB connection works")

	collection = client.Database(dbName).Collection(colName)
	fmt.Println("Collection instance is ready", collection)
}

// mongodb helpers

// insert 1 record
func insertOneMovie(movie models.Netflix) {
	inserted, err := collection.InsertOne(context.Background(), movie)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Inserted 1 movie in db with id:", inserted.InsertedID)
}

// update 1 record
func updateOneMovie(movieId string) {
	id, err := primitive.ObjectIDFromHex(movieId)
	if err != nil {
		log.Fatal(err)
	}
	filter := bson.M{"_id": id}
	update := bson.M{"$set": bson.M{"watched": true}}
	result, err := collection.UpdateOne(context.Background(), filter, update)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Updated 1 movie in db with id:", result.ModifiedCount)
}

// delete 1 record
func deleteOneMovie(movieId string) {
	id, err := primitive.ObjectIDFromHex(movieId)
	if err != nil {
		log.Fatal(err)
	}
	filter := bson.M{"_id": id}
	deleteCount, err := collection.DeleteOne(context.Background(), filter, nil)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Movie got deleted with delete count:", deleteCount)
}

// delete all records
func deleteAllMovies() int64 {
	deleteResult, err := collection.DeleteMany(context.Background(), bson.D{{}}, nil)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Number of movies deleted:", deleteResult.DeletedCount)
	return deleteResult.DeletedCount
}

// get all movies from database
func getAllMovies() []primitive.M {
	cursor, err := collection.Find(context.Background(), bson.D{{}})
	if err != nil {
		log.Fatal(err)
	}
	defer cursor.Close(context.Background())
	var movies []primitive.M
	for cursor.Next(context.Background()) {
		var movie bson.M
		err := cursor.Decode(&movie)
		if err != nil {
			log.Fatal(err)
		}
		movies = append(movies, movie)
	}
	return movies
}

// Controllers are below

func GetAllMovies(res http.ResponseWriter, req *http.Request) {
	res.Header().Set("Content-Type", "application/x-ww-form-urlencode")
	res.Header().Set("Allow-Control-Allow-Methods", "GET")
	allMovies := getAllMovies()
	json.NewEncoder(res).Encode(allMovies)
}

func InsertOneMovie(res http.ResponseWriter, req *http.Request) {
	res.Header().Set("Content-Type", "application/x-ww-form-urlencode")
	res.Header().Set("Allow-Control-Allow-Methods", "POST")

	var movie models.Netflix
	_ = json.NewDecoder(req.Body).Decode(&movie)
	insertOneMovie(movie)
	json.NewEncoder(res).Encode(movie)
}

func UpdateOneMovie(res http.ResponseWriter, req *http.Request) {
	res.Header().Set("Content-Type", "application/x-ww-form-urlencode")
	res.Header().Set("Allow-Control-Allow-Methods", "PUT")

	params := mux.Vars(req)
	updateOneMovie(params["id"])
	json.NewEncoder(res).Encode(params["id"])
}

func DeleteOneMovie(res http.ResponseWriter, req *http.Request) {
	res.Header().Set("Content-Type", "application/x-ww-form-urlencode")
	res.Header().Set("Allow-Control-Allow-Methods", "DELETE")
	params := mux.Vars(req)
	deleteOneMovie(params["id"])
	json.NewEncoder(res).Encode(params["id"])
}

func DeleteAllMovie(res http.ResponseWriter, req *http.Request) {
	res.Header().Set("Content-Type", "application/x-ww-form-urlencode")
	res.Header().Set("Allow-Control-Allow-Methods", "DELETE")
	count := deleteAllMovies()
	json.NewEncoder(res).Encode(count)
}
