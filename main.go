package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/demo/demo-go-service/handlers"
)

func main() {
	l := log.New(os.Stdout, "demo-go-service ", log.LstdFlags)

	helloHandler := handlers.NewHello(l)
	goodByeHandler := handlers.NewGoodBye(l)
	productsHandler := handlers.NewProducts(l)

	serveMux := http.NewServeMux()

	serveMux.Handle("/", helloHandler)
	serveMux.Handle("/goodbye", goodByeHandler)
	serveMux.Handle("/get-products", productsHandler)

	s := &http.Server{
		Addr:         ":8080",
		Handler:      serveMux,
		IdleTimeout:  120 * time.Second,
		ReadTimeout:  1 * time.Second,
		WriteTimeout: 1 * time.Second,
	}

	s.ListenAndServe()

	go func() {
		l.Println("Starting the server for demo-go-service ...")
		err := s.ListenAndServe()
		if err != nil {
			l.Fatal(err)
		}
	}()

	sigChan := make(chan os.Signal)
	signal.Notify(sigChan, os.Interrupt)
	signal.Notify(sigChan, os.Kill)

	sig := <-sigChan
	l.Println("Received terminate, graceful shutdown", sig)

	tc, _ := context.WithTimeout(context.Background(), 30*time.Second)
	s.Shutdown(tc)

}
