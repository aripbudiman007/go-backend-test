package main

import (
	"cart/app"
	"cart/controller"
	"cart/helper"
	"cart/repository"
	"cart/service"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func main() {
	db := app.NewBD()
	cartRepository := repository.NewCartRepository()
	cartegoryService := service.NewCartService(cartRepository, db)
	categoryController := controller.NewCartController(cartegoryService)

	router := httprouter.New()

	router.GET("/api/cart", categoryController.Get)
	router.POST("/api/cart", categoryController.Save)
	router.GET("/api/cart/:cartId", categoryController.FindById)
	router.DELETE("/api/cart/:cartId", categoryController.Delete)

	server := http.Server{
		Addr:    "localhost:3000",
		Handler: router,
	}

	err := server.ListenAndServe()
	helper.PanicIfError(err)
}
