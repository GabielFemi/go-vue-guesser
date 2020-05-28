package main

import (
	"github.com/gabielfemi/go-vue-guesser/guesser"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"time"
)


func main() {
		router := mux.NewRouter()
		router.PathPrefix("/assets/").Handler(http.StripPrefix("/assets/", http.FileServer(http.Dir("assets"))))
		router.HandleFunc("/", guesser.Guess)

		server := &http.Server{
			Addr: "127.0.0.1:7000",
			Handler: router,

			WriteTimeout: 15 * time.Second,
			ReadTimeout: 15 * time.Second,
		}
		log.Println("Guesser App is listening on", server.Addr)
		log.Fatalln(server.ListenAndServe())
}