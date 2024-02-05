package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"

	"github.com/RomeroGabriel/gobrax-challenge/configs"
	"github.com/RomeroGabriel/gobrax-challenge/internal/infra/db"
	"github.com/RomeroGabriel/gobrax-challenge/internal/infra/webserver/handlers"
	"github.com/RomeroGabriel/gobrax-challenge/internal/service"
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"

	// mysql
	_ "github.com/go-sql-driver/mysql"
	// sqlite
	// _ "github.com/mattn/go-sqlite3"
)

func main() {
	configs, err := configs.LoadConfig(".")
	if err != nil {
		panic(err)
	}
	database, err := sql.Open(configs.DBDriver, configs.DBConnection)
	if err != nil {
		panic(err)
	}
	defer database.Close()

	truckDriverDb := db.NewDriverRepository(database)
	truckDb := db.NewTruckRepository(database)
	bindingDb := db.NewDriverTruckBindingRespository(database)

	truckDriverService := service.NewDriverService(truckDriverDb, bindingDb)
	webDriverHandler := handlers.NewWebDriverHandler(*truckDriverService)

	truckService := service.NewTruckService(truckDb, bindingDb)
	webTruckHandler := handlers.NewWebTruckHandler(truckService)

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
	r.Route("/bindings", func(r chi.Router) {
		r.Post("/truck/{idtruck}/driver/{iddriver}", bindingHandler.CreateBinding)
	})

	fmt.Println("Starting web server on port ", configs.WebServerPort)
	if err := http.ListenAndServe(configs.WebServerPort, r); err != nil {
		log.Fatal(err)
	}
}
