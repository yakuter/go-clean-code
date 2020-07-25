package main

import (
	"fmt"
	"log"
	"os"

	"net/http"

	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/yakuter/go-clean-code/pkg/api"
	"github.com/yakuter/go-clean-code/pkg/repository/post"
	"github.com/yakuter/go-clean-code/pkg/service"
)

// App is alias for api.App{}
type App struct {
	Router *mux.Router
	DB     *gorm.DB
}

func main() {
	a := App{}

	// Initialize storage
	a.initialize(
		os.Getenv("APP_DB_HOST"),
		os.Getenv("APP_DB_PORT"),
		os.Getenv("APP_DB_USERNAME"),
		os.Getenv("APP_DB_PASSWORD"),
		os.Getenv("APP_DB_NAME"))

	// Initialize routes
	a.routes()

	// Start server
	a.run(":8010")
}

func (a *App) initialize(host, port, username, password, dbname string) {
	var err error

	connectionString :=
		fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", host, port, username, password, dbname)

	a.DB, err = gorm.Open("postgres", connectionString)
	if err != nil {
		log.Fatal(err)
	}

	// repository := post.NewRepository(a.DB)

	a.Router = mux.NewRouter()
}

func (a *App) run(addr string) {
	fmt.Printf("Server started at %s\n", addr)
	log.Fatal(http.ListenAndServe(addr, a.Router))
}

func (a *App) routes() {
	postAPI := InitPostAPI(a.DB)
	a.Router.HandleFunc("/posts", postAPI.FindAllPosts()).Methods("GET")
	a.Router.HandleFunc("/posts", postAPI.CreatePost()).Methods("POST")
	a.Router.HandleFunc("/posts/{id:[0-9]+}", postAPI.FindByID()).Methods("GET")
	a.Router.HandleFunc("/posts/{id:[0-9]+}", postAPI.UpdatePost()).Methods("PUT")
	a.Router.HandleFunc("/posts/{id:[0-9]+}", postAPI.DeletePost()).Methods("DELETE")
}

// InitPostAPI ..
func InitPostAPI(db *gorm.DB) api.PostAPI {
	postRepository := post.NewRepository(db)
	postService := service.NewPostService(postRepository)
	postAPI := api.NewPostAPI(postService)
	postAPI.Migrate()
	return postAPI
}
