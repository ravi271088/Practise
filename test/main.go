package main

// Implement the handler functions GetKeys, GetKey, CreateKey, and DeleteKey. Use
// in-memory storage; the keys do not need to be persisted across restart. The key
// structure is given. How you manage the storage is up to you.
//
// The get methods should only return keys that are unexpired. If the current
// time is after the key's expiration, it should not be returned in the
// response.
//
// All handler methods should return json data.
//
// Example requests and responses:
//
// # Get all unexpired keys.
// curl --verbose --request GET \
//		--url 'http://localhost:8080/keys'
// -> 200
//		[{"id":"<uuid>","expires":"<iso date>"},{"id":"<uuid>","expires":"<iso date>"}]
//
// # Get a specific key. If the key is expired, the response should be Not Found/404.
// curl --verbose --request GET \
//		--url 'http://localhost:8080/keys/<uuid>'
// -> 200
//		{"id":"<uuid>","expires":"<iso date>"}
// -> 404
//
// # Create a key with no expiration. The expiration should be one hour from
// # when the key is created.
// curl --verbose --request POST \
//		--url 'http://localhost:8080/keys' \
//		--header 'Content-type: application/json' \
//		--data '{}'
// -> 204
//		{"id":"<uuid>","expires":"<iso date now + 1 hour>"}
//
// # Create a key with explicit expiration. An expires date in the past is accepted.
// curl --verbose --request POST \
//		--url 'http://localhost:8080/keys' \
//		--header 'Content-type: application/json' \
//		--data '{"expires":"<iso date>"}'
// -> 204
//		{"id":"<uuid>","expires":"<iso date>"}
//
// # Delete a specific key.
// curl --verbose --request DELETE \
//		--url 'http://localhost:8080/keys/<uuid>'
// -> 204
// -> 404
//

import (
	//"encoding/json"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"slices"
	"time"

	"github.com/google/uuid"
	"github.com/gorilla/mux"
)

type Key struct {
	Id      string    `json:"id"`
	Expires time.Time `json:"expires"`
}

func main() {
	fmt.Println("Starting server on :8080")
	router := mux.NewRouter()
	sub := router.PathPrefix("/keys").Subrouter()
	sub.HandleFunc("", GetKeys).
		Methods("GET")
	sub.HandleFunc("/{key_id}", GetKey).
		Methods("GET")
	sub.HandleFunc("", CreateKey).
		Methods("POST").
		HeadersRegexp("Content-Type", "application/json")
	sub.HandleFunc("/{key_id}", DeleteKey).
		Methods("DELETE")
	router.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusNotFound)
	})
	srv := &http.Server{
		Handler:      router,
		Addr:         "localhost:8080",
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}
	log.Fatal(srv.ListenAndServe())
}

// Get all keys that are not expired. This should return a JSON array of keys.
func GetKeys(w http.ResponseWriter, r *http.Request) {
	fmt.Println("GetKeys called")
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	responseData := []Key{
		{Id: "123e4567-e89b-12d3-a456-426614174000", Expires: time.Now().Add(-1 * time.Hour)},
		{Id: "123e4567-e89b-12d3-a456-426614174001", Expires: time.Now().Add(2 * time.Hour)},
	}
	fmt.Printf("res %+v\n", responseData)
	for i, key := range responseData {
		if key.Expires.Before(time.Now()) {
			responseData = slices.Delete(responseData, i, 1)
		}
		fmt.Printf("Key ID: %s, Expires: %s\n", key.Id, key.Expires)
	}

	response, err := json.Marshal(responseData)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	fmt.Printf("res %+v\n", string(response))
	log.Println("GetKeys")
}

// Get a key by id. If the key is expired, return 404. This should return a JSON object.
func GetKey(w http.ResponseWriter, r *http.Request) {
	getParams := mux.Vars(r)
	keyID := getParams["key_id"]
	responseData := []Key{
		{Id: "123e4567-e89b-12d3-a456-426614174000", Expires: time.Now().Add(-1 * time.Hour)},
		{Id: "123e4567-e89b-12d3-a456-426614174001", Expires: time.Now().Add(2 * time.Hour)},
	}
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	for _, key := range responseData {
		fmt.Println("Checking key:", key.Id)
		if key.Id == keyID {
			if key.Expires.Before(time.Now()) {
				http.Error(w, "Key not found", http.StatusNotFound)
				return
			}
			responseData1 := Key{Id: key.Id, Expires: key.Expires}
			response, err := json.Marshal(responseData1)
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}
			fmt.Printf("res %+v\n", string(response))
		}
	}

	log.Println("GetKey")
}

// Create a new key with the specified expiration. If no expiration is provided,
// use the default of 1 hour. Return the new key as the response.
// Example: {"expires": "2019-01-01T12:00:00Z"}
func CreateKey(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	getParams := mux.Vars(r)
	exDate := getParams["expires"]
	layout := "2006-01-02 15:04:05"
	parsedTime, err := time.Parse(layout, exDate)
	if err != nil {
		fmt.Println("Error parsing time:", err)
		return
	}
	Data := []Key{
		{Id: uuid.New().String(), Expires: parsedTime},
	}
	response, err := json.Marshal(Data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	fmt.Printf("res %+v\n", string(response))
	log.Println("CreateKey")
}

// Delete the specified key. If the key is expired, return 404.
func DeleteKey(w http.ResponseWriter, r *http.Request) {
	getParams := mux.Vars(r)
	keyID := getParams["key_id"]

	Data := []Key{
		{Id: "123e4567-e89b-12d3-a456-426614174000", Expires: time.Now().Add(-1 * time.Hour)},
		{Id: "123e4567-e89b-12d3-a456-426614174001", Expires: time.Now().Add(2 * time.Hour)},
	}
	fmt.Printf("res %+v\n", Data)
	for i, key := range Data {
		if key.Expires.Before(time.Now()) {
			Data = slices.Delete(Data, i, 1)
		}

		if key.Id == keyID {
			Data = slices.Delete(Data, i, 1)
		}
	}

	response, err := json.Marshal(Data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	fmt.Printf("res %+v\n", string(response))

	w.WriteHeader(http.StatusNoContent)
	log.Println("DeleteKey")
}
