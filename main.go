package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	_ "github.com/lib/pq"
)

func main() {
	// Create a file server handler for serving static files from the "frontend" directory
	fs := http.FileServer(http.Dir("frontend"))
	http.Handle("/", fs)
	http.HandleFunc("/messages", Messages)
	// Print message before starting the server
	fmt.Println("Listening on port 8080")

	// Start the server
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Printf("Error starting server: %s\n", err)
	}
}

func Messages(w http.ResponseWriter, req *http.Request) {
	db, err := sql.Open("postgres", "postgresql://postgres:boobs@172.17.0.2/postgres?sslmode=disable")
	if err != nil {
		http.Error(w, fmt.Sprintf("Error connecting to database: %s", err), http.StatusInternalServerError)
		return
	}
	defer db.Close()
    
    if req.Method == http.MethodGet {
        // Handle GET request
	rows, err := db.Query("SELECT * FROM messages;")
	if err != nil {
		http.Error(w, fmt.Sprintf("Error querying database: %s", err), http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	// Iterate over the rows and store the data in a slice of structs
	type Message struct {
		ID   int    `json:"id"`
		Message string `json:"message"`
		DateCreated time.Time `json:date_created` 
	}
	var data []Message

	for rows.Next() {
		var d Message		
		if err := rows.Scan(&d.ID, &d.Message, &d.DateCreated); err != nil {
			http.Error(w, fmt.Sprintf("Error scanning row: %s", err), http.StatusInternalServerError)
			return
		}
		data = append(data, d)
	}
	if err := rows.Err(); err != nil {
		http.Error(w, fmt.Sprintf("Error iterating over rows: %s", err), http.StatusInternalServerError)
		return
	}

	// Marshal the data into JSON
	jsonData, err := json.Marshal(data)
	if err != nil {
		http.Error(w, fmt.Sprintf("Error marshaling JSON: %s", err), http.StatusInternalServerError)
		return
	}

	// Set the content type header and write the JSON data to the response
	w.Header().Set("Content-Type", "application/json")
	w.Write(jsonData)
    } 
}



