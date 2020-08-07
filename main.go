package main

import (
	"log"
	"net/http"

	envConfig "github.com/gfbatista/xy-inc/util"

	"github.com/gfbatista/xy-inc/client"
	"github.com/gfbatista/xy-inc/database"
)

func main() {
	envConfig.LoadConfig()
	database.Create()

	http.HandleFunc("/poi/proximidade/", client.ListByProximity)
	http.HandleFunc("/poi/", client.ListInsert)
	log.Fatal(http.ListenAndServe(":3000", nil))
}
