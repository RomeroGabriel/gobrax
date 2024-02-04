package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"

	"github.com/RomeroGabriel/gobrax-challenge/internal/infra/db"
	"github.com/RomeroGabriel/gobrax-challenge/internal/infra/webserver/handlers"
	"github.com/RomeroGabriel/gobrax-challenge/internal/service"
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	_ "github.com/mattn/go-sqlite3"
)

func main() {
	database, err := sql.Open("sqlite3", ":memory:")
	if err != nil {
		panic(err)
	}
	defer database.Close()

	truckDriverDb := db.NewTruckDriverRepository(database)
	truckDriverService := service.NewTruckDriverService(truckDriverDb)
	webDriverHandler := handlers.NewWebDriverHandler(truckDriverService)

	truckDb := db.NewTruckRepository(database)
	truckService := service.NewTruckService(truckDb)
	webTruckHandler := handlers.NewWebTruckHandler(truckService)

	bindingDb := db.NewDriverTruckBindingRespository(database)
	bindingService := service.NewDriverTruckBindingService(truckDriverDb, truckDb, bindingDb)
	bindingHandler := handlers.NewDriverTruckBindingHandler(bindingService)

	r := chi.NewRouter()
	r.Use(middleware.Logger)

	r.Route("/drivers", func(r chi.Router) {
		r.Post("/", webDriverHandler.Create)
		r.Get("/{id}", webDriverHandler.FindById)
		r.Get("/", webDriverHandler.FindAll)
		r.Put("/{id}", webDriverHandler.Update)
		r.Delete("/{id}", webDriverHandler.Delete)
	})

	r.Route("/trucks", func(r chi.Router) {
		r.Post("/", webTruckHandler.Create)
		r.Get("/{id}", webTruckHandler.FindById)
		r.Get("/", webTruckHandler.FindAll)
		r.Put("/{id}", webTruckHandler.Update)
		r.Delete("/{id}", webTruckHandler.Delete)
	})

	r.Route("/drivers-trucks", func(r chi.Router) {
		r.Post("/", bindingHandler.CreateBinding)
	})

	fmt.Println("Starting web server on port :8080")
	if err := http.ListenAndServe(":8080", r); err != nil {
		log.Fatal(err)
	}
}
