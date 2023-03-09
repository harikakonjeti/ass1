package main

import (
	"api/Config"
	"api/Handler"
	"log"

	"net/http"

	"github.com/bmizerany/pat"
)

func main() {
	

	app := Config.App{
		Mp:   make(map[string]int),
		V:    make([]Config.Video, 0),
		Size: 0,
	}
	repo := Handler.NewRepo(&app)
	Handler.NewHandler(repo)
	mux := pat.New()
	mux.Get("/:id", http.HandlerFunc(Handler.Repo.GetViews))
	mux.Post("/:id", http.HandlerFunc(Handler.Repo.IncViews))
	http.Handle("/", mux)
	err := http.ListenAndServe(":12345", nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	} else {
		log.Println("Port running on :12345")
	}
}
