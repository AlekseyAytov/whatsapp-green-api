package main

import (
	"log"
	"net/http"

	"github.com/AlekseyAytov/whatsapp-green-api/internal"
)

func main() {
	log.Fatal(http.ListenAndServe(":8080", internal.NewAdapter().Router()))
}
