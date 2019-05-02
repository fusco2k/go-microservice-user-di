package main

import (
	"net/http"

	"github.com/fusco2k/go-microservice-user-di/config"
	"go.mongodb.org/mongo-driver/mongo"
)

type Env struct {
	cl *mongo.Collection
}

func main() {
	cl := config.NewSession("mongodb://localhost:27017", "users", "catalog")

	env := Env{cl}

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request){
		http.Redirect(w,r,"/users", 307)
	})
	http.HandleFunc("users", func(w http.ResponseWriter, r *http.Request){
		switch r.Method{
		case "GET":
			
		case "POST":

		case "PUT":

		case "DELETE":
			
		}
	})
	http.ListenAndServe("8080", nil)
}
