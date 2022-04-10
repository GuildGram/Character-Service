package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/GuildGram/Character-Service.git/handlers"
	"github.com/gorilla/mux"
)

func main() {
	l := log.New(os.Stdout, "character-api", log.LstdFlags)

	ch := handlers.NewCharacter(l)

	sm := mux.NewRouter()

	//handle get
	getRouter := sm.Methods(http.MethodGet).Subrouter()
	getRouter.HandleFunc("/getAll", ch.GetCharacters)

	getRouter.HandleFunc("/get{id:[0-9]+}", ch.GetCharacter)

	//handle put
	putRouter := sm.Methods(http.MethodPut).Subrouter()
	putRouter.HandleFunc("/update{id:[0-9]+}", ch.UpdateCharacters)

	//handle add
	postRouter := sm.Methods(http.MethodPost).Subrouter()
	postRouter.HandleFunc("/add", ch.AddCharacter)

	//handle delete
	deleteRouter := sm.Methods(http.MethodDelete).Subrouter()
	deleteRouter.HandleFunc("/delete{id:[0-9]+}", ch.DeleteCharacter)

	//Server stuff for testing, will be deleted soon
	s := &http.Server{
		Addr:         ":9090",
		Handler:      sm,
		IdleTimeout:  120 * time.Second,
		ReadTimeout:  1 * time.Second,
		WriteTimeout: 1 * time.Second,
	}

	go func() {
		err := s.ListenAndServe()
		if err != nil {
			l.Fatal(err)
		}
	}()

	sigChan := make(chan os.Signal)
	signal.Notify(sigChan, os.Interrupt)
	signal.Notify(sigChan, os.Kill)

	sig := <-sigChan
	l.Println("received kill signal, shutting down", sig)

	tc, _ := context.WithTimeout(context.Background(), 30*time.Second)
	s.Shutdown(tc)
}
