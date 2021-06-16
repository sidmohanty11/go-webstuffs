package controllers

import (
	"log"

	"golang.org/x/crypto/bcrypt"

	"github.com/gofiber/fiber/v2"
	"github.com/sidmohanty11/go-webstuffs/setting-up-backend/db"
	"github.com/sidmohanty11/go-webstuffs/setting-up-backend/models"
)

func Register(c *fiber.Ctx) error {
	var data map[string]string

	err := c.BodyParser(&data)

	if err != nil {
		log.Fatalln(err.Error())
	}

	password, _ := bcrypt.GenerateFromPassword([]byte(data["password"]), 14)

	user := models.User{
		Name:     data["name"],
		Email:    data["email"],
		Password: password,
	}

	db.DB.Create(&user)

	return c.JSON(user)
}

func Login(c *fiber.App) error {
	var data map[string]string

	err := c.BodyParser(&data)

	if err != nil {
		log.Fatalln(err.Error())
	}

	var user models.User

	db.DB.Where("email = ?", data["email"]).First(&user)

	if user.ID == 0 {
		c.Status(fiber.StatusNotFound)
		return c.JSON(fiber.Map{
			"message": "user not found",
		})
	}

	err = bcrypt.CompareHashAndPassword(user.Password, []byte(data["password"]))

	if err != nil {
		log.Fatalln(err.Error())
		c.Status(fiber.StatusBadRequest)
		return c.JSON(fiber.Map{
			"message": "incorrect password",
		})
	}

	return c.JSON(user)
}
