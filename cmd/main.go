package main

import (
	"bdpit/bookstore-go/internals/routes"
	"bdpit/bookstore-go/pkg"
	"log"
)

// Dependency Injection (DI)

func main()  {
	// Inisialisasi DB
	_, err := pkg.InitMySql()
	if err != nil {
		log.Fatal(err)
	}
	// Inisialisasi Router
	router := routes.InitRouter()

	// Inisialisasi Server
	server := pkg.InitServer(router)

	// Running Server
	if err := server.ListenAndServe(); err != nil {
		log.Fatal(err)
	}

}