package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/RomeroGabriel/gobrax-challenge/internal/infra/db"
	"github.com/RomeroGabriel/gobrax-challenge/internal/infra/webserver/handlers"
	"github.com/RomeroGabriel/gobrax-challenge/internal/service"
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/jmoiron/sqlx"
	_ "github.com/mattn/go-sqlite3"
)

func main() {
	database := sqlx.MustOpen("sqlite3", ":memory:")
	defer database.Close()

	truckDriverDb := db.NewTruckDriverRepository(database)
	truckDriverService := service.NewTruckDriverService(truckDriverDb)

	webTruckDriverHandler := handlers.NewWebTruckDriverHandler(truckDriverService)

	r := chi.NewRouter()
	r.Use(middleware.Logger)

	r.Route("/truckdrivers", func(r chi.Router) {
		r.Post("/", webTruckDriverHandler.Create)
		r.Get("/{id}", webTruckDriverHandler.FindById)
		r.Get("/", webTruckDriverHandler.FindAll)
	})

	fmt.Println("Starting web server on port :8080")
	if err := http.ListenAndServe(":8080", r); err != nil {
		log.Fatal(err)
	}
}
