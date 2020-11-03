package main

//imports
import (
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

// Book struct (Model)

type Book struct {
	ID     string  `json:"id"`
	Isbn   string  `json:"idbn"`
	Title  string  `json:"title"`
	Author *Author `json:"author"`
}

// Author struct
type Author struct {
	Firstname string `json:"firstname"`
	Lastname  string `json:"lastname"`
}

// Init books var as a slice Book struct
var books []Book

func getBooks(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "applciation/json")
	json.NewEncoder(w).Encode(books)
}

// Get single book
func getBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r) // Gets params
	// Loop through books and find one with the id from the params
	for _, item := range books {
		if item.ID == params["id"] {
			json.NewEncoder(w).Encode(item)
			return
		}
	}
	json.NewEncoder(w).Encode(&Book{})
}

//add a new book
func createBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var book Book
	_ = json.NewDecoder(r.Body).Decode(&book)
	book.ID = strconv.Itoa(rand.Intn(100000000))
	books = append(books, book)
	json.NewEncoder(w).Encode(book)
}

// Update book
func updateBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	for index, item := range books {
		if item.ID == params["id"] {
			books = append(books[:index], books[index+1:]...)
			var book Book
			_ = json.NewDecoder(r.Body).Decode(&book)
			book.ID = params["id"]
			books = append(books, book)
			json.NewEncoder(w).Encode(book)
			return
		}
	}
}

// Delete book
func deleteBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	for index, item := range books {
		if item.ID == params["id"] {
			books = append(books[:index], books[index+1:]...)
			break
		}
	}
	json.NewEncoder(w).Encode(books)
}

//parameter function
func foo(w http.ResponseWriter, req *http.Request) {
	v := req.FormValue("q")
	// v1 := req.FormValue("p")
	// i, _ := strconv.Atoi(v)
	// j, _ := strconv.Atoi(v1)
	// k := i + j
	fmt.Fprintln(w, "Do My Search:", v)

}

// Main function
func main() {
	// Init router
	r := mux.NewRouter()

	// Hardcoded data
	books = append(books, Book{ID: "1", Isbn: "423434", Title: "Book One", Author: &Author{Firstname: "John", Lastname: "Doe"}})
	books = append(books, Book{ID: "2", Isbn: "134543", Title: "Book Two", Author: &Author{Firstname: "Steve", Lastname: "Smith"}})
	books = append(books, Book{ID: "3", Isbn: "441234", Title: "Book Three", Author: &Author{Firstname: "Alex", Lastname: ""}})
	// Route handles & endpoints
	r.HandleFunc("/books", getBooks).Methods("GET")
	r.HandleFunc("/books/{id}", getBook).Methods("GET")
	r.HandleFunc("/books", createBook).Methods("POST")
	r.HandleFunc("/books/{id}", updateBook).Methods("PUT")
	r.HandleFunc("/books/{id}", deleteBook).Methods("DELETE")
	r.HandleFunc("/Calculation", foo).Methods("GET")

	//http handler for foo(parameter adder)
	http.HandleFunc("/", foo)
	http.Handle("/favico.ico", http.NotFoundHandler())

	//Start Server
	log.Fatal(http.ListenAndServe(":8080", r))

	/*
		Links:
		127.0.0.1:8080/books
		127.0.0.1:8080/Calculation

		To Add Parameters
		http://127.0.0.1:8080/Calculation?q=something
		multiple
		127.0.0.1:8080/calculation?q=hello+world


	*/
}
