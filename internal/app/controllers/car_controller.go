package controllers

import (
	"cars/internal/app/models"
	"cars/internal/app/queries"
	"cars/internal/pkg/utils"
	"cars/internal/platform/database"
	"encoding/json"
	"strconv"

	"github.com/gofiber/fiber/v3"
	"github.com/sirupsen/logrus"
)

func CreateCar(c fiber.Ctx) error {
	car := &models.Car{}

	if err := json.Unmarshal(c.Body(), car); err != nil {
		logrus.Info(err)
		return c.Status(fiber.StatusBadRequest).JSON(models.GetResponse("Bad request"))
	}

	validate := utils.NewValidator()

	if err := validate.Struct(car); err != nil {
		logrus.Info(err)
		return c.Status(fiber.StatusBadRequest).JSON(models.GetResponse("Bad request"))
	}

	dbConn, err := database.OpenDBConn()
	if err != nil {
		logrus.Info(err)
		return c.Status(fiber.StatusBadRequest).JSON(models.GetResponse("Internal server error"))
	}

	if err := queries.CreateCar(dbConn, car); err != nil {
		logrus.Info(err)
		return c.Status(fiber.StatusBadRequest).JSON(models.GetResponse("Internal server error"))
	}

	return c.JSON(models.GetResponse("success"))
}

func DeleteCar(c fiber.Ctx) error {
	regNum := c.Params("reg", "")

	if len(regNum) == 0 {
		return c.Status(fiber.StatusBadRequest).JSON(models.GetResponse("Bad request"))
	}

	dbConn, err := database.OpenDBConn()
	if err != nil {
		logrus.Info(err)
		return c.Status(fiber.StatusBadRequest).JSON(models.GetResponse("Internal server error"))
	}

	if ok, err := queries.IsCarExists(dbConn, regNum); err != nil {
		logrus.Info(err)
		return c.Status(fiber.StatusInternalServerError).JSON(models.GetResponse("Internal server error"))
	} else if !ok {
		return c.Status(fiber.StatusBadRequest).JSON(models.GetResponse("Bad request"))
	}

	if err := queries.DeleteCar(dbConn, regNum); err != nil {
		logrus.Info(err)
		return c.Status(fiber.StatusInternalServerError).JSON(models.GetResponse("Internal server error"))
	}

	return c.JSON(models.GetResponse("success"))
}

func GetCar(c fiber.Ctx) error {
	regNum := c.Params("reg", "")

	if len(regNum) == 0 {
		return c.Status(fiber.StatusBadRequest).JSON(models.GetResponse("Bad request"))
	}

	dbConn, err := database.OpenDBConn()
	if err != nil {
		logrus.Info(err)
		return c.Status(fiber.StatusBadRequest).JSON(models.GetResponse("Internal server error"))
	}

	car, err := queries.GetCar(dbConn, regNum)

	if err != nil {
		logrus.Info(err)
		return c.Status(fiber.StatusInternalServerError).JSON(models.GetResponse("Internal server error"))
	} else if len(car.RegNum) == 0 {
		return c.Status(fiber.StatusBadRequest).JSON(models.GetResponse("Bad request"))
	} else {
		return c.JSON(car)
	}
}

func GetCars(c fiber.Ctx) error {
	limit, _ := strconv.Atoi(c.Query("limit", ""))
	if limit <= 0 || limit > 100 {
		limit = 10
	}

	offset, _ := strconv.Atoi(c.Query("offset", ""))
	if offset < 0 {
		offset = 0
	}

	dbConn, err := database.OpenDBConn()
	if err != nil {
		logrus.Info(err)
		return c.Status(fiber.StatusBadRequest).JSON(models.GetResponse("Internal server error"))
	}

	if cars, err := queries.GetCars(dbConn, limit, offset); err != nil {
		logrus.Info(err)
		return c.Status(fiber.StatusInternalServerError).JSON(models.GetResponse("Internal server error"))
	} else {
		return c.JSON(cars)
	}
}

func UpdateCar(c fiber.Ctx) error {
	car := &models.Car{}

	if err := json.Unmarshal(c.Body(), car); err != nil {
		logrus.Info(err)
		return c.Status(fiber.StatusBadRequest).JSON(models.GetResponse("Bad request"))
	}

	validate := utils.NewValidator()

	if err := validate.Struct(car); err != nil {
		logrus.Info(err)
		return c.Status(fiber.StatusBadRequest).JSON(models.GetResponse("Bad request"))
	}

	dbConn, err := database.OpenDBConn()
	if err != nil {
		logrus.Info(err)
		return c.Status(fiber.StatusBadRequest).JSON(models.GetResponse("Internal server error"))
	}

	if ok, err := queries.IsCarExists(dbConn, car.RegNum); err != nil {
		logrus.Info(err)
		return c.Status(fiber.StatusInternalServerError).JSON(models.GetResponse("Internal server error"))
	} else if !ok {
		return c.Status(fiber.StatusBadRequest).JSON(models.GetResponse("Bad request"))
	}

	if err := queries.UpdateCar(dbConn, car); err != nil {
		logrus.Info(err)
		return c.Status(fiber.StatusBadRequest).JSON(models.GetResponse("Internal server error"))
	}

	return c.JSON(models.GetResponse("success"))
}
