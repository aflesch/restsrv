package main

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

type Test struct {
	ID   string `json:"id,omitempty"`
	Name string `json:"name,omitempty"`
}

var test []Test

// our main function
func main() {
	test = append(test, Test{ID: "1", Name: "foo"})
	test = append(test, Test{ID: "2", Name: "how"})
	test = append(test, Test{ID: "3", Name: "bobo"})

	router := mux.NewRouter()
	router.HandleFunc("/test", GetTest).Methods("GET")
	router.HandleFunc("/test/{id}", GetTestId).Methods("GET")
	router.HandleFunc("/test/{id}", CreateTestId).Methods("POST")
	router.HandleFunc("/test/{id}", DeleteTestId).Methods("DELETE")
	log.Fatal(http.ListenAndServe(":8000", router))
}

func GetTest(w http.ResponseWriter, r *http.Request) {
	fmt.Print("GET All", r.RemoteAddr, "\n")
	json.NewEncoder(w).Encode(test)
}

func GetTestId(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	fmt.Print("GET", r.RemoteAddr, "\n")
	for _, item := range test {
		if item.ID == params["id"] {
			json.NewEncoder(w).Encode(item)
			return
		}
	}
	json.NewEncoder(w).Encode(&Test{})
}

func CreateTestId(w http.ResponseWriter, r *http.Request) {
	fmt.Print("Create", r.RemoteAddr, "\n")
	params := mux.Vars(r)
	var tst Test
	_ = json.NewDecoder(r.Body).Decode(&tst)
	tst.ID = params["id"]
	test = append(test, tst)
	json.NewEncoder(w).Encode(test)
}

func DeleteTestId(w http.ResponseWriter, r *http.Request) {
	fmt.Print("Delete", r.RemoteAddr, "\n")
	params := mux.Vars(r)
	for index, item := range test {
		if item.ID == params["id"] {
			test = append(test[:index], test[index+1:]...)
			break
		}
		json.NewEncoder(w).Encode(test)
	}
}
