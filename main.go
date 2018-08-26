package main

import (
	"go_saas/controller"
	"log"
	"net/http"
)

func main() {
	api := &controllers.API{}

	if err := http.ListenAndServe(":8080", api); err != nil {
		log.Println(err)
	}
}
