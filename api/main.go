package main

import (
	"log"
	"os"

	"github.com/dzeleniak/arnold/controllers"
	database "github.com/dzeleniak/arnold/db"
	"github.com/dzeleniak/arnold/services"
	"github.com/dzeleniak/arnold/stores"
	"github.com/joho/godotenv"
)

var DB_URI = os.Getenv("ARNOLD_DB_URI")
var PORT = os.Getenv("ARNOLD_PORT")

func init() {
	err := godotenv.Load()
	if err != nil {
		panic(err)
	}
}

func main() {
	db, err := database.New(DB_URI)
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

	if PORT == "" {
		PORT = "8080"
	}

	log.Fatal(e.Start(":"+PORT))
}