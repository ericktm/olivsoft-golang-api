package main

import (
	"log"
	"net/http"
	"github.com/ericktm/olivsoft-golang-api/util"
)

func main() {
	log.Println("Application Start")
	defer log.Println("Application Stop")

	app := util.GetConfig()

	log.Fatal(http.ListenAndServe(":8080", app.Router))

	defer app.DB.Close()
}
