package routes

import(
	"github.com/gorilla/mux"
	"github.com/freydrey/go-bookstore/pkg/controllers"
)

var RegisterBookStoreRoutes = func(route *mux.Router){
	route.HandleFunc("/book/", controllers.CreateBook).Methods("POST")
	route.HandleFunc("/book/" , controllers.GetBook).Methods("GET")
	route.HandleFunc("/book/{id}", controllers.GetBookById).Methods("GET")
	route.HandleFunc("/book/{id}", controllers.UpdateBook).Methods("PUT")
	route.HandleFunc("/book/{id}", controllers.DeleteBook).Methods("DELETE")

}