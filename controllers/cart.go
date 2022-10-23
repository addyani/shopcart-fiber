package controllers

import (
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/session"
	"gorm.io/gorm"

	"ilmudata/task1/database"
	"ilmudata/task1/models"
)

type CartController struct {
	// declare variables
	Db    *gorm.DB
	store *session.Store
}

func InitCartController(s *session.Store) *CartController {
	db := database.InitDb()
	return &CartController{Db: db, store: s}
}

// GET /products
func (controller *CartController) GetCart(c *fiber.Ctx) error {
	var carts models.Cart
	id := c.Params("id")
	idn, _ := strconv.Atoi(id)

	err := models.ReadCartById(controller.Db, &carts, idn)
	if err != nil {
		return c.SendStatus(500) // http 500 internal server error
	}

	var cartsFK []models.CartProduct
	errs := models.FindCart(controller.Db, &cartsFK, uint(idn))
	if errs != nil {
		return c.SendStatus(500) // http 500 internal server error
	}

	carts.Total = 0
	for _, num := range cartsFK {
		carts.Total += num.Harga
	}

	//Save Update Harga Total To Db Cart
	errss := models.InsertProductToCart(controller.Db, &carts)
	if errss != nil {
		return c.SendStatus(500) // http 500 internal server error
	}

	var user models.User
	errsss := models.FindUserById(controller.Db, &user, idn)
	if errsss != nil {
		return c.Redirect("/login") // Unsuccessful login (cannot find user)
	}

	// return c.JSON(fiber.Map{
	// 	"Title":    "Keranjang",
	// 	"Users":    user,
	// 	"CartUser": carts,
	// 	"Carts":    cartsFK,
	// })

	return c.Render("cart", fiber.Map{
		"Title":    "Keranjang",
		"Users":    user,
		"CartUser": carts,
		"Carts":    cartsFK,
	})
}

// GET /products
func (controller *CartController) AddCart(c *fiber.Ctx) error {
	params := c.AllParams()
	CartId, _ := strconv.Atoi(params["cartid"])
	ProductId, _ := strconv.Atoi(params["productid"])

	var cart models.Cart
	var product models.Product

	err := models.ReadProductById(controller.Db, &product, ProductId)
	if err != nil {
		return c.SendStatus(500) // http 500 internal server error
	}

	errs := models.ReadCartById(controller.Db, &cart, CartId)
	if errs != nil {
		return c.SendStatus(500) // http 500 internal server error
	}

	// var order models.CartProduct
	// order.Harga = product.Price
	// order.Jumlah = product.Quantity

	// fmt.Println(uint(CartId))
	// fmt.Println(uint(ProductId))
	// errss := models.InsertIntoCart(controller.Db, &order, uint(CartId), uint(ProductId))
	// if errss != nil {
	// 	return c.SendStatus(500) // http 500 internal server error
	// }
	cart.Products = append(cart.Products, &product)
	errss := models.InsertProductToCart(controller.Db, &cart)
	if errss != nil {
		return c.SendStatus(500) // http 500 internal server error
	}

	var new models.CartProduct
	errssss := models.FindCartProduct(controller.Db, &new, uint(CartId), uint(ProductId))
	if errssss != nil {
		return c.SendStatus(500) // http 500 internal server error
	}

	new.Jumlah = new.Jumlah + 1
	new.Harga = float32(new.Jumlah) * product.Price

	new.Image = product.Image
	new.Name = product.Name
	new.Deskripsi = product.Deskripsi
	new.Quantity = product.Quantity
	new.Price = product.Price
	new.Owner = product.Owner

	errsss := models.UpdateCart(controller.Db, &new, uint(CartId), uint(ProductId))
	if errsss != nil {
		return c.SendStatus(500) // http 500 internal server error
	}

	// if succeed
	// idns := strconv.FormatUint(uint64(CartId), 10)
	// return c.Redirect("/products/" + idns)

	// return c.JSON(fiber.Map{
	// 	"Title": "Horeeeeeeeeeeeeeeeeeeeee",
	// })
	idns := strconv.FormatUint(uint64(CartId), 10)
	return c.Redirect("/products/" + idns)
}

