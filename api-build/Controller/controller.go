package controller

import (
	model "api-basic/models"
	"context"
	"database/sql"
	"encoding/json"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"log"
	"net/http"
	"time"
)

const connectionString = "mongodb+srv://<username>:<password>@cluster0.jmgkygr.mongodb.net/?retryWrites=true&w=majority"
const dbName = "netflix"
const colName = "watch-list"

var collection *mongo.Collection

var database *sql.DB

func init() {
	var err error
	database, err = sql.Open("mysql", "root:bkashWEB786@tcp(0.0.0.0:3306)/bkash_toll")
	if err != nil {
		panic(err)
	}
	err = database.Ping()

	// handle error
	if err != nil {
		panic(err)
	}

	fmt.Print("Pong\n")
}

func insertOneMovie(movie model.Netflix) {
	created, err := collection.InsertOne(context.Background(), movie)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Inserted movie", created.InsertedID)
}

func updateOneMovie(movieId string) {
	id, _ := primitive.ObjectIDFromHex(movieId)
	filter := bson.M{"_id": id}
	update := bson.M{"$set": bson.M{"watched": true}}

	result, err := collection.UpdateOne(context.Background(), filter, update)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Modified count: ", result.ModifiedCount)
}

func deleteOneMovie(movieId string) {
	id, _ := primitive.ObjectIDFromHex(movieId)
	filter := bson.M{"_id": id}

	result, err := collection.DeleteOne(context.Background(), filter)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("DeleteCount is: ", result.DeletedCount)
}

func deleteAllMovie() {
	result, err := collection.DeleteMany(context.Background(), bson.D{{}})

	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("DeleteCount is: ", result.DeletedCount)
}

func getAllMovie() []primitive.M {
	cursor, err := collection.Find(context.Background(), bson.D{})

	if err != nil {
		log.Fatal(err)
	}

	var movies []primitive.M

	for cursor.Next(context.Background()) {
		var movie bson.M

		err := cursor.Decode(&movie)
		if err != nil {
			log.Fatal(err)
		}

		movies = append(movies, movie)
	}

	defer cursor.Close(context.Background())

	return movies
}

func GetAllMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	allMovies := getAllMovie()
	json.NewEncoder(w).Encode(allMovies)
}

func CreateMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var movie model.Netflix

	_ = json.NewDecoder(r.Body).Decode(&movie)

	insertOneMovie(movie)
	json.NewEncoder(w).Encode("Create success")
}

func MarkAsWatched(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Methods", "PUT")

	params := mux.Vars(r)

	updateOneMovie(params["id"])
	json.NewEncoder(w).Encode(params["id"])
}

func DeleteMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Methods", "DELETE")

	params := mux.Vars(r)

	deleteOneMovie(params["id"])
	json.NewEncoder(w).Encode("Delete the movie")
}

func DeleteAllMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Methods", "DELETE")

	deleteAllMovie()
	json.NewEncoder(w).Encode("Delete all movies success")
}

func HealthCheck(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	fmt.Println("Health Check")
	json.NewEncoder(w).Encode("Health is okk")
}

func GetAllClasses(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Methods", "GET")

	if database == nil {
		http.Error(w, "Database connection is nil", http.StatusInternalServerError)
		return
	}

	rows, err := database.Query("SELECT * FROM vehicle_classes")
	if err != nil {
		panic(err)
	}
	defer rows.Close()

	// Define a struct to hold the result
	type VehicleClass struct {
		ID        int       `json:"id"`
		Name      string    `json:"name"`
		CreatedAt time.Time `json:"created_at"`
		UpdatedAt time.Time `json:"updated_at"`
		// Add more fields as needed
	}

	var vehicleClasses []VehicleClass

	// Iterate through the result set and scan rows into the struct
	for rows.Next() {
		var vehicleClass VehicleClass
		var createdAt, updatedAt []uint8

		if err := rows.Scan(&vehicleClass.ID, &vehicleClass.Name, &createdAt, &updatedAt); err != nil {
			panic(err)
		}
		vehicleClasses = append(vehicleClasses, vehicleClass)
	}

	// Encode the result as JSON and send it as a response
	if err := json.NewEncoder(w).Encode(vehicleClasses); err != nil {
		panic(err)
	}
}
