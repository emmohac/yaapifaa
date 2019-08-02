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
    yr := yr[0]

    qtr, ok := r.URL.Query()["qtr"]
    if !ok || len(qtr[0]) < 1 {
        log.Println("URL Param 'qtr' is missing")
        return
    }
    qtr := qtr[0]

    code, ok := r.URL.Query()["code"]
    if !ok || len(code[0]) < 1 {
        log.Println("URL Param 'code' is missing")
        return
    }
    code := code[0]

    getJsonResponse(w, yr, qtr, code)
}

func getJsonResponse(w http.ResponseWriter, yr, qtr, code) {
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
