package main

import (
	"EcomerceProject/internmal/products"
	"log"
	"net/http"
	"time"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
)

type aplicaction struct {
	config config
}
type config struct {
	addr string
	db   dbConfig
}

type dbConfig struct {
	dsn string
}

func (app *aplicaction) run(h http.Handler) error {
	srv := &http.Server{
		Addr:         app.config.addr,
		Handler:      h,
		WriteTimeout: time.Second * 30,
		ReadTimeout:  time.Second * 10,
		IdleTimeout:  time.Minute,
	}
	log.Printf("server has started at addr %s", app.config.addr)
	return srv.ListenAndServe()
}

func (app *aplicaction) mount() http.Handler {
	r := chi.NewRouter()
	r.Use(middleware.RequestID)
	r.Use(middleware.RealIP)
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)

	r.Use(middleware.Timeout(60 * time.Second))

	r.Get("/health", func(w http.ResponseWriter, r *http.Request) { w.Write([]byte("bye")) })
	// http.ListenAndServe(":3333", r)
	productsService:= products.NewService()
	productHandler := products.NewHandler(productsService)
	r.Get("/products", productHandler.ListProducts)
	return r
}
