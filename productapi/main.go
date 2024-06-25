package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"productapi/handlers"
	"time"
)

func main() {

	// http.HandleFunc("/", func(rw http.ResponseWriter, r *http.Request) {
	// 	log.Println("Handle request")
	// })

	l := log.New(os.Stdout, "prouct-api", log.LstdFlags)
	ph := handlers.NewProduct(l)
	sm := http.NewServeMux()
	sm.Handle("/", ph)

	s := &http.Server{

		Addr:         ":8010",
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

	sigcha := make(chan os.Signal)
	signal.Notify(sigcha, os.Interrupt)
	signal.Notify(sigcha, os.Kill)

	sig := <-sigcha
	l.Println("recived terminate ,graceful shutdown", sig)

	tc, _ := context.WithTimeout(context.Background(), 120*time.Second)
	s.Shutdown(tc)

}
