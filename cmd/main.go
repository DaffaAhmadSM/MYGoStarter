package main

import (
	"MyGoStarterPack/internal/database"
	"MyGoStarterPack/internal/helpers"
	"MyGoStarterPack/internal/http/routers"
	"github.com/joho/godotenv"
	"log"
	"net/http"
)

func main() {
	// LOAD .ENV FILE
	err := godotenv.Load()
	if err != nil {
		panic(err)
	}

	// INIT DATABASE
	db := database.PostgresInit()
	defer func() {
		debe, _ := db.DB()
		err := debe.Close()
		if err != nil {
			panic(err)
		}
	}()
	// INIT SERVER
	e := helpers.EchoSetup(routers.RouterInit(db))

	e.Static("/assets", "./web/static")
	serverInit := http.Server{
		Addr:    ":8080",
		Handler: e,
	}
	log.Println("Server running on http://127.0.0.1:8080")
	log.Fatal(serverInit.ListenAndServe())

}
