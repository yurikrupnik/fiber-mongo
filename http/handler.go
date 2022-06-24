package http

import (
	"fiber-mongo/domain"
	"fmt"
	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"log"
)

type Handler struct {
	Svc domain.UserSvc
}

func NewHandler(svc domain.UserSvc) *Handler {
	return &Handler{
		Svc: svc,
	}
}

func (h *Handler) GetUser(c *fiber.Ctx) error {
	id := c.Params("id")
	objId, _ := primitive.ObjectIDFromHex(id)
	log.Println("-----> OBJECT ID", objId)
	data, err := h.Svc.Get(objId)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(err.Error())
	}
	return c.Status(200).JSON(data)
}

func (h *Handler) ListUsers(c *fiber.Ctx) error {
	u := new(domain.User)
	if err := c.QueryParser(u); err != nil {
		return err
	}
	result, err := h.Svc.List(u)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(err.Error())
	}
	return c.Status(fiber.StatusOK).JSON(result)
}

func (h *Handler) AddUser(c *fiber.Ctx) error {
	var user domain.User
	if err := c.BodyParser(&user); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(err.Error())
	}
	result, err := h.Svc.Create(&user)
	if err != nil {
		fmt.Println(fmt.Errorf("error - adding new user detail to db failed, err : %v", err))
		return c.Status(fiber.StatusInternalServerError).JSON(err.Error())
	}

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{"id": result.InsertedID})
}

func (h *Handler) DeleteUser(c *fiber.Ctx) error {
	id := c.Params("id")

	objId, _ := primitive.ObjectIDFromHex(id)
	log.Println("objId", objId)
	err := h.Svc.Delete(objId)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(err.Error())
	}
	return c.Status(fiber.StatusOK).JSON(objId)
}

func (h *Handler) Update(c *fiber.Ctx) error {
	id := c.Params("id")
	objId, _ := primitive.ObjectIDFromHex(id)
	var user domain.User
	if err := c.BodyParser(&user); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(err.Error())
	}
	err := h.Svc.Update(objId, &user)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(err.Error())
	}
	return c.SendString("updated")
}
