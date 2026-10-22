package main

import (
	"fmt"
	"log"
	"net/http"
)

func home(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello from Snippetbox"))
}

func snippetView(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Display a specific snippet..."))
}

func snippetCreate(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Create a new snippet..."))
}

func main() {
	fmt.Println("Iniciando aplicación")

	//Se usa http.NewServeMux para registrar las rutas
	mux := http.NewServeMux()
	mux.HandleFunc("/", home)
	mux.HandleFunc("/snippet/create", snippetCreate)
	mux.HandleFunc("/snippet/view", snippetView)

	//Se usar http.ListenAnd Serve() para iniciar el servidor web

	log.Print("Iniciando servidor en el puerto :4000")
	err := http.ListenAndServe(":4000", mux)
	log.Fatal(err)

}
