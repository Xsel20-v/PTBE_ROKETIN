package main

import (
	"log"
	"net/http"
	"ptbe_roketin/challenge_2/routers"
)

func main() {
    r := routers.InitRouter()

    log.Println("Server running on http://localhost:8080")
    err := http.ListenAndServe(":8080", r)
    if err != nil {
        log.Fatalf("Server failed to start: %v", err)
    }
}