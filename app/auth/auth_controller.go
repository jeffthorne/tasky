package auth

import (
	"context"
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/jeffthorne/tasky/app/database"
	"github.com/jeffthorne/tasky/app/models"
	"github.com/jinzhu/copier"
	"go.mongodb.org/mongo-driver/bson"
)

func Auth(c *fiber.Ctx) error {
	return c.Render("login", fiber.Map{"User": LoginForm{}})
}

func Register(c *fiber.Ctx) error {
	return c.Render("index", fiber.Map{"User": SignupForm{}})
}

type LoginForm struct {
	Email    string `json:"Email" validate:"email,required" form:"Email"`
	Password string `json:"Password" validate:"required,min=5,max=32" form:"Password"`
}

func Login(c *fiber.Ctx) error {

	fmt.Println("IN LOGIN")
	loginForm := LoginForm{}
	if err := c.BodyParser(&loginForm); err != nil {
		fmt.Println("ERROR1: ", err)
		return c.Render("login", fiber.Map{"User": loginForm, "Errors": err})
	}

	if err := ValidateStruct(loginForm, "login"); err != nil {
		return c.Render("login", fiber.Map{"User": loginForm, "Errors": err})
	}

	user := models.User{}
	err := database.DB.Collection("Users").FindOne(context.TODO(), bson.D{{"Email", loginForm.Email}}).Decode(&user)
	errors := []map[string]string{}
	if err != nil {
		errors = append(errors, map[string]string{"Login": "invalid"})
		return c.Render("login", fiber.Map{"User": loginForm, "Errors": errors})
	}

	if !user.CheckPassword(loginForm.Password) {
		errors = append(errors, map[string]string{"Login": "invalid"})
		return c.Render("login", fiber.Map{"User": loginForm, "Errors": errors})
	}

	errors = append(errors, map[string]string{"Login": "good"})
	return c.Render("login", fiber.Map{"User": loginForm, "Errors": errors})

}

type SignupForm struct {
	FirstName       string `json:"FirstName" validate:"required" form:"FirstName"`
	LastName        string `json:"LastName" validate:"required" form:"LastName"`
	Email           string `json:"Email" validate:"email,required" form:"Email"`
	Password        string `json:"Password" validate:"required,min=5,max=32" form:"Password"`
	ConfirmPassword string `json:"ConfirmPassword" validate:"required,min=5,max=32" form:"ConfirmPassword"`
}

func Signup(c *fiber.Ctx) error {
	fmt.Println("IN SIGNUP")
	signupForm := SignupForm{}
	if err := c.BodyParser(&signupForm); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"message": err})
	}

	fmt.Println("In signup: ", signupForm)
	if err := ValidateStruct(signupForm, "signup"); err != nil {
		return c.Render("index", fiber.Map{"User": signupForm, "Errors": err})
	}
	fmt.Println("In signup: ", signupForm)
	user := models.User{}
	if err := copier.Copy(&user, &signupForm); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"message": err})
	}

	user.HashPassword()
	result, err := database.DB.Collection("Users").InsertOne(context.TODO(), user)
	if err != nil {
		fmt.Println("INSERT ERROR: ", err)
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"message": err})
	}

	fmt.Println("INSERT RESULT: ", result)
	errors := []map[string]string{}
	errors = append(errors, map[string]string{"Signup": "successful"})

	return c.Render("index", fiber.Map{"User": SignupForm{}, "Errors": errors})
}
