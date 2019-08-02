package main

import (
	"encoding/json"
	"log"
	"net/http"
)

// Data is a model
type Data struct {
	Year    string `json:"year"`
	Quarter string `json:"quarter"`
	Code    string `json:"code"`
}

// TheData is an array of model Data. Naming convention will be changed later
type TheData []Data

func redirect(w http.ResponseWriter, r *http.Request) {
	// Will make a generic function to verify query param in URL
	yr, ok := r.URL.Query()["yr"]
	if !ok || len(yr[0]) < 1 {
		log.Println("URL Param 'yr' is missing")
		return
	}
	year := yr[0]
	log.Println("URL param yr is: " + string(year)) // If yr is present, then log it out to verify

	qtr, ok := r.URL.Query()["qtr"]
	if !ok || len(qtr[0]) < 1 {
		log.Println("URL Param 'qtr' is missing")
		return
	}
	quarter := qtr[0]
	log.Println("URL param qtr is: " + string(quarter))

	code, ok := r.URL.Query()["code"]
	if !ok || len(code[0]) < 1 {
		log.Println("URL Param 'code' is missing")
		return
	}
	codes := code[0]
	log.Println("URL Param code is: " + string(codes))

	getJSONResponse(w, year, quarter, codes)
}

func getJSONResponse(w http.ResponseWriter, year string, quarter string, codes string) { // Need to have type in parameter
	data := TheData{
		Data{
			Year: year, Quarter: quarter, Code: codes}, // Assigning parameters to property of Data
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(data)
}
func main() {
	// This later can be implemented with gorilla/mux
	http.HandleFunc("/", redirect)
	err := http.ListenAndServe(":9090", nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
