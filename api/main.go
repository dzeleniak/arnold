package main

import (
	"log"
	"os"

	"github.com/dzeleniak/arnold/controllers"
	database "github.com/dzeleniak/arnold/db"
	"github.com/dzeleniak/arnold/services"
	"github.com/dzeleniak/arnold/stores"
)

var GO_ENV = os.Getenv("GO_ENV")

func main() {
	db, err := database.New(GO_ENV == "development")
	if err != nil {
		log.Fatal("Failed to connect to database")
	}

	defer db.Close();

	e := controllers.Echo()

	s := stores.New(db)
	ss := services.New(s)
	c := controllers.New(ss)

	controllers.SetDefault(e);
	controllers.SetApi(e, c, nil);

	PORT := os.Getenv("PORT")

	if PORT == "" {
		PORT = "8888"
	}

	log.Fatal(e.Start(":"+PORT))
}