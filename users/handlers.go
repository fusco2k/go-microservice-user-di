package users

import (
	"encoding/json"
	"fmt"
	"net/http"

	"go.mongodb.org/mongo-driver/bson/primitive"

	"github.com/fusco2k/go-microservice-user-di/config"
	"github.com/julienschmidt/httprouter"
)

//Index returns all users
func Index(env *config.Env) httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
		//get all users from db
		users := AllUsers(env.CL)
		//loop throht users slice and prints
		for _, user := range users {
			fmt.Fprintf(w, "%s, %s, %s, %s\n", user.ID, user.FName, user.LName, user.Email)
		}
	}
}

//Get returns a user by its id users
func Get(env *config.Env) httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
		//parse the param id and ask for the correspondent user
		id, _ := primitive.ObjectIDFromHex(ps.ByName("id"))
		user := OneUser(env.CL, id)
		//prints the user
		fmt.Fprintf(w, "%s, %s, %s, %s\n", user.ID, user.FName, user.LName, user.Email)
	}
}

//Create a user with the given JSON data
func Create(env *config.Env) httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
		u := User{}
		json.NewDecoder(r.Body).Decode(&u)
		obj := CreateUser(env.CL, u)
		//prints the id of the created user
		fmt.Fprintf(w, "create a user with the following id: %s", obj.Hex())
	}
}

//Delete a user by its id users
func Delete(env *config.Env) httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
		//parse the param id and ask for the correspondent user
		id, _ := primitive.ObjectIDFromHex(ps.ByName("id"))
		res := DeleteUser(env.CL, id)
		//prints the number of documents deleted
		fmt.Fprintf(w, "number of documents deleted: %v", res)
	}
}

//Modify a user
func Modify(env *config.Env) httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
		u := []User{}
		//decode the json with the a slice with the user to be modified
		json.NewDecoder(r.Body).Decode(&u)
		obj := ModifyUser(env.CL, u)
		//prints the id of the modified user
		fmt.Fprintf(w, "modified the user with the following id: %s", obj.Hex())
	}
}
