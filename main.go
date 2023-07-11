package main

import (
	"context"
	"encoding/json"
	"log"
	"math/rand"
	"net/http"

	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type User struct {
	ID       string `json:"id,omitempty"`
	Username string `json:"username,omitempty"`
	Email    string `json:"email,omitempty"`
}

var client *mongo.Client
var collection *mongo.Collection

func main() {
	// Connect to MongoDB
	clientOptions := options.Client().ApplyURI("mongodb://mongoadmin:secret@localhost:27017")
	// for private docker repo
	//clientOptions := options.Client().ApplyURI("mongodb://mongoadmin:secret@mongodb")
	client, err := mongo.Connect(context.Background(), clientOptions)
	if err != nil {
		log.Fatal(err)
	}

	// Get a handle to the user-account collection
	collection = client.Database("user-account").Collection("user-details")

	// Create a new HTTP router
	router := mux.NewRouter()

	// Serve static files
	fs := http.FileServer(http.Dir("./static"))
	router.PathPrefix("/static/").Handler(http.StripPrefix("/static/", fs))

	// Define the API routes
	router.HandleFunc("/users", createUser).Methods("POST")
	router.HandleFunc("/users/{id}", getUser).Methods("GET")
	router.HandleFunc("/users", getUsers).Methods("GET")

	// Serve the HTML file
	router.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "static/index.html")
	}).Methods("GET")

	// Start the HTTP server
	log.Println("Server started on port 8080")
	log.Fatal(http.ListenAndServe(":8080", router))
}
func generateID() string {
	// Generate a unique ID using any suitable method (e.g., uuid or shortid)
	// For example, using the github.com/google/uuid package:
	// id := uuid.New().String()

	// Placeholder function to generate a random string for simplicity
	const idLength = 8
	const idCharacters = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	id := make([]byte, idLength)
	for i := 0; i < idLength; i++ {
		id[i] = idCharacters[rand.Intn(len(idCharacters))]
	}
	return string(id)
}
func createUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var user User
	json.NewDecoder(r.Body).Decode(&user)
	// Generate a unique ID
	user.ID = generateID()
	result, err := collection.InsertOne(context.Background(), user)
	if err != nil {
		log.Fatal(err)
	}
	json.NewEncoder(w).Encode(result.InsertedID)
}
func getUsers(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var users []User

	// Retrieve all users from the MongoDB collection
	cur, err := collection.Find(context.Background(), bson.D{})
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(map[string]string{"error": "Failed to fetch user details"})
		return
	}
	defer cur.Close(context.Background())

	// Iterate over the cursor and append each user to the users slice
	for cur.Next(context.Background()) {
		var user User
		if err := cur.Decode(&user); err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			json.NewEncoder(w).Encode(map[string]string{"error": "Failed to fetch user details"})
			return
		}
		users = append(users, user)
	}
	if err := cur.Err(); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(map[string]string{"error": "Failed to fetch user details"})
		return
	}

	json.NewEncoder(w).Encode(users)
}

func getUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	var user User
	err := collection.FindOne(context.Background(), bson.M{"id": params["id"]}).Decode(&user)
	if err != nil {
		log.Fatal(err)
	}
	json.NewEncoder(w).Encode(user)
}
