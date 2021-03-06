package Database

import "go.mongodb.org/mongo-driver/bson/primitive"

type Note struct {
	ID     primitive.ObjectID `json:"id,omitempty"`
	Task   string             `json:"problem,omitempty"`
	Status bool               `json:"status,omitempty"`
}