// GET /products
func (controller *CartController) AddCartInCart(c *fiber.Ctx) error {
	params := c.AllParams()
	CartId, _ := strconv.Atoi(params["cartid"])
	ProductId, _ := strconv.Atoi(params["productid"])

	var cart models.Cart
	var product models.Product

	err := models.ReadProductById(controller.Db, &product, ProductId)
	if err != nil {
		return c.SendStatus(500) // http 500 internal server error
	}

	errs := models.ReadCartById(controller.Db, &cart, CartId)
	if errs != nil {
		return c.SendStatus(500) // http 500 internal server error
	}

	cart.Products = append(cart.Products, &product)
	errss := models.InsertProductToCart(controller.Db, &cart)
	if errss != nil {
		return c.SendStatus(500) // http 500 internal server error
	}

	var new models.CartProduct
	errssss := models.FindCartProduct(controller.Db, &new, uint(CartId), uint(ProductId))
	if errssss != nil {
		return c.SendStatus(500) // http 500 internal server error
	}

	new.Jumlah = new.Jumlah + 1
	new.Harga = float32(new.Jumlah) * product.Price

	new.Image = product.Image
	new.Name = product.Name
	new.Deskripsi = product.Deskripsi
	new.Quantity = product.Quantity
	new.Price = product.Price
	new.Owner = product.Owner

	errsss := models.UpdateCart(controller.Db, &new, uint(CartId), uint(ProductId))
	if errsss != nil {
		return c.SendStatus(500) // http 500 internal server error
	}

	idns := strconv.FormatUint(uint64(CartId), 10)
	return c.Redirect("/cart/" + idns)
}

// GET /products
func (controller *CartController) MinusInCart(c *fiber.Ctx) error {
	params := c.AllParams()
	CartId, _ := strconv.Atoi(params["cartid"])
	ProductId, _ := strconv.Atoi(params["productid"])

	var cart models.Cart
	var product models.Product

	err := models.ReadProductById(controller.Db, &product, ProductId)
	if err != nil {
		return c.SendStatus(500) // http 500 internal server error
	}

	errs := models.ReadCartById(controller.Db, &cart, CartId)
	if errs != nil {
		return c.SendStatus(500) // http 500 internal server error
	}

	cart.Products = append(cart.Products, &product)
	errss := models.InsertProductToCart(controller.Db, &cart)
	if errss != nil {
		return c.SendStatus(500) // http 500 internal server error
	}

	var new models.CartProduct
	errssss := models.FindCartProduct(controller.Db, &new, uint(CartId), uint(ProductId))
	if errssss != nil {
		return c.SendStatus(500) // http 500 internal server error
	}

	new.Jumlah = new.Jumlah - 1
	new.Harga = float32(new.Jumlah) * product.Price

	new.Image = product.Image
	new.Name = product.Name
	new.Deskripsi = product.Deskripsi
	new.Quantity = product.Quantity
	new.Price = product.Price
	new.Owner = product.Owner

	errsss := models.UpdateCart(controller.Db, &new, uint(CartId), uint(ProductId))
	if errsss != nil {
		return c.SendStatus(500) // http 500 internal server error
	}

	idns := strconv.FormatUint(uint64(CartId), 10)
	return c.Redirect("/cart/" + idns)
}

// GET /products
func (controller *CartController) DeleteInCart(c *fiber.Ctx) error {
	params := c.AllParams()
	CartId, _ := strconv.Atoi(params["cartid"])
	ProductId, _ := strconv.Atoi(params["productid"])

	var cart models.CartProduct

	err := models.DeleteCart(controller.Db, &cart, uint(CartId), uint(ProductId))
	if err != nil {
		return c.SendStatus(500) // http 500 internal server error
	}

	idns := strconv.FormatUint(uint64(CartId), 10)
	return c.Redirect("/cart/" + idns)
}
