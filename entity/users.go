package enitity

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type FixedImages struct {
	id        primitive.ObjectID `json:"_id" bson:"_id"`
	name      string             `json:"name" bson:"name"`
	img_url   string             `json:"img_url" bson:"img_url"`
	createdAt int                `json:"createdAt" bson:"createdAt"`
	updatedAt int                `json:"updatedAt" bson:"updatedAt"`
}
