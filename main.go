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
	"github.com/rs/cors"
)

func main() {
	// go handlers.StartMsgBrokerConnection()

	//old code
	l := log.New(os.Stdout, "character-api", log.LstdFlags)

	ch := handlers.NewCharacter(l)

	router := mux.NewRouter()

	//handle get
	getRouter := router.Methods(http.MethodGet).Subrouter()
	getRouter.HandleFunc("/characters/getall", ch.GetCharacters)

	//should change to get by name for when user services are implemented
	getRouter.HandleFunc("/characters/get{id:[0-9]+}", ch.GetCharacter)

	//handle put
	putRouter := router.Methods(http.MethodPut).Subrouter()
	putRouter.HandleFunc("/characters/update{id:[0-9]+}", ch.UpdateCharacters)
	putRouter.HandleFunc("/characters/updateguild{id:[0-9]+}", ch.UpdateCharacterGuild)

	//handle Post
	postRouter := router.Methods(http.MethodPost).Subrouter()
	postRouter.HandleFunc("/characters/add", ch.AddCharacter)

	//handle delete
	deleteRouter := router.Methods(http.MethodDelete).Subrouter()
	deleteRouter.HandleFunc("/characters/delete{id:[0-9]+}", ch.DeleteCharacter)

	//handle open rabbit mq listener
	msgBrokerRouter := router.Methods(http.MethodGet).Subrouter()
	msgBrokerRouter.HandleFunc("/characters/msg{id:[0-9]+}", ch.SendCharactersMessageBroker)

	c := cors.New(cors.Options{
		AllowedOrigins:   []string{"http://localhost:3000"},
		AllowCredentials: true,
		AllowedMethods:   []string{"POST", "DELETE", "GET", "PUT"},
	})

	handler := c.Handler(router)

	//Server stuff
	s := &http.Server{
		Addr:         ":9090",
		Handler:      handler,
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
