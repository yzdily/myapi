package main
import (
	"fmt"
	"log"
	"net/http"
	"github.com/gorilla/mux"
)
func getUsers(w http.ResponseWriter, r *http.Request){
	fmt.Fprintf(w, "Get All Users Endpoint Hit")
}

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/api/users", getUsers).Methods("Get")
	log.Fatal(http.ListenAndServe(":8000", router))
}