package controllers

import (
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/session"
	"gorm.io/gorm"

	"ilmudata/task1/database"
	"ilmudata/task1/models"

	"golang.org/x/crypto/bcrypt"
)

type UserController struct {
	// declare variables
	Db    *gorm.DB
	store *session.Store
}

func InitUserController(s *session.Store) *UserController {
	db := database.InitDb()
	return &UserController{Db: db, store: s}
}

// GET /login
func (controller *UserController) Login(c *fiber.Ctx) error {
	return c.Render("login", fiber.Map{
		"Title": "Login",
	})
}

// post /login
func (controller *UserController) LoginPosted(c *fiber.Ctx) error {
	sess, err := controller.store.Get(c)
	if err != nil {
		panic(err)
	}
	var user models.User
	var myform models.LoginForm

	if err := c.BodyParser(&myform); err != nil {
		return c.Redirect("/login")
	}

	// Find user
	errs := models.FindUserByUsername(controller.Db, &user, myform.Username)
	if errs != nil {
		return c.Redirect("/login") // Unsuccessful login (cannot find user)
	}

	// Compare password
	compare := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(myform.Password))
	if compare == nil { // compare == nil artinya hasil compare di atas true
		sess.Set("username", user.Username)
		sess.Set("userId", user.Id)
		sess.Save()

		idn := strconv.FormatUint(uint64(user.Id), 10)
		return c.Redirect("/products/" + idn)
	}

	return c.Redirect("/login")
}

// GET /register
func (controller *UserController) Register(c *fiber.Ctx) error {
	return c.Render("register", fiber.Map{
		"Title": "Register",
	})
}

// POST /register
func (controller *UserController) AddRegisteredUser(c *fiber.Ctx) error {
	var user models.User
	// var cart models.Cart

	if err := c.BodyParser(&user); err != nil {
		return c.SendStatus(400) // Bad Request, RegisterForm is not complete
	}

	// Hash password
	bytes, _ := bcrypt.GenerateFromPassword([]byte(user.Password), 10)
	sHash := string(bytes)

	// Simpan hashing, bukan plain passwordnya
	user.Password = sHash

	// save user
	err := models.CreateUser(controller.Db, &user)
	if err != nil {
		return c.SendStatus(500) // Server error, gagal menyimpan user
	}

	// Find user
	errs := models.FindUserByUsername(controller.Db, &user, user.Username)
	if errs != nil {
		return c.SendStatus(500) // Server error, gagal menyimpan user
	}

	// // also create cart
	// errCart := models.CreateCart(controller.Db, &cart, user.ID)
	// if errCart != nil {
	// 	return c.SendStatus(500) // Server error, gagal menyimpan user
	// }

	// if succeed
	return c.Redirect("/login")
}

// /logout
func (controller *UserController) Logout(c *fiber.Ctx) error {

	sess, err := controller.store.Get(c)
	if err != nil {
		panic(err)
	}
	sess.Destroy()
	return c.Render("login", fiber.Map{
		"Title": "Login",
	})
}

func (controller *UserController) AuthVerify(c *fiber.Ctx) error {
	sess, _ := controller.store.Get(c)
	val := sess.Get("username")
	if val != nil {
		return c.Next()
	}
	return c.Redirect("/login")
}
