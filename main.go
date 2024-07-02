package main

import (
	"fmt"
	"log"
	"net/http"

	routers "github.com/lucifer126741/mongo_app/routers"
)

func main() {
	fmt.Println("MongoDB connection established")
	r := routers.Router()
	fmt.Println("Server is initiating")
	log.Fatal(http.ListenAndServe(":4000", r))
	fmt.Println("Server is running")
}
