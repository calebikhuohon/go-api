package http

import (
	"github.com/gorilla/mux"
	"go-api/server"
	"go-api/server/http/handlers"
	user2 "go-api/server/http/handlers/user"
	"go-api/services/user"
	"go-api/services/user/repository"
	"log"
	"net/http"
)

var (
	userHandler user2.Handler
)

func init() {
	mongoDB, err := server.ConnectToMongo()
	fatalIfErr(err)

	userRepository := repository.NewMongoRepository(mongoDB)

	userService := user.NewService(userRepository)
	userHandler = user2.NewUserHandler(userService)

}

func fatalIfErr(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

//Server configures and returns a new http.Server
func Server() *http.Server {
	r := mux.NewRouter()

	r.Handle("/users", handlers.Handler(userHandler.CreateUser)).Methods(http.MethodPost)
	r.Handle("/users/{userId}", handlers.Handler(userHandler.GetUser)).Methods(http.MethodGet)

	srv := &http.Server{Handler: r, Addr: ":8080"}
	return srv
}
