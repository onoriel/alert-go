package models

import "go.mongodb.org/mongo-driver/bson/primitive"

// Alert definition
type Alert struct {
	ID      primitive.ObjectID `json:"id,omitempty" bson:"_id,omitempty"`
	Name    string             `json:"name,omitempty" bson:"name,omitempty"`
	Type    string             `json:"type,omitempty" bson:"type,omitempty"`
	Message string             `json:"message,omitempty" bson:"message,omitempty"`
}
