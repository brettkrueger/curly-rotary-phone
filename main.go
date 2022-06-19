package main

import (
	"api/api"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"
	"time"

	"github.com/gorilla/mux"
)

var Data = api.ApiData()

var mappedCharacters = map[string]api.Character{
	"aang":     Data[0],
	"appa":     Data[1],
	"momo":     Data[2],
	"katara":   Data[3],
	"sokka":    Data[4],
	"toph":     Data[5],
	"zuko":     Data[6],
	"suki":     Data[7],
	"iroh":     Data[8],
	"ozai":     Data[9],
	"azula":    Data[10],
	"zhao":     Data[11],
	"longfeng": Data[12],
}

type Status struct {
	Status   string    `json:"status"`
	Datetime time.Time `json:"datetime"`
}

func status(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Printing status")
	status := Status{"ok", time.Now()}
	json.NewEncoder(w).Encode(status)
}

func character(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	raw_name := params["name"]
	char := mappedCharacters[raw_name]

	fmt.Printf("Printing character Name: %s, %d\n", char.Name, char.Id)
	json.NewEncoder(w).Encode(char)
}

func characterId(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	raw_id := params["id"]

	id, err := strconv.Atoi(raw_id)
	newId := id - 1

	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}

	fmt.Printf("Printing character ID: %d, %s\n", id, Data[newId].Name)
	json.NewEncoder(w).Encode(Data[newId])
}

func characters(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Printing characters")
	json.NewEncoder(w).Encode(Data)
}

func main() {
	fmt.Println("Starting http server on port 8080")
	router := mux.NewRouter()
	router.HandleFunc("/", status)
	router.HandleFunc("/status", status)
	router.HandleFunc("/characters", characters).Methods("GET")
	router.HandleFunc("/characters/{name}", character).Methods("GET")
	router.HandleFunc("/character/{id}", characterId).Methods("GET")

	log.Fatal(http.ListenAndServe(":8080", router))
	fmt.Println("Shutting down http server")
}
