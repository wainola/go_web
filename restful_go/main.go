package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/gorilla/mux"
)

type Note struct {
	Title       string    `json:"title"`
	Descripcion string    `json:"descripcion"`
	CreatedOn   time.Time `json:"createdon"`
}

// runtime storage for notes
var noteStore = make(map[string]Note)

// variable to generate key for collection
var id int = 0

// HTTP Post - /api/notes
func PostNoteHandler(w http.ResponseWriter, r *http.Request) {
	var note Note
	fmt.Println("Imprimiendo el r.Body")
	fmt.Println(r.Body)
	// Decoding incoming json note
	err := json.NewDecoder(r.Body).Decode(&note)
	if err != nil {
		panic(err)
	}
	fmt.Println("Imprimiendo err")
	fmt.Println(err)
	note.CreatedOn = time.Now()
	id++
	k := strconv.Itoa(id)
	noteStore[k] = note
	j, err := json.Marshal(note)
	if err != nil {
		panic(err)
	}
	fmt.Println("imprimiendo j")
	fmt.Println(string(j[:]))
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	w.Write(j)
}

// HTTP GET - /api/notes
func GetNoteHandler(w http.ResponseWriter, r *http.Request) {
	var notes []Note
	for _, v := range noteStore {
		notes = append(notes, v)
	}
	fmt.Println(notes)

	w.Header().Set("Content-Type", "application/json")
	j, err := json.Marshal(notes)
	if err != nil {
		panic(err)
	}
	w.WriteHeader(http.StatusOK)
	w.Write(j)
}

// HTTP Put - /api/notes/{id}
func PutNoteHandler(w http.ResponseWriter, r *http.Request) {
	var err error
	vars := mux.Vars(r)
	fmt.Println("vars es...")
	fmt.Println(vars)
	// obteniendo el id que es pasado mediante la url
	k := vars["id"]
	var noteToUpdate Note
	// Decodificando la nota que viene en json
	err = json.NewDecoder(r.Body).Decode(&noteToUpdate)
	fmt.Println("la nota decodificada es...")
	fmt.Println(noteToUpdate)
	if err != nil {
		panic(err)
	}
	if note, ok := noteStore[k]; ok {
		noteToUpdate.CreatedOn = note.CreatedOn
		// borrando el item existente y añadiendo el item actualizado
		delete(noteStore, k)
		noteStore[k] = noteToUpdate
	} else {
		log.Printf("No pude encontrar la llave para la nota %s para borrarla", k)
	}
	w.WriteHeader(http.StatusNoContent)
}

// HTTP Delete - /api/notes/{id}
func DeleteNoteHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	k := vars["id"]
	// Remove item from store
	if _, ok := noteStore[k]; ok {
		// delete existing item
		delete(noteStore, k)
	} else {
		log.Printf("No pude encontrar la llave para la nota %s para borrarla", k)
	}
	w.WriteHeader(http.StatusNoContent)
}

// Entry point of the program
func main() {
	r := mux.NewRouter().StrictSlash(false)
	r.HandleFunc("/api/notes", GetNoteHandler).Methods("GET")
	r.HandleFunc("/api/notes", PostNoteHandler).Methods("POST")
	r.HandleFunc("/api/notes/{id}", PutNoteHandler).Methods("PUT")
	r.HandleFunc("/api/notes/{id}", DeleteNoteHandler).Methods("DELETE")
	server := &http.Server{
		Addr:    ":8080",
		Handler: r,
	}
	log.Println("Listening...")
	server.ListenAndServe()
}
