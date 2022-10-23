package controllers

import (
	"fmt"
	"strconv"
	"time"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"

	"ilmudata/task1/database"
	"ilmudata/task1/models"
)

type TestingController struct {
	// declare variables
	Db *gorm.DB
}

func InitTestingController() *TestingController {
	db := database.InitDb()
	return &TestingController{Db: db}
}

// Testing
func (controller *TestingController) GetProductUser(c *fiber.Ctx) error {
	// load all products
	var users []models.User
	err := models.GetAllProduct(controller.Db, &users)
	if err != nil {
		return c.SendStatus(500) // http 500 internal server error
	}
	return c.JSON(fiber.Map{
		"Title": "Daftar Produk",
		"nilai": users,
	})
}

// Testing
func (controller *TestingController) GetProductUser2(c *fiber.Ctx) error {
	// load all products
	var users models.User
	id := c.Params("id")
	idn, _ := strconv.Atoi(id)

	err := models.GetAllProductUser(controller.Db, &users, idn)
	if err != nil {
		return c.SendStatus(500) // http 500 internal server error
	}
	return c.JSON(fiber.Map{
		"Title": "Daftar Produk",
		"nilai": users,
	})
}

// POST /products/create
func (controller *TestingController) PostAddProd(c *fiber.Ctx) error {
	//myform := new(models.Product)
	var myform models.Product

	id := c.Params("id")
	idn, _ := strconv.Atoi(id)
	var user models.User

	if form, err := c.MultipartForm(); err == nil {
		// => *multipart.Form

		// Get all files from "documents" key:
		files := form.File["image"]
		// => []*multipart.FileHeader

		// Loop through files:
		for _, file := range files {
			fmt.Println(file.Filename, file.Size, file.Header["Content-Type"][0])
			// => "tutorial.pdf" 360641 "application/pdf"
			exp := time.Now().Format("20060102150405")

			// Save the files to disk:
			if err := c.SaveFile(file, fmt.Sprintf("./public/images/%s", (exp+file.Filename))); err != nil {
				return err
			}
			//return c.SendString("Succeed.. " + (exp + file.Filename))
			myform.Image = (exp + file.Filename)
		}
		//return err
	}
	//return c.SendStatus(400)

	if err := c.BodyParser(&myform); err != nil {
		return c.Redirect("/products")
	}

	errs := models.FindUserById(controller.Db, &user, idn)
	if errs != nil {
		return c.Redirect("/login") // Unsuccessful login (cannot find user)
	}

	myform.Owner = (user.Name)
	myform.UserIdProduct = uint(user.Id)

	// save product
	err := models.CreateProduct(controller.Db, &myform)
	if err != nil {
		return c.Redirect("/products")
	}

	// if succeed
	//idns := strconv.FormatUint(uint64(user.Id), 10)
	return c.JSON(fiber.Map{
		"Title": "Daftar Produk",
		"nilai": myform,
	})
}

func (controller *TestingController) GetAllCart(c *fiber.Ctx) error {
	// load all products
	var carts []models.Cart
	// err := models.ReadCart(controller.Db, &carts)
	// if err != nil {
	// 	return c.SendStatus(500) // http 500 internal server error
	// }
	return c.JSON(fiber.Map{
		"Title": "Daftar Produk",
		"nilai": carts,
	})
}

func (controller *CartController) TestingAddCart(c *fiber.Ctx) error {
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

	var order models.CartProduct
	order.Harga = product.Price
	order.Jumlah = product.Quantity
	// errss := models.InsertIntoCart(controller.Db, &order, uint(CartId), uint(ProductId))
	// if errss != nil {
	// 	return c.Redirect("/products")
	// }

	// errss := models.InsertProductToCart(controller.Db, &cart, &product)
	// if errss != nil {
	// 	return c.SendStatus(500) // http 500 internal server error
	// }

	// if succeed
	// idns := strconv.FormatUint(uint64(CartId), 10)
	// return c.Redirect("/products/" + idns)

	return c.JSON(fiber.Map{
		"Title": "Horeeeeeeeeeeeeeeeeeeeee",
	})
}
