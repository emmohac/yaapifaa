package main

import (
    "log"
    "net/http"
    "strings"
    "encoding/json"
)

func redirect(w http.ResponseWriter, r *http.Request) {
    yr, ok := r.URL.Query()["yr"]
    if !ok || len(yr[0]) < 1 {
        log.Println("URL Param 'yr' is missing")
        return
    }
    year := yr[0]
	log.Println("URL param yr is: " + string(year))

    qtr, ok := r.URL.Query()["qtr"]
    if !ok || len(qtr[0]) < 1 {
        log.Println("URL Param 'qtr' is missing")
        return
    }
    quarter := qtr[0]
	log.Println("URL param qtr is: "+string(quarter))

    code, ok := r.URL.Query()["code"]
    if !ok || len(code[0]) < 1 {
        log.Println("URL Param 'code' is missing")
        return
    }
    codes := code[0]
	log.Println("URL Param code is: "+string(codes))
    getJsonResponse(w, year, quarter, codes)
}

func getJsonResponse(w http.ResponseWriter, theYear, theQuarter, theCode) {
  data := SomeStruct{}
  w.Header().Set("Content-Type", "application/json")
  w.WriteHeader(http.StatusCreated)
  json.NewEncoder(w).Encode(data)
}

func main() {
    http.HandleFunc("/", redirect)
    err := http.ListenAndServe(":9090", nil)
    if err != nil {
        log.Fatal("ListenAndServe: ", err)
    }
}
