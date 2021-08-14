package Structures

import "go.mongodb.org/mongo-driver/bson/primitive"

type Note struct {
	ID      primitive.ObjectID `json:"id,omitempty"`
	Problem string             `json:"problem,omitempty"`
	Status  bool               `json:"status,omitempty"`
}
