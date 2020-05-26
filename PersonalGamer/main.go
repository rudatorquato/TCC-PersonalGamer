package main

import (
	//. "PersonalGamer/models"
	"PersonalGamer/helper"
	"PersonalGamer/models"
	"image/png"
	"os"

	"context"
	"encoding/json"
	"log"
	"net/http"

	"github.com/boombuler/barcode"
	"github.com/boombuler/barcode/qr"
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
	log.Println(users.Traning)
	log.Println(users.Traning.Place)
}

func getUserQrcode(w http.ResponseWriter, r *http.Request) {

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
	exercícios, err := json.Marshal(users.Traning)
	//log.Println(string(training))

	//qrCode, _ := qr.Encode(users.Traning.Place, qr.M, qr.Auto)
	qrCode, _ := qr.Encode(string(exercícios), qr.M, qr.Auto)
	//log.Println(users.Traning.Place)

	// Scale the barcode to 200x200 pixels
	qrCode, _ = barcode.Scale(qrCode, 200, 200)

	// create the output file
	file, _ := os.Create("qrcode/qrcode.png")
	defer file.Close()

	// encode the barcode as png
	png.Encode(file, qrCode)
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
			{"username", users.Username},
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
				{"right_relaxed_arm", users.Measures.RightRelaxedArm},
				{"left_relaxed_arm", users.Measures.LeftRelaxedArm},
				{"left_thigh", users.Measures.LeftThigh},
				{"right_thigh", users.Measures.RightThigh},
				{"left_forearm", users.Measures.LeftForearm},
				{"right_forearm", users.Measures.RightForearm},
				{"left_contracted_arm", users.Measures.LeftContractedArm},
				{"right_contracted_arm", users.Measures.RightContractedArm},
				{"waist", users.Measures.Waist},
				{"abdomen", users.Measures.Abdomen},
				{"hip", users.Measures.Hip},
				{"left_leg", users.Measures.LeftLeg},
				{"right_leg", users.Measures.RightLeg},
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
				{"images", users.Traning.Images},
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

	r.HandleFunc("/personalgamer/users", getUsers).Methods("GET")            //aluno
	r.HandleFunc("/personalgamer/users/{id}", getUserName).Methods("GET")    //personal
	r.HandleFunc("/personalgamer/QRCODE/{id}", getUserQrcode).Methods("GET") //personal
	r.HandleFunc("/personalgamer/users", createUsers).Methods("POST")
	r.HandleFunc("/personalgamer/users/measures/{id}", updateMeasures).Methods("PUT")
	r.HandleFunc("/personalgamer/users/traning/{id}", updateTraning).Methods("PUT") //personal
	r.HandleFunc("/personalgamer/users/{id}", updateUsers).Methods("PUT")           //personal
	r.HandleFunc("/personalgamer/users/{id}", deleteUsers).Methods("DELETE")        //personal

	log.Fatal(http.ListenAndServe(":8000", r))

}
