package main

import (
	"github.com/skuttleman/tilda-poc/services"
	_ "github.com/skuttleman/tilda-poc/Godeps/_workspace/src/github.com/joho/godotenv/autoload"
	"os"
	"net/http"
	"fmt"
)

func main() {
	http.Handle("/socket.io/", services.Socket())
	http.Handle("/", http.FileServer(http.Dir("public")))
	port := os.Getenv("PORT")
	http.ListenAndServe(":" + port, nil)
	fmt.Println("Server is listening on:", port)
}
