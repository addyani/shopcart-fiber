package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/session"
	"github.com/gofiber/template/html"

	"ilmudata/task1/controllers"
	"ilmudata/task1/models"
)

func main() {
	// session
	store := session.New()

	// load template engine
	engine := html.New("./views", ".html")

	app := fiber.New(fiber.Config{
		Views: engine,
	})

	// static
	app.Static("/public", "./public")
	models.InitDbModels()

	// controllers
	// helloController := controllers.InitHelloController(store)
	prodController := controllers.InitProductController(store)
	userController := controllers.InitUserController(store)
	cartController := controllers.InitCartController(store)
	historyController := controllers.InitHistoryController(store)

	user := app.Group("")
	user.Get("/login", userController.Login)
	user.Post("/login", userController.LoginPosted)
	user.Get("/logout", userController.Logout)
	user.Get("/register", userController.Register)
	user.Post("/register", userController.AddRegisteredUser)

	prod := app.Group("/products")
	prod.Get("/", prodController.IndexProduct)
	prod.Get("/:id", userController.AuthVerify, prodController.IndexxProduct)
	prod.Get("/user/:id", userController.AuthVerify, prodController.IndexxxProduct)
	prod.Get("/create/:id", userController.AuthVerify, prodController.AddProduct)
	prod.Post("/create/:id", userController.AuthVerify, prodController.AddPostedProduct)
	prod.Get("/detail/:id", userController.AuthVerify, prodController.GetDetailProduct2)
	prod.Get("/editproduct/:id", userController.AuthVerify, prodController.EditlProduct)
	prod.Post("/editproduct/:id", userController.AuthVerify, prodController.EditlPostedProduct)
	prod.Get("/deleteproduct/:id", userController.AuthVerify, prodController.DeleteProduct)

	cart := app.Group("/cart")
	cart.Get("/:id", userController.AuthVerify, cartController.GetCart)
	cart.Get("/:cartid/product/:productid", userController.AuthVerify, cartController.AddCart)
	cart.Get("/:cartid/product/:productid/redirect", userController.AuthVerify, cartController.AddCartInCart)
	cart.Get("/:cartid/product/:productid/kurang", userController.AuthVerify, cartController.MinusInCart)
	cart.Get("/:cartid/product/:productid/batal", userController.AuthVerify, cartController.DeleteInCart)
	cart.Get("/cekout/:id", userController.AuthVerify, cartController.CekOutCart)

	history := app.Group("/history")
	history.Get("/:id", userController.AuthVerify, historyController.GetHistory)
	history.Get("user/:userid/detail/:id", userController.AuthVerify, historyController.GetDetailHistory)

	app.Listen(":3000")
}
