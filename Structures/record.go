package Structures

import "go.mongodb.org/mongo-driver/bson/primitive"

type Records struct {
	ID      primitive.ObjectID `json:"id,omitempty"`
	Problem string             `json:"problem,omitempty"`
	Status  bool               `json:"status,omitempty"`
}
