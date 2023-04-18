package activityController

import (
	"boilerplate/entities"
	"boilerplate/services/activityService"
	"fmt"
	"strconv"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func GetAll(ctx *fiber.Ctx) error {
	Datas, err := activityService.GetAll()
	if err != nil {
		return ctx.JSON(entities.Response{
			Status:  "Failed",
			Message: err.Error(),
		})
	}

	resp := entities.Response{
		Status:  "Success",
		Message: "Success",
		Data:    Datas,
	}
	return ctx.JSON(resp)
}

func Get(ctx *fiber.Ctx) error {
	id, _ := strconv.Atoi(ctx.Params("id"))

	site, err := activityService.Get(id)
	if err != nil {
		return ctx.JSON(entities.Response{
			Status:  "Failed",
			Message: err.Error(),
		})
	}

	resp := entities.Response{
		Status:  "Success",
		Message: "Success",
		Data:    site,
	}
	return ctx.JSON(resp)
}

func Create(ctx *fiber.Ctx) error {
	var Req *entities.Activity
	if err := ctx.BodyParser(&Req); err != nil {
		return ctx.JSON(entities.Response{
			Status:  "Failed",
			Message: err.Error(),
		})
	}
	Datas, err := activityService.Create(Req)
	if err != nil {
		return ctx.JSON(entities.Response{
			Status:  "Failed",
			Message: err.Error(),
		})
	}

	resp := entities.Response{
		Status:  "Success",
		Message: "Success",
		Data:    Datas,
	}
	return ctx.JSON(resp)
}

func Delete(ctx *fiber.Ctx) error {
	id, _ := strconv.Atoi(ctx.Params("id"))
	var message string
	err := activityService.Delete(id)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			message = "Activity with ID " + ctx.Params("id", "0") + " Not Found"
		} else {
			message = err.Error()
		}
		return ctx.JSON(entities.Response{
			Status:  "Failed",
			Message: message,
		})
	}
	data := make(map[string]interface{})

	resp := entities.Response{
		Status:  "Success",
		Message: "Success",
		Data:    data,
	}
	return ctx.JSON(resp)
}

func Update(ctx *fiber.Ctx) error {
	id, _ := strconv.Atoi(ctx.Params("id", "0"))
	var Req map[string]interface{}
	if err := ctx.BodyParser(&Req); err != nil || len(Req) == 0 {
		return ctx.JSON(entities.Response{
			Status:  "Bad Request",
			Message: "title cannot be null",
		})
	}

	fmt.Println(Req)
	Datas, err := activityService.Update(id, Req)
	var message string
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			message = "Activity with ID " + ctx.Params("id", "0") + " Not Found"
		} else {
			message = err.Error()
		}
		return ctx.JSON(entities.Response{
			Status:  "Not Found",
			Message: message,
		})
	}

	resp := entities.Response{
		Status:  "Success",
		Message: "Success",
		Data:    Datas,
	}
	return ctx.JSON(resp)
}
