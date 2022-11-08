package main

import (
  "net/http"
  "log"
)
//currentTime := time.Now()

func handleRequests(){
  http.HandleFunc("/api/insert/patient/", insertPatientRequest)
  log.Fatal(http.ListenAndServe(":8080", nil))
}

func main(){
  handleRequests()

}
