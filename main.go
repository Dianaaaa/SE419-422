package main

import (
	"log"
	"math/rand"
	"net/http"
	"time"

	"github.com/julienschmidt/httprouter"
)

func init() {
	rand.Seed(time.Now().UTC().UnixNano())
}

func main() {
	InitDB();
	router := httprouter.New()
	router.GET("/get-url", Geturl)
	router.POST("/generate", Generate)
	router.GET("/heart-beat", Heartbeat)

	log.Fatal(http.ListenAndServe(":8000", router))
}