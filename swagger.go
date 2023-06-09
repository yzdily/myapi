package main

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
)

type response struct {
	Message string `json:"message"`
}

// hello godoc
// @Summary 返回问候消息
// @Description 返回一个带有问候信息的JSON响应
// @Tags hello
// @Produce json
// @Success 200 {object} response
// @Router /hello [get]
func hello(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response{
		Message: "Hello, world!",
	})
}

func main() {
	router := mux.NewRouter()

	router.HandleFunc("/hello", hello).Methods("GET")

	// 将 Swagger UI 目录映射到 /swagger URL
	router.PathPrefix("/swagger/").Handler(http.StripPrefix("/swagger/", http.FileServer(http.Dir("./swagger-ui-master"))))

	http.ListenAndServe(":8000", router)
}