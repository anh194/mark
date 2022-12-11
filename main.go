package main

import (
	"fmt"
	"net/http"
	"mark/api/essayApi"
	"github.com/gorilla/mux"
)


func main(){
	router := mux.NewRouter()

	router.HandleFunc("/api/v1/health", hello).Methods("GET")
	router.HandleFunc("/api/v1/essay/find", essayApi.FindEssay).Methods("GET")
	router.HandleFunc("/api/v1/essay/getall", essayApi.GetAll).Methods("GET")
	router.HandleFunc("/api/v1/essay/create", essayApi.CreateEssay).Methods("POST")
	router.HandleFunc("/api/v1/essay/update", essayApi.UpdateEssay).Methods("PUT")
	router.HandleFunc("/api/v1/essay/delete", essayApi.DeleteEssay).Methods("DELETE")

	fmt.Println("Starting the server on: 8080")
	err := http.ListenAndServe(":8080", router)

	if err != nil {
		panic(err)
	}
}


// func defaultMux() *http.ServeMux{
// 	mux := http.NewServeMux()
// 	mux.HandleFunc("/", hello)
// 	mux.HandleFunc("/essays", submitEssay)
// 	return mux
// }


func hello(w http.ResponseWriter, r *http.Request){
	fmt.Fprintln(w, "Hello world")
}
