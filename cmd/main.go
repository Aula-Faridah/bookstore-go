package main

import (
	"bdpit/bookstore-go/internals/routes"
	"bdpit/bookstore-go/pkg"
	"log"

	_"github.com/jmoiron/sqlx"
)

// Dependency Injection (DI)

func main()  {
	// Inisialisasi DB
	db, err := pkg.InitMySql()
	if err != nil {
		log.Fatal(err)
	}
	// Inisialisasi Router
	router := routes.InitRouter(db)

	// Inisialisasi Server
	server := pkg.InitServer(router)

	// Running Server
	if err := server.ListenAndServe(); err != nil {
		log.Fatal(err)
	}

}