package handler

import (
	"backend-koperasi/database"
	"backend-koperasi/models/entity"
	"backend-koperasi/models/request"
	"backend-koperasi/utils"
	"log"
	"time"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
)

func LoginHandler(ctx *fiber.Ctx) error {
	LoginRequest := new(request.LoginRequest)

	if err := ctx.BodyParser(LoginRequest); err != nil {
		return err
	}

	// log.Println(LoginRequest)

	// VALIDASI API LOGIN
	validate := validator.New()
	errValidator := validate.Struct(LoginRequest)
	if errValidator != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error":   true,
			"message": errValidator.Error(),
		})
	}

	// CHECK AVAILABEL USER & PASSWORD
	var user entity.User
	// jika user tidak ada
	err := database.DB.First(&user, "email = ?", LoginRequest.Email).Error
	if err != nil {
		return ctx.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"error":   true,
			"message": "wrong credential",
		})
	}

	log.Println("user login:", user.Email, user.Password)

	// CHECK AVAILABLE PASSWORD
	isValid := utils.CheckPasswordHash(LoginRequest.Password, user.Password)
	if !isValid {
		return ctx.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"error":   true,
			"message": "wrong credential",
		})
	}

	// GENERATE TOKEN JWT
	claims := jwt.MapClaims{}
	claims["name"] = user.Name
	claims["email"] = user.Email
	claims["address"] = user.Address
	claims["exp"] = time.Now().Add(time.Minute * 2).Unix() // token ini akan valid selama 2 menit

	role := user.Role // mendapatkan role admin / role dari database untuk mendapatkan token login

	if role == "admin" {
		claims["role"] = "admin"
	} else {
		claims["role"] = "user"
	}

	// if user.Email == "eman13@gmail.com" {
	// 	claims["role"] = "admin"
	// } else {
	// 	claims["role"] = "user"
	// }

	token, errGenereateToken := utils.GenerateToken(&claims)

	if errGenereateToken != nil {
		log.Println(errGenereateToken)

		return ctx.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"error":   true,
			"message": "wrong credential",
		})
	}

	// jika user ada
	return ctx.JSON(fiber.Map{
		"message": token,
		// "message": "success login",
	})
}
