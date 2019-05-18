package users

import (
	"fmt"
	"net/http"

	"github.com/fusco2k/go-microservice-user-di/config"
	"github.com/julienschmidt/httprouter"
)

//UserIndex returns all users
func UserIndex(env *config.Env) httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
		users := AllUsers(env.CL)

		for _, user := range users {
			fmt.Fprintf(w, "%s, %s, %s, %s\n", user.ID, user.FName, user.LName, user.Email)
		}
	}
}
