package entity

import "go.mongodb.org/mongo-driver/bson/primitive"

type Role struct {
	Id     primitive.ObjectID `json:"id" bson:"_id,omitempty"`
	Code   string             `json:"code" bson:"code,omitempty"`
	Scopes []string           `json:"scopes" bson:"scopes,omitempty"`
}
