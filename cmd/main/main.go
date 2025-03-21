package main

import(
	"log"
	"net/http"
	"github.com/gorilla/mux"
	_ "gorm.io/driver/postgres"
	"github.com/Kush-316/book_mgmt/pkg/routes"
)

func main(){
	r := mux.NewRouter()
	routes.RegisterBookStoreRoutes(r)
	http.Handle("/", r)
	log.Fatal(http.ListenAndServe("localhost:9010", r))
}