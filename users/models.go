package users

import (
	"context"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

//User basic structure
type User struct {
	ID    primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	FName string             `json:"fname,omitempty" bson:"fname,omitempty"`
	LName string             `json:"lname,omitempty" bson:"lname,omitempty"`
	Email string             `json:"email,omitempty" bson:"email,omitempty"`
}

//AllUsers returns a slice of Users
func AllUsers(cl *mongo.Collection) []User {
	//initialize a slice model to get data
	var Users []User
	//creates a context with a timeout of 3 seconds
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	//cancel de ctx, all jobs done
	defer cancel()
	//gets the cursos with data
	cursor, err := cl.Find(ctx, bson.D{})
	if err != nil {
		cancel()
		log.Fatal(err)
	}
	defer cursor.Close(ctx)
	// loop throght the cursor decoding the data and append it to the slice model
	for cursor.Next(ctx) {
		//initialize a model user to receive data from the cursor
		user := User{}
		//decode cursor data into user
		err = cursor.Decode(&user)
		if err != nil {
			cancel()
			log.Fatal(err)
		}
		//append the results into the slice of users
		Users = append(Users, user)
	}
	if err := cursor.Err(); err != nil {
		log.Fatal(err)
	}
	//returns the slice model
	return Users
}

//OneUser returns the ObjectID user
func OneUser(cl *mongo.Collection, id primitive.ObjectID) User {
	//initialize the model to decoded the mongo data
	user := User{}
	//creates a context with a timeout of 3 seconds
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	//cancel de ctx, all jobs done
	defer cancel()
	//gets the patient related to id and decode to the pointe patient model
	err := cl.FindOne(ctx, bson.M{"_id": id}).Decode(&user)
	if err != nil {
		cancel()
		return user
	}
	//returns the patient
	return user
}

//CreateUser creates a user and returns the create user
func CreateUser(cl *mongo.Collection, u User) primitive.ObjectID {
	//creates a context with a timeout of 3 seconds
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	//cancel de ctx, all jobs done
	defer cancel()
	//creates a new patient on the collection
	res, err := cl.InsertOne(ctx, u)
	if err != nil {
		cancel()
	}
	//decode response
	obj := res.InsertedID.(primitive.ObjectID)
	//returns the objectID of the created user
	return obj
}

//DeleteUser deletes the user of given id
func DeleteUser(cl *mongo.Collection, id primitive.ObjectID) int64 {
	//creates a context with a timeout of 3 seconds
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	//cancel de ctx, all jobs done
	defer cancel()
	//delete the patient from the collection
	res, err := cl.DeleteOne(ctx, bson.M{"_id": id})
	if err != nil {
		cancel()
	}
	//return the number of deleted documents
	return res.DeletedCount
}

//ModifyUser replace the user given on pos 0 from slice by the user on pos 1
func ModifyUser(cl *mongo.Collection, u []User) primitive.ObjectID {
	//creates a context with a timeout of 3 seconds
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	//cancel de ctx, all jobs done
	defer cancel()
	//Replace the data on the collection
	res, err := cl.ReplaceOne(ctx, bson.M{"_id": u[0].ID}, u[1])
	if err != nil {
		cancel()
	}
	//decode response
	obj := res.UpsertedID.(primitive.ObjectID)
	//returns the objectID of the created user
	return obj
}
