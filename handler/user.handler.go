package handler

import (
	"backend-koperasi/database"
	"backend-koperasi/models/entity"
	"backend-koperasi/models/request"
	"backend-koperasi/models/response"
	"backend-koperasi/utils"
	"errors"
	"log"
	"strconv"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func GetAllUsers(ctx *fiber.Ctx) error {
	var users []*entity.User

	result := database.DB.Debug().Find(&users)

	if result.Error != nil {
		log.Fatal(result.Error)
	}

	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{
		"error":   false,
		"message": users,
	})
}

func CreateUser(ctx *fiber.Ctx) error {
	user := new(request.UserCreateRequest)

	if err := ctx.BodyParser(user); err != nil {
		return ctx.Status(fiber.StatusServiceUnavailable).JSON(err.Error())
	}

	// VALIDATION CREATE USER
	validate := validator.New()
	errValidator := validate.Struct(user)
	if errValidator != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error":   true,
			"message": errValidator.Error(),
		})
	}

	var existingEmail entity.User
	var existingNIK entity.User

	// Jika Email sudah terdaftar
	errExisitingUser := database.DB.Debug().Where("email = ?", user.Email).First(&existingEmail).Error
	if errExisitingUser == nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error":   true,
			"message": "email already registered",
		})
	} else if !errors.Is(errExisitingUser, gorm.ErrRecordNotFound) {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error":   true,
			"message": "error checking existing email",
		})
	}

	// Jika NIK sudah terdaftar
	errExistingNIK := database.DB.Debug().Where("nik = ?", user.NIK).First(&existingNIK).Error
	if errExistingNIK == nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error":   true,
			"message": "NIK already registered",
		})
	} else if !errors.Is(errExistingNIK, gorm.ErrRecordNotFound) {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error":   true,
			"message": "error chechking existing NIK",
		})
	}

	newUser := entity.User{
		Name:    user.Name,
		NIK:     user.NIK,
		Email:   user.Email,
		Address: user.Address,
		Phone:   user.Phone,
		Role:    user.Role,
	}

	// Proses Hashing Password
	hashedPassword, err := utils.HashingPassword(user.Password)
	if err != nil {
		log.Println(err)
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error":   true,
			"message": "internal server error",
		})
	}
	newUser.Password = hashedPassword
	// End Hashed Password

	errCreateUser := database.DB.Debug().Create(&newUser)
	if errCreateUser.Error != nil {
		log.Fatal(errCreateUser.Error)
	}

	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{ // StatusCode 200
		"error":   false,
		"message": "user has been created",
		"data":    newUser,
	})
}

func GetById(ctx *fiber.Ctx) error {
	var user []*response.User

	errById := database.DB.Debug().First(&user, ctx.Params("id"))

	// if errById.Error != nil {
	// 	return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
	// 		"error":   true,
	// 		"message": "user not found",
	// 	})
	// }

	if errById.RowsAffected == 0 {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error":   true,
			"message": "user not found",
		})
	} else if errById.Error != nil {
		return ctx.Status(fiber.StatusBadGateway).JSON(fiber.Map{
			"error":   true,
			"message": errById.Error,
		})
	}

	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{
		"error":   false,
		"message": user,
	})
}

func UpdateUser(ctx *fiber.Ctx) error {
	// user := new(entity.User)
	userRequest := new(request.UserUpdateRequest)

	var existingUser entity.User
	var existingNIK entity.User

	if err := ctx.BodyParser(userRequest); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(err.Error())
	}

	validate := validator.New()
	errValidator := validate.Struct(userRequest)
	if errValidator != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error":   true,
			"message": errValidator.Error(),
		})
	}

	id, _ := strconv.Atoi(ctx.Params("id"))

	// cek jika email belum terdaftar
	err := database.DB.Debug().Where("email = ?", userRequest.Email).First(&existingUser).Error
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error":   true,
			"message": "error checking existing email",
		})
	}
	// Cek jika email sudah terdaftar pada database
	if existingUser.ID != 0 && existingUser.ID != uint(id) {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error":   true,
			"message": "email already exist",
		})
	}

	// cek jika NIK sudah terdaftar
	errUpdateNIK := database.DB.Debug().Where("nik = ?", userRequest.NIK).First(&existingNIK).Error

	if existingNIK.ID != 0 && existingNIK.ID != uint(id) {
		return ctx.Status(fiber.StatusBadGateway).JSON(fiber.Map{
			"error":   true,
			"message": "NIK already exist",
		})
	} else if errUpdateNIK != nil && !errors.Is(errUpdateNIK, gorm.ErrRecordNotFound) {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error":   true,
			"message": "error checking existing nik",
		})
	}

	errUpdate := database.DB.Debug().Model(&entity.User{}).Where("id = ?", id).Updates(map[string]interface{}{
		"name":    userRequest.Name,
		"nik":     userRequest.NIK,
		"email":   userRequest.Email,
		"address": userRequest.Address,
		"phone":   userRequest.Phone,
	})

	if errUpdate.RowsAffected == 0 {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error":   true,
			"message": "id not found",
		})
	} else if errUpdate.Error != nil {
		return ctx.Status(fiber.StatusBadGateway).JSON(fiber.Map{
			"error":   true,
			"message": errUpdate.Error,
		})
	}

	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{
		"error":   false,
		"message": "user has been updated",
	})
}

func DeleteUser(ctx *fiber.Ctx) error {
	user := new(entity.User)

	id, _ := strconv.Atoi(ctx.Params("id"))

	result := database.DB.Debug().Where("id = ?", id).Delete(&user)

	// Handler jika ID User tidak tersedia
	if result.RowsAffected == 0 {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error":   true,
			"message": "user not found",
		})
	} else if result.Error != nil {
		return ctx.Status(fiber.StatusBadGateway).JSON(fiber.Map{
			"error":   true,
			"message": result.Error,
		})
	}

	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{
		"error":   false,
		"message": "user has been deleted",
	})
}
