package controllers

import (
	"context"
	"encoding/json"
	"log"
	"net/http"

	"go.mongodb.org/mongo-driver/bson"

	"golearn/database"
)

func GetUsers(w http.ResponseWriter, r *http.Request) {
	filter := bson.M{"type": "banner"}
	cur, err := database.Connector.Find(context.TODO(), filter)
	if err != nil {
		log.Println(err.Error())
		defer cur.Close(context.TODO())
	}

	fixedImages := make([]bson.M, 0)

	for cur.Next(context.TODO()) {
		var result bson.M
		err := cur.Decode(&result)

		if err != nil {
			log.Println("error in next ", err.Error())
		}

		log.Println(result)
		fixedImages = append(fixedImages, result)
	}

	w.Header().Set("{Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(fixedImages)
}
