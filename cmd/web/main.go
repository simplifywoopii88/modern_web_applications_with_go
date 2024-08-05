package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/simplifywoopii88/modern_web_applications_with_go/pkg/handlers"
)

const portNumber = ":8080"



func main(){
	http.HandleFunc("/", handlers.Home)
	http.HandleFunc("/about", handlers.About)
	fmt.Printf("listening server http://localhost%s\n", portNumber)
	if err := http.ListenAndServe(portNumber, nil); err != nil {
		log.Fatal(err)
	}
}