package users

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type Users struct {
	ID    primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	FName string             `json:"fname,omitempty" bson:"fname,omitempty"`
	LName string             `json:"lname,omitempty" bson:"lname,omitempty"`
	Email string             `json:"email,omitempty" bson:"email,omitempty"`
}

func AllUsers(cl *mongo.Collection) {

}

func OneUser(cl *mongo.Collection) {

}

func CreateUser(cl *mongo.Collection) {

}

func DeleteUser(cl *mongo.Collection) {

}

func ModifyUser(cl *mongo.Collection) {

}
