package app

import (
	"go-api-template/app/database"
	"log"
	"net/http"
	"os"
)

func Boostrap() {
	database.InitDatabase()

	mux := http.NewServeMux()

	e := http.ListenAndServe(os.Getenv("PORT"), mux)

	log.Fatal(e)
}
