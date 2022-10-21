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

type ProductController struct {
	// declare variables
	Db *gorm.DB
}

func InitProductController() *ProductController {
	db := database.InitDb()
	return &ProductController{Db: db}
}

// routing
// GET /products
func (controller *ProductController) IndexProduct(c *fiber.Ctx) error {
	// load all products
	var products []models.Product
	err := models.ReadProducts(controller.Db, &products)
	if err != nil {
		return c.SendStatus(500) // http 500 internal server error
	}
	return c.Render("products", fiber.Map{
		"Title":    "Daftar Produk",
		"Products": products,
	})
}

// GET /products
func (controller *ProductController) IndexxProduct(c *fiber.Ctx) error {
	// load all products
	var products []models.Product
	var user models.User
	id := c.Params("id")
	idn, _ := strconv.Atoi(id)

	errs := models.FindUserById(controller.Db, &user, idn)
	if errs != nil {
		return c.Redirect("/login") // Unsuccessful login (cannot find user)
	}

	errss := models.ReadProductByNoUser(controller.Db, &products, user.Username)
	if errss != nil {
		return c.Redirect("/login") // Unsuccessful login (cannot find user)
	}
	//if succeed
	return c.Render("products", fiber.Map{
		"Title":    "Rincian",
		"Users":    user,
		"Products": products,
	})
}

// GET /products
func (controller *ProductController) IndexxxProduct(c *fiber.Ctx) error {
	// load all products
	var products []models.Product
	var user models.User
	id := c.Params("id")
	idn, _ := strconv.Atoi(id)

	errs := models.FindUserById(controller.Db, &user, idn)
	if errs != nil {
		return c.Redirect("/login") // Unsuccessful login (cannot find user)
	}

	errss := models.ReadProductByUser(controller.Db, &products, user.Username)
	if errss != nil {
		return c.Redirect("/login") // Unsuccessful login (cannot find user)
	}
	//if succeed
	return c.Render("products", fiber.Map{
		"Title":    "Rincian",
		"Users":    user,
		"Products": products,
	})
}

// GET /products/create
func (controller *ProductController) AddProduct(c *fiber.Ctx) error {
	id := c.Params("id")
	idn, _ := strconv.Atoi(id)
	var user models.User

	// Find user
	errs := models.FindUserById(controller.Db, &user, idn)
	if errs != nil {
		return c.Redirect("/login") // Unsuccessful login (cannot find user)
	}

	return c.Render("addproduct", fiber.Map{
		"Title": "Tambah Produk",
		"users": user,
	})
}

// POST /products/create
func (controller *ProductController) AddPostedProduct(c *fiber.Ctx) error {
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

	// save product
	err := models.CreateProduct(controller.Db, &myform)
	if err != nil {
		return c.Redirect("/products")
	}

	errs := models.FindUserById(controller.Db, &user, idn)
	if errs != nil {
		return c.Redirect("/login") // Unsuccessful login (cannot find user)
	}

	var products []models.Product
	errss := models.ReadProductByUser(controller.Db, &products, user.Username)
	if errss != nil {
		return c.Redirect("/login") // Unsuccessful login (cannot find user)
	}
	// if succeed
	return c.Render("products", fiber.Map{
		"Title":    "Rincian",
		"users":    user,
		"Products": products,
	})
}

// GET /products/productdetail?id=xxx
func (controller *ProductController) GetDetailProduct(c *fiber.Ctx) error {
	id := c.Query("id")
	idn, _ := strconv.Atoi(id)

	var product models.Product
	err := models.ReadProductById(controller.Db, &product, idn)
	if err != nil {
		return c.SendStatus(500) // http 500 internal server error
	}
	return c.Render("productdetail", fiber.Map{
		"Title":   "Detail Produk",
		"Product": product,
	})
}

// GET /products/detail/xxx
func (controller *ProductController) GetDetailProduct2(c *fiber.Ctx) error {
	id := c.Params("id")
	idn, _ := strconv.Atoi(id)

	var product models.Product
	err := models.ReadProductById(controller.Db, &product, idn)
	if err != nil {
		return c.SendStatus(500) // http 500 internal server error
	}
	return c.Render("productdetail", fiber.Map{
		"Title":   "Detail Produk",
		"Product": product,
	})
}

// / GET products/editproduct/xx
func (controller *ProductController) EditlProduct(c *fiber.Ctx) error {
	id := c.Params("id")
	idn, _ := strconv.Atoi(id)

	var product models.Product
	err := models.ReadProductById(controller.Db, &product, idn)
	if err != nil {
		return c.SendStatus(500) // http 500 internal server error
	}
	return c.Render("editproduct", fiber.Map{
		"Title":   "Edit Produk",
		"Product": product,
	})
}

// / POST products/editproduct/xx
func (controller *ProductController) EditlPostedProduct(c *fiber.Ctx) error {
	id := c.Params("id")
	idn, _ := strconv.Atoi(id)

	var product models.Product
	var myform models.Product

	err := models.ReadProductById(controller.Db, &product, idn)
	if err != nil {
		return c.SendStatus(500) // http 500 internal server error
	}

	if err := c.BodyParser(&myform); err != nil {
		return c.Redirect("/products")
	}

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
			product.Image = (exp + file.Filename)
		}
		//return err
	}
	//return c.SendStatus(400)

	product.Name = myform.Name
	product.Deskripsi = myform.Deskripsi
	product.Quantity = myform.Quantity
	product.Price = myform.Price
	// save product
	models.UpdateProduct(controller.Db, &product)

	return c.Redirect("/products")

}

// / GET /products/deleteproduct/xx
func (controller *ProductController) DeleteProduct(c *fiber.Ctx) error {
	id := c.Params("id")
	idn, _ := strconv.Atoi(id)

	var product models.Product
	models.DeleteProductById(controller.Db, &product, idn)
	return c.Redirect("/products")
}
