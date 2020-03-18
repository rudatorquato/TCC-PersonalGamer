package main

import (
	//. "PersonalGamer/models"
	"PersonalGamer/helper"
	"PersonalGamer/models"

	"context"
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func getUsers(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var users []models.Users

	collection := helper.ConnectDB()

	cur, err := collection.Find(context.TODO(), bson.M{})

	if err != nil {
		helper.GetError(err, w)
		return
	}

	defer cur.Close(context.TODO())

	for cur.Next(context.TODO()) {

		var user models.Users

		err := cur.Decode(&user)
		if err != nil {
			log.Fatal(err)
		}

		users = append(users, user)
	}

	if err := cur.Err(); err != nil {
		log.Fatal(err)
	}

	json.NewEncoder(w).Encode(users)
}

func getUserName(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")

	var users models.Users

	var params = mux.Vars(r)

	id, _ := primitive.ObjectIDFromHex(params["id"])

	collection := helper.ConnectDB()

	filter := bson.M{"_id": id}
	err := collection.FindOne(context.TODO(), filter).Decode(&users)

	if err != nil {
		helper.GetError(err, w)
		return
	}

	json.NewEncoder(w).Encode(users)
}

func createUsers(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var users models.Users

	_ = json.NewDecoder(r.Body).Decode(&users)

	collection := helper.ConnectDB()

	result, err := collection.InsertOne(context.TODO(), users)

	if err != nil {
		helper.GetError(err, w)
		return
	}

	json.NewEncoder(w).Encode(result)
}

func updateUsers(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var params = mux.Vars(r)

	id, _ := primitive.ObjectIDFromHex(params["id"])

	var users models.Users

	collection := helper.ConnectDB()

	filter := bson.M{"_id": id}

	_ = json.NewDecoder(r.Body).Decode(&users)

	update := bson.D{
		{"$set", bson.D{
			{"name", users.Name},
			{"email", users.Email},
			{"telephone", users.Telephone},
			{"info", users.Info},
			{"typeuser", users.TypeUser},
		}},
	}

	err := collection.FindOneAndUpdate(context.TODO(), filter, update).Decode(&users)

	if err != nil {
		helper.GetError(err, w)
		return
	}

	users.ID = id

	json.NewEncoder(w).Encode(users)
}

func updateMeasures(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var params = mux.Vars(r)

	id, _ := primitive.ObjectIDFromHex(params["id"])

	var users models.Users

	collection := helper.ConnectDB()

	// Create filter
	filter := bson.M{"_id": id}

	// Read update model from body request
	_ = json.NewDecoder(r.Body).Decode(&users)

	// prepare update model.

	update := bson.D{
		{"$set", bson.D{
			{"measures", bson.D{
				{"weight", users.Measures.Weight},
				{"stature", users.Measures.Stature},
				{"shoulder", users.Measures.Shoulder},
				{"inspired_chest", users.Measures.InspiredChest},
				{"relaxed_arm", users.Measures.RelaxedArm},
				{"thigh", users.Measures.Thigh},
				{"forearm", users.Measures.Forearm},
				{"contracted_arm", users.Measures.ContractedArm},
				{"waist", users.Measures.Waist},
				{"abdomen", users.Measures.Abdomen},
				{"hip", users.Measures.Hip},
				{"leg", users.Measures.Leg},
			}},
		}},
	}

	err := collection.FindOneAndUpdate(context.TODO(), filter, update).Decode(&users)

	if err != nil {
		helper.GetError(err, w)
		return
	}

	users.ID = id

	json.NewEncoder(w).Encode(users)
}

func updateTraning(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var params = mux.Vars(r)

	//Get id from parameters
	id, _ := primitive.ObjectIDFromHex(params["id"])

	//var book models.Book
	var users models.Users

	collection := helper.ConnectDB()

	// Create filter
	filter := bson.M{"_id": id}

	// Read update model from body request
	_ = json.NewDecoder(r.Body).Decode(&users)

	// prepare update model.

	update := bson.D{
		{"$set", bson.D{
			{"traning", bson.D{
				{"sequence", users.Traning.Sequence},
				{"place", users.Traning.Place},
				{"exercise", users.Traning.Exercise},
				{"series", users.Traning.Series},
				{"repetition", users.Traning.Repetition},
				{"charge", users.Traning.Charge},
			}},
		}},
	}

	err := collection.FindOneAndUpdate(context.TODO(), filter, update).Decode(&users)

	if err != nil {
		helper.GetError(err, w)
		return
	}

	users.ID = id

	json.NewEncoder(w).Encode(users)
}

func deleteUsers(w http.ResponseWriter, r *http.Request) {
	// Set header
	w.Header().Set("Content-Type", "application/json")

	// get params
	var params = mux.Vars(r)

	// string to primitve.ObjectID
	id, err := primitive.ObjectIDFromHex(params["id"])

	collection := helper.ConnectDB()

	// prepare filter.
	filter := bson.M{"_id": id}

	deleteResult, err := collection.DeleteOne(context.TODO(), filter)

	if err != nil {
		helper.GetError(err, w)
		return
	}

	json.NewEncoder(w).Encode(deleteResult)
}

func main() {
	//Init Router
	r := mux.NewRouter()

	r.HandleFunc("/personalgamer/users", getUsers).Methods("GET")         //aluno
	r.HandleFunc("/personalgamer/users/{id}", getUserName).Methods("GET") //personal
	r.HandleFunc("/personalgamer/users", createUsers).Methods("POST")
	r.HandleFunc("/personalgamer/users/measures/{id}", updateMeasures).Methods("PUT")
	//problema pra depois se não colocar todos os campos ele coloca um valor vazio
	r.HandleFunc("/personalgamer/users/traning/{id}", updateTraning).Methods("PUT") //personal
	r.HandleFunc("/personalgamer/users/{id}", updateUsers).Methods("PUT")           //personal
	r.HandleFunc("/personalgamer/users/{id}", deleteUsers).Methods("DELETE")        //personal

	log.Fatal(http.ListenAndServe(":8000", r))

}
