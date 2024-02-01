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

	truckDb := db.NewTruckRepository(database)
	truckService := service.NewTruckService(truckDb)
	webTruckHandler := handlers.NewWebTruckHandler(truckService)

	r := chi.NewRouter()
	r.Use(middleware.Logger)

	r.Route("/truckdrivers", func(r chi.Router) {
		r.Post("/", webTruckDriverHandler.Create)
		r.Get("/{id}", webTruckDriverHandler.FindById)
		r.Get("/", webTruckDriverHandler.FindAll)
		r.Put("/{id}", webTruckDriverHandler.Update)
		r.Delete("/{id}", webTruckDriverHandler.Delete)
	})

	r.Route("/trucks", func(r chi.Router) {
		r.Post("/", webTruckHandler.Create)
		r.Get("/{id}", webTruckHandler.FindById)
		r.Get("/", webTruckHandler.FindAll)
		r.Put("/{id}", webTruckHandler.Update)
		r.Delete("/{id}", webTruckHandler.Delete)
	})

	fmt.Println("Starting web server on port :8080")
	if err := http.ListenAndServe(":8080", r); err != nil {
		log.Fatal(err)
	}
}
