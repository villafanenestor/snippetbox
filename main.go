package main

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
)

func home(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}
	w.Write([]byte("Hello from Snippetbox"))
}

func snippetView(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(r.URL.Query().Get("id"))
	if err != nil || id < 1 {
		http.NotFound(w, r)
		return
	}

	fmt.Fprint(w, "Display a specific snippet with id %d", id)
}

func snippetCreate(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		w.Header().Set("Allow", "POST")
		w.Header().Set("Cache-Control", "public, max-age=31536000")
		//Esto es una forma de enviarlo separado
		// w.WriteHeader(http.StatusMethodNotAllowed)
		// w.Write([]byte("Method not allowed"))
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		w.Header()["X-XSS-Protection"] = []string{"1; mode=block"}

		return
	}
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
