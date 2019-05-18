package main

import (
	"context"
	"log"
	"net/http"

	"github.com/fusco2k/go-microservice-user-di/config"
	"github.com/fusco2k/go-microservice-user-di/users"
	"github.com/julienschmidt/httprouter"
)

func main() {
	//mux
	router := httprouter.New()
	//dependency config
	cl := config.NewSession("mongodb://localhost:27017")
	defer cl.Disconnect(context.Background())
	env := &config.Env{CL: cl.Database("users").Collection("catalog")}
	//routes
	router.GET("/api/users", users.UserIndex(env))
	// router.POST("/api/users")
	// router.PUT("/api/users")
	// router.DELETE("/api/users")
	// router.GET("/api/users/:id", users.UserIndex(env))
	//server at port 8080 using httprouter as router
	log.Fatal(http.ListenAndServe(":8080", router))
}
