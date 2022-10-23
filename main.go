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
	testingController := controllers.InitTestingController()

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
	cart.Get("/:id", cartController.GetCart)
	cart.Get("/:cartid/product/:productid", cartController.AddCart)
	cart.Get("/:cartid/product/:productid/redirect", cartController.AddCartInCart)
	cart.Get("/:cartid/product/:productid/kurang", cartController.MinusInCart)
	cart.Get("/:cartid/product/:productid/batal", cartController.DeleteInCart)

	user := app.Group("")
	user.Get("/login", userController.Login)
	user.Post("/login", userController.LoginPosted)
	user.Get("/logout", userController.Logout)
	user.Get("/register", userController.Register)
	user.Post("/register", userController.AddRegisteredUser)

	test := app.Group("/testing")
	test.Post("/create/:id", testingController.PostAddProd)
	test.Post("/allcart", testingController.GetAllCart)

	// //app.Get("/testing", userController.userTest)

	// app.Get("/login", authController.Login)
	// app.Post("/login", authController.LoginPosted)
	// app.Get("/logout", authController.Logout)
	// //app.Get("/profile",authController.Profile)

	// // app.Use("/profile", func(c *fiber.Ctx) error {
	// // 	sess,_ := store.Get(c)
	// // 	val := sess.Get("username")
	// // 	if val != nil {
	// // 		return c.Next()
	// // 	}

	// // 	return c.Redirect("/login")

	// // })
	// app.Get("/profile", func(c *fiber.Ctx) error {
	// 	sess, _ := store.Get(c)
	// 	val := sess.Get("username")
	// 	if val != nil {
	// 		return c.Next()
	// 	}

	// 	return c.Redirect("/login")

	// }, authController.Profile)

	app.Listen(":3000")
}
