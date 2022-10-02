package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

func helloHandler(w http.ResponseWriter, r *http.Request) {
   response := os.Getenv("RESPONSE")
   if len(response) == 0 {
      response = "OpenShift for GO developers! V1.0"
   }
   
   fmt.Fprintln(w, response)
   fmt.Println("Servicing an impatient beginner's request'")   
}

func main() {
    http.HandleFunc("/", helloHandler)
    log.Fatal(http.ListenAndServe(":8080", nil))
}

